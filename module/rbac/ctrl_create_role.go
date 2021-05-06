package rbac

import (
	"errors"
	"fmt"

	"github.com/fdymylja/tmos/module/rbac/v1alpha1"
	runtimev1alpha1 "github.com/fdymylja/tmos/module/runtime/v1alpha1"
	"github.com/fdymylja/tmos/runtime"
	"github.com/fdymylja/tmos/runtime/controller"
	"github.com/fdymylja/tmos/runtime/meta"
	"github.com/fdymylja/tmos/runtime/module"
	"github.com/scylladb/go-set/strset"
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
	return CreateRoleController{
		client: client,
	}
}

type CreateRoleController struct {
	client module.Client
}

func (c CreateRoleController) Deliver(req controller.StateTransitionRequest) (controller.StateTransitionResponse, error) {
	msg := req.Transition.(*v1alpha1.MsgCreateRole)
	return controller.StateTransitionResponse{}, c.client.Create(msg.NewRole)
}

func NewCreateRoleAdmissionController(client module.Client) CreateRoleAdmissionController {
	return CreateRoleAdmissionController{
		client:   client,
		rtClient: runtimev1alpha1.NewClient(client),
	}
}

type CreateRoleAdmissionController struct {
	client   module.Client
	rtClient *runtimev1alpha1.Client
}

func (c CreateRoleAdmissionController) Validate(req controller.AdmissionRequest) (controller.AdmissionResponse, error) {
	msg := req.Transition.(*v1alpha1.MsgCreateRole)
	if msg.NewRole == nil {
		return controller.AdmissionResponse{}, fmt.Errorf("new role is nil")
	}
	role := msg.NewRole
	// NOTE we check only for role id and nothing else
	// as we might want to create a role which has access to nothing
	if role.Id == "" {
		return controller.AdmissionResponse{}, fmt.Errorf("no role id defined")
	}
	if err := c.roleNotExist(msg.NewRole.Id); err != nil {
		return controller.AdmissionResponse{}, err
	}
	if err := c.verifyStateObjects(msg.NewRole); err != nil {
		return controller.AdmissionResponse{}, err
	}
	if err := c.verifyStateTransitions(msg.NewRole); err != nil {
		return controller.AdmissionResponse{}, err
	}
	// pass
	return controller.AdmissionResponse{}, nil
}

func (c CreateRoleAdmissionController) roleNotExist(id string) error {
	err := c.client.Get(meta.NewStringID(id), new(v1alpha1.Role))
	switch {
	case err == nil:
		return fmt.Errorf("role %s already exists", id)
	case errors.Is(err, runtime.ErrNotFound):
		return nil
	default:
		return err
	}
}

func (c CreateRoleAdmissionController) verifyStateObjects(role *v1alpha1.Role) error {
	stateObjects, err := c.rtClient.GetStateObjectsList()
	if err != nil {
		return err
	}
	set := strset.New(stateObjects.StateObjects...)
	if !set.Has(role.Gets...) {
		return fmt.Errorf("unknown state object in get types %#v", role.Gets)
	}
	if !set.Has(role.Lists...) {
		return fmt.Errorf("unknown state object in list types %#v", role.Lists)
	}
	if !set.Has(role.Creates...) {
		return fmt.Errorf("unknown state object in create types %#v", role.Creates)
	}
	if !set.Has(role.Updates...) {
		return fmt.Errorf("unkown state object in update types %#v", role.Updates)
	}
	if !set.Has(role.Deletes...) {
		return fmt.Errorf("unknown state object in delete types %#v", role.Deletes)
	}
	return nil
}

func (c CreateRoleAdmissionController) verifyStateTransitions(role *v1alpha1.Role) error {
	stateTransitions, err := c.rtClient.GetStateTransitionsList()
	if err != nil {
		return err
	}
	set := strset.New(stateTransitions.StateTransitions...)
	if !set.Has(role.Delivers...) {
		return fmt.Errorf("unknown state transition types %#v", role.Delivers)
	}
	return nil
}
