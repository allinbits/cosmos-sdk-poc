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
		OwnsStateObject(&v1alpha1.ModuleDescriptors{}, v1alpha1.ModuleDescriptorsSchema).
		HandlesStateTransition(&v1alpha1.CreateModuleDescriptors{}, newCreateModuleDescriptorsHandler(client), false).
		Build()
}

func newCreateModuleDescriptorsHandler(client module.Client) statetransition.ExecutionHandlerFunc {
	return func(req statetransition.ExecutionRequest) (resp statetransition.ExecutionResponse, err error) {
		msg := req.Transition.(*v1alpha1.CreateModuleDescriptors)
		return resp, client.Create(&v1alpha1.ModuleDescriptors{Modules: msg.Modules})
	}
}
