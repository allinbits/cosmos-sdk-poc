package rbac

import (
	"fmt"

	"github.com/fdymylja/tmos/module/rbac/v1alpha1"
	runtime "github.com/fdymylja/tmos/module/runtime/v1alpha1"
	"github.com/fdymylja/tmos/runtime/controller"
	"github.com/fdymylja/tmos/runtime/module"
)

func NewBindRoleController(client module.Client) controller.StateTransition {
	return BindRoleController{client: client}
}

type BindRoleController struct {
	client module.Client
}

func (b BindRoleController) Deliver(req controller.StateTransitionRequest) (controller.StateTransitionResponse, error) {
	panic("implement me")
}

func NewCreateRoleController(client module.Client) controller.StateTransition {
	return CreateRoleController{client: client}
}

type CreateRoleController struct {
	client module.Client
}

func (c CreateRoleController) Deliver(req controller.StateTransitionRequest) (controller.StateTransitionResponse, error) {
	msg := req.Transition.(*v1alpha1.MsgCreateRole)
	return controller.StateTransitionResponse{}, c.client.Create(msg.NewRole)
}

type CreateRoleAdmissionController struct {
	client module.Client
}

func (c CreateRoleAdmissionController) Validate(req controller.AdmissionRequest) (controller.AdmissionResponse, error) {
	msg := req.Transition.(*v1alpha1.MsgCreateRole)
	if msg.NewRole == nil {
		return controller.AdmissionResponse{}, fmt.Errorf("new role is nil")
	}
	role := msg.NewRole
	if len(role.Verbs) == 0 {
		return controller.AdmissionResponse{}, fmt.Errorf("at least one verb required")
	}
	if len(role.Resources) == 0 {
		return controller.AdmissionResponse{}, fmt.Errorf("at least one resource required")
	}
	// iterate over verbs to check for unknowns
	for _, verb := range role.Verbs {
		if verb == runtime.Verb_Unknown {
			return controller.AdmissionResponse{}, fmt.Errorf("unknown verb not allowed")
		}
	}
	// check for resources existence

}
