package application

import "github.com/fdymylja/tmos/pkg/codec"

// HandlerPolicy defines handler options
type HandlerPolicy struct {
	// Internal if true will make the state transition invalid if called externally
	Internal bool
}

// StateObject defines an object that can be saved inside the state
type StateObject interface {
	codec.DeterministicObject
}

// StateObjectRegisterer defines an object that takes care of registering
// state objects for the given application
type StateObjectRegisterer interface {
	// RegisterStateObject registers a state object
	// which is going to belong to the Application's namespace
	RegisterStateObject(object StateObject)
}

// StateTransitionObject defines an object that handles state transitions
type StateTransitionObject interface {
	codec.DeterministicObject
}

// DeliverClient defines an application client that can be used
// to get or store information in state
type DeliverClient interface {
	// Deliver can be used to deliver state transitions to other applications
	// from within an application (instead of externally).
	// NOTE: RBAC policies apply.
	Deliver(object StateTransitionObject) error
	// Get gets a state object, it can be used to fetch
	// StateObject from other applications
	Get(id string, object StateObject) (exists bool)
	// Set sets a state object which belongs to the application.
	// it cannot be used to modify state of other applications.
	Set(id string, object StateObject) error
	// Delete deletes a state object which belongs to the application
	// it cannot be used to modify state of other applications.
	Delete(id string)
}

// DeliverRequest defines the request used to
type DeliverRequest struct {
	StateTransitionObject StateTransitionObject
	Client                DeliverClient
}

// DeliverHandler defines a handler which should handle delivering state transitions requests
type DeliverHandler interface {
	Deliver(DeliverRequest) error
}

type CheckRequest struct {
	StateTransitionObject StateTransitionObject
}

// CheckHandler defines a handler which should check for correctness of state transitions
type CheckHandler interface {
	Check(CheckRequest) error
}

// HandlerRegisterer defines an object that takes care of registering handlers
// for the given application
type HandlerRegisterer interface {
	// RegisterDeliverHandler registers a deliver handler
	RegisterDeliverHandler(StateTransitionObject, DeliverHandler, CheckHandler)
	// RegisterBeginBlockHandler registers the begin block handler
	RegisterBeginBlockHandler(BeginBlockHandler)
	// RegisterEndBlockHandler registers the end block handler
	RegisterEndBlockHandler(EndBlockHandler)
	// RegisterHandlerHook registers hooks that can be used to execute actions
	// after a certain state transition is processed
	RegisterHandlerHook(StateTransitionObject, HookHandler)
}

type BeginBlockHandler interface {
}

type EndBlockHandler interface {
}

// HookHandler is used to handle StateTransitionObject hooks
type HookHandler interface {
}

type HookExecutionPolicy struct {
	ExecuteBefore bool
	ExecuteAfter  bool
}

type HandlerHookRegisterer interface {
	RegisterHandlerHook(StateTransitionObject, HookHandler, HookExecutionPolicy)
}

type InvariantsHandlerRegisterer interface {
}

// Application defines an application of the tendermint operating system
type Application interface {
	// Identifier identifies an application in the runtime
	Identifier() string
	// RegisterStateObjects is used by applications to register the
	// objects which will be saved in the state
	RegisterStateObjects(StateObjectRegisterer)
	// RegisterHandlers is used by the application to register DeliverTx,
	// CheckTx, BeginBlock, EndBlock, Hook and Invariance handlers
	RegisterHandlers(HandlerRegisterer)
}
