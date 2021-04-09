package application

import (
	metav1alpha1 "github.com/fdymylja/tmos/apis/core/meta/v1alpha1"
	"github.com/fdymylja/tmos/pkg/codec"
)

// HandlerPolicy defines handler options
type HandlerPolicy struct {
	// Internal if true will make the state transition invalid if called externally
	Internal bool
}

// StateObject defines an object that can be saved inside the state
type StateObject interface {
	// GetMeta comes from implementing proto.Messages in .proto files
	// and adding the field core.meta.v1alpha1 wth name meta
	GetState() *metav1alpha1.State
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
	GetTransition() *metav1alpha1.Transition
	codec.DeterministicObject
}

// StateObjectClient defines a client that allows to
// get, set, delete StateObject in the state
type StateObjectClient interface {
	// Get gets a state object, it can be used to fetch
	// StateObject from other applications
	Get(object StateObject) (exists bool)
	// Set sets a state object which belongs to the application.
	// it cannot be used to modify state of other applications.
	Set(object StateObject)
	// Delete deletes a state object which belongs to the application
	// it cannot be used to modify state of other applications.
	Delete(object StateObject)
}

// DeliverClient defines an application client that can be used
// to get or store information in state
type DeliverClient interface {
	StateObjectClient
	// Deliver can be used to deliver state transitions to other applications
	// from within an application (instead of externally).
	// NOTE: RBAC policies apply.
	Deliver(object StateTransitionObject) error
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

// DeliverHandlerFn implements DeliverHandler but allows to define DeliverHandler via function instead of a struct
type DeliverHandlerFn func(DeliverRequest) error

func (d DeliverHandlerFn) Deliver(request DeliverRequest) error {
	return d(request)
}

type CheckRequest struct {
	StateTransitionObject StateTransitionObject
}

// CheckHandler defines a handler which should check for correctness of state transitions
type CheckHandler interface {
	Check(CheckRequest) error
}

// CheckHandlerFn implements CheckHandler but allows to define a CheckHandler via function instead of a struct
type CheckHandlerFn func(CheckRequest) error

func (f CheckHandlerFn) Check(request CheckRequest) error {
	return f(request)
}

// HandlerRegisterer defines an object that takes care of registering handlers
// for the given application
type HandlerRegisterer interface {
	// RegisterDeliverHandler registers a deliver handler, if CheckHandler is nil
	// a no-op CheckHandler will be used
	RegisterDeliverHandler(StateTransitionObject, DeliverHandler, CheckHandler, HandlerPolicy)
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
