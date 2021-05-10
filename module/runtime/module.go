package runtime

import (
	"github.com/fdymylja/tmos/module/runtime/v1alpha1"
	"github.com/fdymylja/tmos/runtime/controller"
	"github.com/fdymylja/tmos/runtime/module"
)

func NewModule() Module { return Module{} }

type Module struct {
}

func (m Module) Initialize(client module.Client, builder *module.Builder) {
	builder.
		Named("runtime").
		OwnsStateObject(&v1alpha1.StateObjectsList{}).
		OwnsStateObject(&v1alpha1.StateTransitionsList{}).
		HandlesStateTransition(&v1alpha1.CreateStateTransitionsList{}, newCreateStateTransitionsController(client), false).
		HandlesStateTransition(&v1alpha1.CreateStateObjectsList{}, newCreateStateObjectsController(client), false)
}

func newCreateStateObjectsController(client module.Client) controller.StateTransitionFn {
	return func(req controller.StateTransitionRequest) (resp controller.StateTransitionResponse, err error) {
		msg := req.Transition.(*v1alpha1.CreateStateObjectsList)
		return resp, client.Create(&v1alpha1.StateObjectsList{StateObjects: msg.StateObjects})
	}
}

func newCreateStateTransitionsController(client module.Client) controller.StateTransitionFn {
	return func(req controller.StateTransitionRequest) (resp controller.StateTransitionResponse, err error) {
		msg := req.Transition.(*v1alpha1.CreateStateTransitionsList)
		return resp, client.Create(&v1alpha1.StateTransitionsList{
			StateTransitions: msg.StateTransitions,
		})
	}
}
