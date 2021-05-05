package rbac

import (
	"github.com/fdymylja/tmos/module/rbac/v1alpha1"
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
