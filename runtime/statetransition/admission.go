package statetransition

import (
	"github.com/fdymylja/tmos/runtime/authentication/user"
)

// AdmissionRequest is the request sent to admission controllers
type AdmissionRequest struct {
	// Transition defines the meta.StateTransition that needs to be validated
	Transition StateTransition
	// Users contains the authentication.Users that have authorized
	// the state transition
	Users user.Users
}

// AdmissionHandler defines the admission controller
// its duty is to verify that the state transition
// can be executed.
// It can read state but it cannot modify it.
type AdmissionHandler interface {
	// Validate validates the AdmissionRequest and returns
	// an error if the request is invalid.
	// If one AdmissionHandler in the meta.StateTransition admission chain
	// fails then the execution is fully stopped.
	// And state rolled back accordingly to the current execution phase.
	Validate(AdmissionRequest) error
}

// AdmissionHandlerFunc implements AdmissionHandler and allows to create a AdmissionHandler
// using a function.
type AdmissionHandlerFunc func(req AdmissionRequest) error

// Validate implements AdmissionHandler
func (h AdmissionHandlerFunc) Validate(req AdmissionRequest) error {
	return h(req)
}

// AdmissionChain contains multiple AdmissionHandler and executes them in a chain
// NOTE: each AdmissionHandler MUST handle the same meta.StateTransition
type AdmissionChain struct {
	handlers []AdmissionHandler
}

func (c AdmissionChain) Validate(req AdmissionRequest) error {
	for _, handler := range c.handlers {
		err := handler.Validate(req)
		if err != nil {
			return err
		}
	}
	return nil
}

// NewAdmissionChain instantiates a new AdmissionChain of AdmissionHandler
// NOTE: each AdmissionHandler MUST handle the same meta.StateTransition
func NewAdmissionChain(handlers ...AdmissionHandler) AdmissionChain {
	return AdmissionChain{handlers: handlers}
}
