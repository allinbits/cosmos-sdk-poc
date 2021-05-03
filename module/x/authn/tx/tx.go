package tx

import (
	coin "github.com/fdymylja/tmos/module/core/coin/v1alpha1"
	"github.com/fdymylja/tmos/module/x/authn/crypto"
	"github.com/fdymylja/tmos/module/x/authn/v1alpha1"
	"github.com/fdymylja/tmos/runtime/authentication"
	"github.com/fdymylja/tmos/runtime/meta"
)

var _ authentication.Tx = (*Wrapper)(nil)

// Wrapper wraps the raw *v1alpha1.Tx but contains parsed information
// regarding pub keys and such.
type Wrapper struct {
	raw         *v1alpha1.Tx
	bytes       []byte
	transitions []meta.StateTransition
	signers     *authentication.Subjects
	pubKeys     []crypto.PubKey
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

func (t *Wrapper) PubKeys() []crypto.PubKey {
	return t.pubKeys
}
