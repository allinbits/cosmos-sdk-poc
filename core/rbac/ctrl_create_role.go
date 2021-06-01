package rbac

import (
	"errors"
	"fmt"

	"github.com/fdymylja/tmos/core/rbac/v1alpha1"
	runtimev1alpha1 "github.com/fdymylja/tmos/core/runtime/v1alpha1"
	rterr "github.com/fdymylja/tmos/runtime/errors"
	"github.com/fdymylja/tmos/runtime/meta"
	"github.com/fdymylja/tmos/runtime/module"
	"github.com/fdymylja/tmos/runtime/statetransition"
	"github.com/scylladb/go-set/strset"
)

func NewCreateRoleHandler(client module.Client) statetransition.ExecutionHandler {
	return CreateRoleHandler{
		client: client,
	}
}

type CreateRoleHandler struct {
	client module.Client
}

func (c CreateRoleHandler) Exec(req statetransition.ExecutionRequest) (statetransition.ExecutionResponse, error) {
	msg := req.Transition.(*v1alpha1.MsgCreateRole)
	return statetransition.ExecutionResponse{}, c.client.Create(msg.NewRole)
}

func NewCreateRoleAdmissionHandler(client module.Client) CreateRoleAdmissionHandler {
	return CreateRoleAdmissionHandler{
		client:   client,
		rtClient: runtimev1alpha1.NewClientSet(client),
	}
}

type CreateRoleAdmissionHandler struct {
	client   module.Client
	rtClient runtimev1alpha1.ClientSet
}

func (c CreateRoleAdmissionHandler) Validate(req statetransition.AdmissionRequest) error {
	msg := req.Transition.(*v1alpha1.MsgCreateRole)
	if msg.NewRole == nil {
		return fmt.Errorf("new role is nil")
	}
	role := msg.NewRole
	// NOTE we check only for role id and nothing else
	// as we might want to create a role which has access to nothing
	if role.Id == "" {
		return fmt.Errorf("no role id defined")
	}
	if err := c.roleNotExist(msg.NewRole.Id); err != nil {
		return err
	}
	if err := c.verifyStateObjects(msg.NewRole); err != nil {
		return err
	}
	if err := c.verifyStateTransitions(msg.NewRole); err != nil {
		return err
	}
	// pass
	return nil
}

func (c CreateRoleAdmissionHandler) roleNotExist(id string) error {
	err := c.client.Get(meta.NewStringID(id), new(v1alpha1.Role))
	switch {
	case err == nil:
		return fmt.Errorf("role %s already exists", id)
	case errors.Is(err, rterr.ErrNotFound):
		return nil
	default:
		return err
	}
}

func (c CreateRoleAdmissionHandler) verifyStateObjects(role *v1alpha1.Role) error {
	stateObjects, err := c.rtClient.StateObjectsList().Get()
	if err != nil {
		return err
	}
	set := strset.New(stateObjects.StateObjects...)
	if len(role.Gets) != 0 && !set.Has(role.Gets...) {
		return fmt.Errorf("unknown state object in get types %#v", role.Gets)
	}
	if len(role.Lists) != 0 && !set.Has(role.Lists...) {
		return fmt.Errorf("unknown state object in list types %#v", role.Lists)
	}
	if len(role.Creates) != 0 && !set.Has(role.Creates...) {
		return fmt.Errorf("unknown state object in create types %#v", role.Creates)
	}
	if len(role.Updates) != 0 && !set.Has(role.Updates...) {
		return fmt.Errorf("unkown state object in update types %#v", role.Updates)
	}
	if len(role.Deletes) != 0 && !set.Has(role.Deletes...) {
		return fmt.Errorf("unknown state object in delete types %#v", role.Deletes)
	}
	return nil
}

func (c CreateRoleAdmissionHandler) verifyStateTransitions(role *v1alpha1.Role) error {
	stateTransitions, err := c.rtClient.StateTransitionsList().Get()
	if err != nil {
		return err
	}
	set := strset.New(stateTransitions.StateTransitions...)
	if len(role.Delivers) != 0 && !set.Has(role.Delivers...) {
		return fmt.Errorf("unknown state transition types %#v", role.Delivers)
	}
	return nil
}
