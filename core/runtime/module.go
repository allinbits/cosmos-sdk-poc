package runtime

import (
	"github.com/fdymylja/tmos/core/runtime/v1alpha1"
	"github.com/fdymylja/tmos/runtime/authentication/user"
	"github.com/fdymylja/tmos/runtime/module"
	"github.com/fdymylja/tmos/runtime/statetransition"
)

func NewModule() Module { return Module{} }

type Module struct {
}

func (m Module) Initialize(client module.Client) module.Descriptor {
	return module.NewDescriptorBuilder().
		Named(user.Runtime).
		OwnsStateObject(&v1alpha1.StateObjectsList{}, v1alpha1.StateObjectsListSchema).
		OwnsStateObject(&v1alpha1.StateTransitionsList{}, v1alpha1.StateTransitionsListSchema).
		HandlesStateTransition(&v1alpha1.CreateStateTransitionsList{}, newCreateStateTransitionsController(client), false).
		HandlesStateTransition(&v1alpha1.CreateStateObjectsList{}, newCreateStateObjectsController(client), false).Build()
}

func newCreateStateObjectsController(client module.Client) statetransition.ExecutionHandlerFunc {
	return func(req statetransition.ExecutionRequest) (resp statetransition.ExecutionResponse, err error) {
		msg := req.Transition.(*v1alpha1.CreateStateObjectsList)
		return resp, client.Create(&v1alpha1.StateObjectsList{StateObjects: msg.StateObjects})
	}
}

func newCreateStateTransitionsController(client module.Client) statetransition.ExecutionHandlerFunc {
	return func(req statetransition.ExecutionRequest) (resp statetransition.ExecutionResponse, err error) {
		msg := req.Transition.(*v1alpha1.CreateStateTransitionsList)
		return resp, client.Create(&v1alpha1.StateTransitionsList{
			StateTransitions: msg.StateTransitions,
		})
	}
}
