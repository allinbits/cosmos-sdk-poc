package basic

import (
	"github.com/fdymylja/tmos/apis/meta"
)

// DeliverRequest is the request model used to transition state
type DeliverRequest struct {
	StateTransition meta.StateTransition
}

// DeliverResponse is the response returned by StateTransitionHandler
type DeliverResponse struct {
	Error error
}

type CheckRequest struct {
}

type CheckResponse struct {
}

// StateTransitionHandler defines a handler that takes care of transitioning state
type StateTransitionHandler interface {
	Check(CheckRequest) CheckResponse
	Deliver(DeliverRequest) DeliverResponse
}

// RegisterStateObjectsFn is used to register state objects to the runtime store
type RegisterStateObjectsFn func(objectMeta meta.StateObject)

// RegisterTransitionFn is used to register a state transition to the runtime
type RegisterTransitionFn func(transition meta.StateTransition, handler StateTransitionHandler)

// Controller defines the interface exposed by an basic
type Controller interface {
	// Name returns the unique name of the module
	Name() string
	// RegisterStateTransitions provides a Client to the module. This Client has an unique
	// identity in the system, that can then be used to interact with the store or interact
	// with other modules. The register function can then be used to register a meta.StateTransition
	// to the runtime and its respective handler.
	RegisterStateTransitions(client Client, register RegisterTransitionFn)
	// RegisterStateObjects provides a register function which allows modules to define
	// and register the meta.StateObjects they own to the runtime store.
	RegisterStateObjects(register RegisterStateObjectsFn)
	// RegisterStateTransitionInterceptors provides a Client to the module.
	// It can be used as a HOOK to intercept meta.StateTransition of other modules
	// before or after execution of the meta.StateTransition.

}
