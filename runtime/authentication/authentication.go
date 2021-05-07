package authentication

import (
	coin "github.com/fdymylja/tmos/module/core/coin/v1alpha1"
	"github.com/fdymylja/tmos/runtime/meta"
)

// Authenticator represents the module that takes care of authenticating transactions
// and returning the state transitions.
type Authenticator interface {
	// Authenticate authenticates the provided Tx (signature verification)
	// if it fails - a runtime.ErrUnauthorized will be returned to the caller
	Authenticate(tx Tx) error
	// DecodeTx simply decodes a transaction and returns a Tx.
	DecodeTx(txBytes []byte) (tx Tx, err error)
}

// Tx represents an authenticated request made to the *runtime.Runtime
type Tx interface {
	// StateTransitions returns the state transitions
	StateTransitions() []meta.StateTransition
	// Subjects returns the entities authenticated in the transaction
	Subjects() *Subjects
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
