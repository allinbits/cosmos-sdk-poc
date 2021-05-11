package authentication

import (
	coin "github.com/fdymylja/tmos/module/core/coin/v1alpha1"
	"github.com/fdymylja/tmos/runtime/authentication/user"
	"github.com/fdymylja/tmos/runtime/meta"
)

// TxDecoder represents the component which decodes transactions given raw bytes
type TxDecoder interface {
	// DecodeTx simply decodes a transaction and returns a Tx.
	DecodeTx(txBytes []byte) (tx Tx, err error)
}

// Tx represents an authenticated request made to the *runtime.Runtime
type Tx interface {
	// StateTransitions returns the state transitions
	StateTransitions() []meta.StateTransition
	// Users returns the entities authenticated in the transaction
	Users() user.Users
	// Fee returns the fees of the transaction
	// as a slice of *coin.Coin
	Fee() []*coin.Coin
	// Payer returns the subject who is paying for the transaction
	Payer() string
	// Raw returns the raw transaction as interface
	// for modules which wish to interact with concrete types
	// of the transaction which are dependent on the authentication
	// module which is being used
	Raw() interface{}
	// RawBytes returns the raw transaction bytes
	RawBytes() []byte
}
