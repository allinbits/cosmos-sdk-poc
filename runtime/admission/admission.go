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

// Handler defines the admission controller
// its duty is to verify that the state transition
// can be executed.
// It can read state but it cannot modify it.
type Handler interface {
	// Validate validates the StateTransitionRequest and returns
	// an error if the request is invalid.
	// If one Handler in the meta.StateTransition admission chain
	// fails then the execution is fully stopped.
	// And state rolled back accordingly to the current execution phase.
	Validate(StateTransitionRequest) error
}

// HandlerFunc implements Handler and allows to create a Handler
// using a function.
type HandlerFunc func(req StateTransitionRequest) error

// Validate implements Handler
func (h HandlerFunc) Validate(req StateTransitionRequest) error {
	return h(req)
}
