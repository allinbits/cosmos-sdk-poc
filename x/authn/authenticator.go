package authn

import (
	"fmt"

	abciv1alpha1 "github.com/fdymylja/tmos/module/abci/v1alpha1"
	"github.com/fdymylja/tmos/runtime/authentication"
	"github.com/fdymylja/tmos/runtime/module"
	tx2 "github.com/fdymylja/tmos/x/authn/tx"
	v1alpha12 "github.com/fdymylja/tmos/x/authn/v1alpha1"
	"google.golang.org/protobuf/proto"
)

func newAuthenticator(c module.Client) authenticator {
	return authenticator{
		abci:    abciv1alpha1.NewClient(c),
		auth:    v1alpha12.NewClient(c),
		decoder: tx2.NewDecoder("test"),
		c:       c,
	}
}

type authenticator struct {
	abci         *abciv1alpha1.Client
	auth         *v1alpha12.Client
	decoder      *tx2.Decoder
	bech32Prefix string
	c            module.Client
}

// Authenticate takes care of authenticating an authentication.Tx
func (a authenticator) Authenticate(aTx authentication.Tx) error {
	wrapper := aTx.(*tx2.Wrapper)
	raw := wrapper.TxRaw()
	sigs := wrapper.Signers()
	// get chainID
	chainID, err := a.abci.GetChainID()
	if err != nil {
		return err
	}
	for _, signer := range sigs {
		// get account
		acc, err := a.auth.GetAccount(signer.Address)
		if err != nil {
			return err
		}
		if acc.PubKey == nil {
			return fmt.Errorf("pub key not set on account %s", signer.Address)
		}

		expectedBytes, err := DirectSignBytes(raw.BodyBytes, raw.AuthInfoBytes, chainID, acc.AccountNumber)
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

func (a authenticator) DecodeTx(txBytes []byte) (tx authentication.Tx, err error) {
	return a.decoder.Decode(txBytes)
}
