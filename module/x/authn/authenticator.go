package authn

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/x/auth/signing"
	abciv1alpha1 "github.com/fdymylja/tmos/module/abci/v1alpha1"
	"github.com/fdymylja/tmos/module/x/authn/tx"
	"github.com/fdymylja/tmos/module/x/authn/v1alpha1"
	"github.com/fdymylja/tmos/runtime/authentication"
	"github.com/fdymylja/tmos/runtime/module"
)

func newAuthenticator(c module.Client) authenticator {
	return authenticator{
		abci:    abciv1alpha1.NewClient(c),
		auth:    v1alpha1.NewClient(c),
		decoder: tx.NewDecoder("test"),
		c:       c,
	}
}

type authenticator struct {
	abci         *abciv1alpha1.Client
	auth         *v1alpha1.Client
	decoder      *tx.Decoder
	bech32Prefix string
	c            module.Client
}

// Authenticate takes care of authenticating an authentication.Tx
func (a authenticator) Authenticate(aTx authentication.Tx) error {
	w := aTx.(*tx.Wrapper)
	sigs := w.Signers()
	for _, signer := range sigs {
		// get account
		acc, err := a.auth.GetAccount(signer.Address)
		if err != nil {
			return err
		}
		if acc.PubKey == nil {
			return fmt.Errorf("pub key not set on account %s", signer.Address)
		}

		signerData := signing.SignerData{
			ChainID:       "",
			AccountNumber: 0,
			Sequence:      0,
		}
	}
}

func (a authenticator) DecodeTx(txBytes []byte) (tx authentication.Tx, err error) {
	return a.decoder.Decode(txBytes)
}
