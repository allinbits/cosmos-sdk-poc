package extensions

import (
	"fmt"

	kmultisig "github.com/cosmos/cosmos-sdk/crypto/keys/multisig"
	abciv1alpha1 "github.com/fdymylja/tmos/module/abci/v1alpha1"
	"github.com/fdymylja/tmos/module/x/authn/crypto"
	"github.com/fdymylja/tmos/module/x/authn/tx"
	"github.com/fdymylja/tmos/module/x/authn/v1alpha1"
	"github.com/fdymylja/tmos/runtime/authentication"
	"github.com/fdymylja/tmos/runtime/module"
)

func New(c module.Client) module.AuthenticationExtension {
	return authExtension{c: c}
}

type authExtension struct {
	c module.Client
}

func (a authExtension) Initialize(builder *module.AuthenticationExtensionBuilder) {
	// add admission controllers, they will only read state
	// never modify it.
	builder.
		WithAdmissionController(mempoolFee{}).                  // verifies if fee matches the minimum
		WithAdmissionController(newAccountExists(a.c)).         // verifies that all signer accounts exist
		WithAdmissionController(newTimeoutBlockExtension(a.c)). // verifies if tx is not timed-out compared to block
		WithAdmissionController(newValidateMemoExtension(a.c)). // validates memo length
		WithAdmissionController(newValidateSigCount(a.c))       // validates number of signatures
	// add transition controllers for tx, they CAN modify state after
	// a tx is authenticated
	builder.
		WithTransitionController(newConsumeGasForTxSize(a.c)). // consumes gas for tx size
		WithTransitionController(newSetPubKeys(a.c)).          // sets pub keys
		WithTransitionController(newIncreaseSequence(a.c))     // increases sequence

}

func newAccountExists(c module.Client) accountExists {
	return accountExists{
		c: v1alpha1.NewClient(c),
	}
}

type accountExists struct {
	c *v1alpha1.Client
}

func (a accountExists) Validate(tx authentication.Tx) error {
	// assert that all signer accounts exist
	for _, sig := range tx.Subjects().List() {
		_, err := a.c.GetAccount(sig)
		if err != nil {
			return err
		}
	}
	return nil
}

func newTimeoutBlockExtension(c module.Client) timeoutBlockExtension {
	return timeoutBlockExtension{abci: abciv1alpha1.NewClient(c)}
}

type timeoutBlockExtension struct {
	abci *abciv1alpha1.Client
}

func (t timeoutBlockExtension) Validate(reqTx authentication.Tx) error {
	tx := reqTx.Raw().(*v1alpha1.Tx)
	currentBlock, err := t.abci.GetCurrentBlock()
	if err != nil {
		return err
	}
	if tx.Body.TimeoutHeight != 0 && currentBlock.BlockNumber > tx.Body.TimeoutHeight {
		return fmt.Errorf("invalid block height")
	}

	return nil
}

func newValidateMemoExtension(c module.Client) validateMemoExtension {
	return validateMemoExtension{c: v1alpha1.NewClient(c)}
}

type validateMemoExtension struct {
	c *v1alpha1.Client
}

func (e validateMemoExtension) Validate(reqTx authentication.Tx) error {
	tx := reqTx.Raw().(*v1alpha1.Tx)
	memo := tx.Body.Memo
	memoLength := len(memo)
	params, err := e.c.GetParams()
	if err != nil {
		return err
	}
	if uint64(memoLength) > params.MaxMemoCharacters {
		return fmt.Errorf("invalid memo length")
	}
	return nil
}

// mempoolFee is used to check if the transaction meets the minimum mempool requirements
type mempoolFee struct {
}

func (e mempoolFee) Validate(tx authentication.Tx) error {
	return nil
}

func newValidateSigCount(c module.Client) sigCount {
	return sigCount{c: v1alpha1.NewClient(c)}
}

type sigCount struct {
	c *v1alpha1.Client
}

func (s sigCount) Validate(reqTx authentication.Tx) error {
	wrapper := reqTx.(*tx.Wrapper) // TODO: should we return an error? or simply skip validation?
	params, err := s.c.GetParams()
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

func (s sigCount) countSubKeys(pk crypto.PubKey) int {
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

func (s consumeGasForTxSize) Deliver(req authentication.DeliverRequest) (authentication.DeliverResponse, error) {
	// TODO
	return authentication.DeliverResponse{}, nil
}

func newSetPubKeys(c module.Client) setPubKeys {
	return setPubKeys{
		c: v1alpha1.NewClient(c),
	}
}

type setPubKeys struct {
	c *v1alpha1.Client
}

func (s setPubKeys) Deliver(req authentication.DeliverRequest) (authentication.DeliverResponse, error) {
	wrapper := req.Tx.(*tx.Wrapper)
	for _, sig := range wrapper.Signers() {
		acc, err := s.c.GetAccount(sig.Address)
		if err != nil {
			return authentication.DeliverResponse{}, err
		}
		// skip if its set
		if acc.PubKey != nil {
			continue
		}
		err = s.c.UpdatePublicKey(sig.Address, sig.PubKey)
		if err != nil {
			return authentication.DeliverResponse{}, err
		}
	}
	return authentication.DeliverResponse{}, nil
}

func newIncreaseSequence(c module.Client) increaseSequence {
	return increaseSequence{c: v1alpha1.NewClient(c)}
}

type increaseSequence struct {
	c *v1alpha1.Client
}

func (i increaseSequence) Deliver(req authentication.DeliverRequest) (authentication.DeliverResponse, error) {
	signers := req.Tx.Subjects()
	for _, signer := range signers.List() {
		err := i.c.IncreaseSequence(signer)
		if err != nil {
			return authentication.DeliverResponse{}, err
		}
	}
	return authentication.DeliverResponse{}, nil
}
