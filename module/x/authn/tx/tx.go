package tx

import (
	coin "github.com/fdymylja/tmos/module/core/coin/v1alpha1"
	"github.com/fdymylja/tmos/module/x/authn/crypto"
	"github.com/fdymylja/tmos/module/x/authn/v1alpha1"
	"github.com/fdymylja/tmos/runtime/authentication"
	"github.com/fdymylja/tmos/runtime/meta"
)

var _ authentication.Tx = (*Tx)(nil)

type Tx struct {
	raw         *v1alpha1.Tx
	bytes       []byte
	transitions []meta.StateTransition
	signers     *authentication.Subjects
	pubKeys     []crypto.PubKey
	payer       string
}

func (t *Tx) StateTransitions() []meta.StateTransition {
	return t.transitions
}

func (t *Tx) Subjects() *authentication.Subjects {
	return t.signers
}

func (t *Tx) Fee() []*coin.Coin {
	return t.raw.AuthInfo.Fee.Amount
}

func (t *Tx) Payer() string {
	return t.payer
}

func (t *Tx) Raw() interface{} {
	return t.raw
}

func (t *Tx) RawBytes() []byte {
	return t.bytes
}
