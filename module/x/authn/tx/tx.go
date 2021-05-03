package tx

import (
	coin "github.com/fdymylja/tmos/module/core/coin/v1alpha1"
	"github.com/fdymylja/tmos/module/x/authn/crypto"
	"github.com/fdymylja/tmos/module/x/authn/v1alpha1"
	"github.com/fdymylja/tmos/runtime/authentication"
	"github.com/fdymylja/tmos/runtime/meta"
)

var _ authentication.Tx = (*Wrapper)(nil)

type Signer struct {
	Address   string
	PubKey    crypto.PubKey
	Signature []byte
}

// Wrapper wraps the raw *v1alpha1.Tx but contains parsed information
// regarding pub keys and such.
type Wrapper struct {
	raw         *v1alpha1.Tx
	bytes       []byte
	transitions []meta.StateTransition
	signers     *authentication.Subjects
	pubKeys     []Signer
	payer       string
}

func (t *Wrapper) StateTransitions() []meta.StateTransition {
	return t.transitions
}

func (t *Wrapper) Subjects() *authentication.Subjects {
	return t.signers
}

func (t *Wrapper) Fee() []*coin.Coin {
	return t.raw.AuthInfo.Fee.Amount
}

func (t *Wrapper) Payer() string {
	return t.payer
}

func (t *Wrapper) Raw() interface{} {
	return t.raw
}

func (t *Wrapper) RawBytes() []byte {
	return t.bytes
}

// Signers returns a map containing the account identifier (address) and the public key the user used to sign.
func (t *Wrapper) Signers() []Signer {
	return t.pubKeys
}
