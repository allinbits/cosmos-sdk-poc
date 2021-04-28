package authentication

import "github.com/fdymylja/tmos/runtime/meta"

type Authenticator interface {
	// DecodeTx simply decodes a transaction and returns its state transitions
	DecodeTx(txBytes []byte) (transitions []meta.StateTransition, err error)
}
