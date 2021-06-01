package statetransition

import (
	meta "github.com/fdymylja/tmos/core/meta"
	"github.com/fdymylja/tmos/runtime/authentication/user"
)

// ExecutionRequest is the request forwarded to the ExecutionHandler controller
type ExecutionRequest struct {
	// Users contains information on the entities that have authorized the meta.StateTransition
	Users user.Users
	// Transition is the meta.StateTransition that needs to be processed
	Transition meta.StateTransition
}

// ExecutionResponse is the response returned by ExecutionHandler controller
type ExecutionResponse struct{}

// ExecutionHandler identifies the state transition handler
// which handles the state transition and modifies state
// based on the received execution request.
type ExecutionHandler interface {
	// Exec is called when the ExecutionRequest needs to be processed
	Exec(req ExecutionRequest) (ExecutionResponse, error)
}

type ExecutionHandlerFunc func(req ExecutionRequest) (resp ExecutionResponse, err error)

func (s ExecutionHandlerFunc) Exec(req ExecutionRequest) (ExecutionResponse, error) {
	return s(req)
}
