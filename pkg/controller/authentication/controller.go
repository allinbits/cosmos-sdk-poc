package authentication

import (
	"github.com/fdymylja/tmos/apis/meta"
	"github.com/fdymylja/tmos/pkg/controller/basic"
)

// Controller defines the authentication controller
type Controller interface {
	basic.Controller
	// DecodeTx takes raw transaction bytes and returns the state transitions, the accounts which have
	// authenticated the transitions, and an error
	Authenticate(txBytes []byte) (transitions []meta.StateTransition, authenticatedAccounts []string, err error)
}
