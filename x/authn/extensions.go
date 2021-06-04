package authn

import (
	"fmt"

	"github.com/fdymylja/tmos/runtime/client"

	kmultisig "github.com/cosmos/cosmos-sdk/crypto/keys/multisig"
	abciv1alpha1 "github.com/fdymylja/tmos/core/abci/v1alpha1"
	"github.com/fdymylja/tmos/runtime/authentication"
	"github.com/fdymylja/tmos/runtime/module"
	crypto2 "github.com/fdymylja/tmos/x/authn/crypto"
	tx2 "github.com/fdymylja/tmos/x/authn/tx"
	v1alpha12 "github.com/fdymylja/tmos/x/authn/v1alpha1"
	"google.golang.org/protobuf/proto"
)

func newAccountExists() accountExists {
	return accountExists{}
}

type accountExists struct{}

func (a accountExists) Validate(client client.RuntimeClient, tx authentication.Tx) error {
	c := v1alpha12.NewClient(client)

	// assert that all signer accounts exist
	for _, sig := range tx.Users().List() {
		_, err := c.GetAccount(sig.GetName())
		if err != nil {
			return err
		}
	}
	return nil
}

func newTimeoutBlockExtension() timeoutBlockExtension {
	return timeoutBlockExtension{}
}

type timeoutBlockExtension struct{}

func (t timeoutBlockExtension) Validate(client client.RuntimeClient, reqTx authentication.Tx) error {
	abciClient := abciv1alpha1.NewClientSet(client)
	tx := reqTx.Raw().(*v1alpha12.Tx)
	currentBlock, err := abciClient.CurrentBlock().Get()
	if err != nil {
		return err
	}

	if tx.Body.TimeoutHeight != 0 && currentBlock.BlockNumber > tx.Body.TimeoutHeight {
		return fmt.Errorf("invalid block height")
	}

	return nil
}

func newValidateMemoExtension() validateMemoExtension {
	return validateMemoExtension{}
}

type validateMemoExtension struct{}

func (e validateMemoExtension) Validate(client client.RuntimeClient, reqTx authentication.Tx) error {
	c := v1alpha12.NewClient(client)

	tx := reqTx.Raw().(*v1alpha12.Tx)
	memo := tx.Body.Memo
	memoLength := len(memo)
	params, err := c.GetParams()
	if err != nil {
		return err
	}
	if uint64(memoLength) > params.MaxMemoCharacters {
		return fmt.Errorf("invalid memo length")
	}
	return nil
}

// mempoolFee is used to check if the transaction meets the minimum mempool requirements
type mempoolFee struct{}

func (e mempoolFee) Validate(client client.RuntimeClient, tx authentication.Tx) error {
	return nil
}

func newValidateSigCount() sigCount {
	return sigCount{}
}

type sigCount struct{}

func (s sigCount) Validate(client client.RuntimeClient, reqTx authentication.Tx) error {
	c := v1alpha12.NewClient(client)
	wrapper := reqTx.(*tx2.Wrapper) // TODO: should we return an error? or simply skip validation?
	params, err := c.GetParams()
	if err != nil {
		return err
	}
	pubKeys := wrapper.Signers()
	sigs := 0
	for _, pk := range pubKeys {
		subKeys := s.countSubKeys(pk.PubKey)
		sigs += subKeys
		if uint64(sigs) > params.TxSigLimit {
			return fmt.Errorf("number of maximum signatures is %d got %d", params.TxSigLimit, sigs)
		}
	}

	return nil
}

func (s sigCount) countSubKeys(pk crypto2.PubKey) int {
	v, ok := pk.(*kmultisig.LegacyAminoPubKey)
	if !ok {
		return 1
	}

	numKeys := 0
	for _, subkey := range v.GetPubKeys() {
		numKeys += s.countSubKeys(subkey)
	}
	return numKeys
}

func newConsumeGasForTxSize(c module.Client) consumeGasForTxSize {
	return consumeGasForTxSize{}
}

type consumeGasForTxSize struct {
}

func (s consumeGasForTxSize) Exec(req authentication.PostAuthenticationRequest) (authentication.PostAuthenticationResponse, error) {
	// TODO
	return authentication.PostAuthenticationResponse{}, nil
}

func newSetPubKeys(c module.Client) setPubKeys {
	return setPubKeys{
		c: v1alpha12.NewClient(c),
	}
}

type setPubKeys struct {
	c *v1alpha12.Client
}

func (s setPubKeys) Exec(req authentication.PostAuthenticationRequest) (authentication.PostAuthenticationResponse, error) {
	wrapper := req.Tx.(*tx2.Wrapper)
	for _, sig := range wrapper.Signers() {
		acc, err := s.c.GetAccount(sig.Address)
		if err != nil {
			return authentication.PostAuthenticationResponse{}, err
		}
		// skip if its set
		if acc.PubKey != nil {
			continue
		}
		err = s.c.UpdatePublicKey(sig.Address, sig.PubKey)
		if err != nil {
			return authentication.PostAuthenticationResponse{}, err
		}
	}
	return authentication.PostAuthenticationResponse{}, nil
}

func newIncreaseSequence(c module.Client) increaseSequence {
	return increaseSequence{c: v1alpha12.NewClient(c)}
}

type increaseSequence struct {
	c *v1alpha12.Client
}

func (i increaseSequence) Exec(req authentication.PostAuthenticationRequest) (authentication.PostAuthenticationResponse, error) {
	signers := req.Tx.Users()
	for _, signer := range signers.List() {
		err := i.c.IncreaseSequence(signer.GetName())
		if err != nil {
			return authentication.PostAuthenticationResponse{}, err
		}
	}

	return authentication.PostAuthenticationResponse{}, nil
}

func newSigVerifier() sigVerifier {
	return sigVerifier{}
}

type sigVerifier struct{}

func (a sigVerifier) Validate(client client.RuntimeClient, aTx authentication.Tx) error {
	abciClient := abciv1alpha1.NewClientSet(client)
	authClient := v1alpha12.NewClient(client)

	wrapper := aTx.(*tx2.Wrapper)
	raw := wrapper.TxRaw()
	sigs := wrapper.Signers()
	// get chainInfo
	chainInfo, err := abciClient.InitChainInfo().Get()
	if err != nil {
		return err
	}
	for _, signer := range sigs {
		// get account
		acc, err := authClient.GetAccount(signer.Address)
		if err != nil {
			return err
		}
		if acc.PubKey == nil {
			return fmt.Errorf("pub key not set on account %s", signer.Address)
		}
		// check sequence
		if acc.Sequence != signer.Sequence {
			return fmt.Errorf("invalid sequence %d expected %d", signer.Sequence, acc.Sequence)
		}
		expectedBytes, err := DirectSignBytes(raw.BodyBytes, raw.AuthInfoBytes, chainInfo.ChainId, acc.AccountNumber)
		if err != nil {
			return err
		}
		if !signer.PubKey.VerifySignature(expectedBytes, signer.Signature) {
			return fmt.Errorf("bad sig")
		}
	}
	return nil
}

// DirectSignBytes returns the SIGN_MODE_DIRECT sign bytes for the provided TxBody bytes, AuthInfo bytes, chain ID,
// account number and sequence.
func DirectSignBytes(bodyBytes, authInfoBytes []byte, chainID string, accnum uint64) ([]byte, error) {
	signDoc := &v1alpha12.SignDoc{
		BodyBytes:     bodyBytes,
		AuthInfoBytes: authInfoBytes,
		ChainId:       chainID,
		AccountNumber: accnum,
	}
	return proto.Marshal(signDoc)
}
