package v1alpha1

import (
	meta "github.com/fdymylja/tmos/runtime/meta"
	module "github.com/fdymylja/tmos/runtime/module"
)

func (x *Params) StateObject() {}

func (x *Params) New() meta.StateObject {
	return new(Params)
}

type ParamsClient interface {
	Get() (*Params, error)
	Create(params *Params) error
	Delete(params *Params) error
	Update(params *Params) error
}
type paramsClient struct {
	client module.Client
}

func (x *paramsClient) Get() (*Params, error) {
	_spfGenO := new(Params)
	_spfGenErr := x.client.Get(meta.SingletonID, _spfGenO)
	if _spfGenErr != nil {
		return nil, _spfGenErr
	}
	return _spfGenO, nil
}
func (x *Role) StateObject() {}

func (x *Role) New() meta.StateObject {
	return new(Role)
}

type RoleClient interface {
	Get(id string) (*Role, error)
	Create(role *Role) error
	Delete(role *Role) error
	Update(role *Role) error
}
type roleClient struct {
	client module.Client
}

func (x *roleClient) Get(id string) (*Role, error) {
	_spfGenO := new(Role)
	_spfGenID := meta.NewStringID(id)
	_spfGenErr := x.client.Get(_spfGenID, _spfGenO)
	if _spfGenErr != nil {
		return nil, _spfGenErr
	}
	return _spfGenO, nil
}
func (x *RoleBinding) StateObject() {}

func (x *RoleBinding) New() meta.StateObject {
	return new(RoleBinding)
}

type RoleBindingClient interface {
	Get(subject string) (*RoleBinding, error)
	Create(roleBinding *RoleBinding) error
	Delete(roleBinding *RoleBinding) error
	Update(roleBinding *RoleBinding) error
}
type roleBindingClient struct {
	client module.Client
}

func (x *roleBindingClient) Get(subject string) (*RoleBinding, error) {
	_spfGenO := new(RoleBinding)
	_spfGenID := meta.NewStringID(subject)
	_spfGenErr := x.client.Get(_spfGenID, _spfGenO)
	if _spfGenErr != nil {
		return nil, _spfGenErr
	}
	return _spfGenO, nil
}
func (x *MsgCreateRole) StateTransition() {}
func (x *MsgCreateRole) New() meta.StateTransition {
	return new(MsgCreateRole)
}
func (x *MsgBindRole) StateTransition() {}
func (x *MsgBindRole) New() meta.StateTransition {
	return new(MsgBindRole)
}
