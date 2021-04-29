package authentication

import (
	"github.com/fdymylja/tmos/runtime/meta"
)

// Authenticator represents the module that takes care of authenticating transactions
// and returning the state transitions.
type Authenticator interface {
	// Authenticate authenticates a transaction and returns the subjects (authenticated identities)
	// and the effective state transitions
	Authenticate(txBytes []byte) (subjects []string, transitions []meta.StateTransition, err error)
	// DecodeTx simply decodes a transaction and returns its state transitions
	DecodeTx(txBytes []byte) (transitions []meta.StateTransition, err error)
}

// Tx represents the transaction
type Tx interface {
	StateTransitions() []meta.StateTransition
	Subjects() []string
}
