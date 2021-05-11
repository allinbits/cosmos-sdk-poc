package admission

import (
	"github.com/fdymylja/tmos/runtime/authentication/user"
	"github.com/fdymylja/tmos/runtime/meta"
)

// Request is the request sent to admission controllers
type Request struct {
	// Transition defines the meta.StateTransition that needs to be validated
	Transition meta.StateTransition
	// Users contains the authentication.Users that have authorized
	// the state transition
	Users user.Users
}

// Handler defines the admission controller
// its duty is to verify that the state transition
// can be executed.
// It can read state but it cannot modify it.
type Handler interface {
	// Validate validates the Request and returns
	// an error if the request is invalid.
	// If one Handler in the meta.StateTransition admission chain
	// fails then the execution is fully stopped.
	// And state rolled back accordingly to the current execution phase.
	Validate(Request) error
}

// HandlerFunc implements Handler and allows to create a Handler
// using a function.
type HandlerFunc func(req Request) error

// Validate implements Handler
func (h HandlerFunc) Validate(req Request) error {
	return h(req)
}

// Chain contains multiple Handler and executes them in a chain
// NOTE: each Handler MUST handle the same meta.StateTransition
type Chain struct {
	handlers []Handler
}

func (c Chain) Validate(req Request) error {
	for _, handler := range c.handlers {
		err := handler.Validate(req)
		if err != nil {
			return err
		}
	}
	return nil
}

// NewChain instantiates a new Chain of Handler
// NOTE: each Handler MUST handle the same meta.StateTransition
func NewChain(handlers ...Handler) Chain {
	return Chain{handlers: handlers}
}
