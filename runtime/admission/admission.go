package admission

import (
	"github.com/fdymylja/tmos/runtime/authentication"
	"github.com/fdymylja/tmos/runtime/meta"
)

// StateTransitionRequest is the request sent to admission controllers
type StateTransitionRequest struct {
	// Transition defines the meta.StateTransition that needs to be validated
	Transition meta.StateTransition
	// Subjects contains the authentication.Subjects that have authorized
	// the state transition
	Subjects *authentication.Subjects
}

// Controller defines the admission controller
// its duty is to verify without accessing state
// that the provided state transition is valid
type Controller interface {
	// Validate validates the StateTransitionRequest and returns
	// an error if the request is invalid.
	// If one Controller in the meta.StateTransition admission chain
	// fails then the execution is fully stopped.
	// And state rolled back accordingly to the current execution phase.
	Validate(StateTransitionRequest) error
}
