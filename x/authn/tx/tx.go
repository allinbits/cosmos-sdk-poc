package tx

import (
	v1alpha12 "github.com/fdymylja/tmos/core/coin/v1alpha1"
	"github.com/fdymylja/tmos/runtime/authentication"
	"github.com/fdymylja/tmos/runtime/authentication/user"
	"github.com/fdymylja/tmos/runtime/meta"
	"github.com/fdymylja/tmos/x/authn/crypto"
	"github.com/fdymylja/tmos/x/authn/v1alpha1"
)

var _ authentication.Tx = (*Wrapper)(nil)

type Signer struct {
	Address   string
	PubKey    crypto.PubKey
	Signature []byte
	Sequence  uint64
}

// Wrapper wraps the raw *v1alpha1.Tx but contains parsed information
// regarding pub keys and such.
type Wrapper struct {
	txRaw       *v1alpha1.TxRaw
	raw         *v1alpha1.Tx
	bytes       []byte
	transitions []meta.StateTransition
	signers     user.Users
	pubKeys     []Signer
	payer       string
}

func (t *Wrapper) StateTransitions() []meta.StateTransition {
	return t.transitions
}

func (t *Wrapper) Users() user.Users {
	return t.signers
}

func (t *Wrapper) Fee() []*v1alpha12.Coin {
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

func (t *Wrapper) TxRaw() *v1alpha1.TxRaw {
	return t.txRaw
}

// Signers returns a map containing the account identifier (address) and the public key the user used to sign.
func (t *Wrapper) Signers() []Signer {
	return t.pubKeys
}
