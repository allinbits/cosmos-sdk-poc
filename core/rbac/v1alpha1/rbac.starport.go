package v1alpha1

import (
	client "github.com/fdymylja/tmos/runtime/client"
	meta "github.com/fdymylja/tmos/runtime/meta"
	schema "github.com/fdymylja/tmos/runtime/orm/schema"
)

func (x *Params) StateObject() {}

func (x *Params) New() meta.StateObject {
	return new(Params)
}

type ParamsClient interface {
	Get(opts ...client.GetOption) (*Params, error)
	Create(params *Params, opts ...client.CreateOption) error
	Delete(params *Params, opts ...client.DeleteOption) error
	Update(params *Params, opts ...client.UpdateOption) error
}

type paramsClient struct {
	client client.RuntimeClient
}

func (x *paramsClient) Get(opts ...client.GetOption) (*Params, error) {
	_spfGenO := new(Params)
	_spfGenErr := x.client.Get(meta.SingletonID, _spfGenO, opts...)
	if _spfGenErr != nil {
		return nil, _spfGenErr
	}
	return _spfGenO, nil
}

func (x *paramsClient) Create(params *Params, opts ...client.CreateOption) error {
	return x.client.Create(params, opts...)
}

func (x *paramsClient) Delete(params *Params, opts ...client.DeleteOption) error {
	return x.client.Delete(params, opts...)
}

func (x *paramsClient) Update(params *Params, opts ...client.UpdateOption) error {
	return x.client.Update(params, opts...)
}

func (x *Role) StateObject() {}

func (x *Role) New() meta.StateObject {
	return new(Role)
}

type RoleClient interface {
	Get(id string, opts ...client.GetOption) (*Role, error)
	Create(role *Role, opts ...client.CreateOption) error
	Delete(role *Role, opts ...client.DeleteOption) error
	Update(role *Role, opts ...client.UpdateOption) error
}

type roleClient struct {
	client client.RuntimeClient
}

func (x *roleClient) Get(id string, opts ...client.GetOption) (*Role, error) {
	_spfGenO := new(Role)
	_spfGenID := meta.NewStringID(id)
	_spfGenErr := x.client.Get(_spfGenID, _spfGenO, opts...)
	if _spfGenErr != nil {
		return nil, _spfGenErr
	}
	return _spfGenO, nil
}

func (x *roleClient) Create(role *Role, opts ...client.CreateOption) error {
	return x.client.Create(role, opts...)
}

func (x *roleClient) Delete(role *Role, opts ...client.DeleteOption) error {
	return x.client.Delete(role, opts...)
}

func (x *roleClient) Update(role *Role, opts ...client.UpdateOption) error {
	return x.client.Update(role, opts...)
}

func (x *RoleBinding) StateObject() {}

func (x *RoleBinding) New() meta.StateObject {
	return new(RoleBinding)
}

type RoleBindingClient interface {
	Get(subject string, opts ...client.GetOption) (*RoleBinding, error)
	Create(roleBinding *RoleBinding, opts ...client.CreateOption) error
	Delete(roleBinding *RoleBinding, opts ...client.DeleteOption) error
	Update(roleBinding *RoleBinding, opts ...client.UpdateOption) error
}

type roleBindingClient struct {
	client client.RuntimeClient
}

func (x *roleBindingClient) Get(subject string, opts ...client.GetOption) (*RoleBinding, error) {
	_spfGenO := new(RoleBinding)
	_spfGenID := meta.NewStringID(subject)
	_spfGenErr := x.client.Get(_spfGenID, _spfGenO, opts...)
	if _spfGenErr != nil {
		return nil, _spfGenErr
	}
	return _spfGenO, nil
}

func (x *roleBindingClient) Create(roleBinding *RoleBinding, opts ...client.CreateOption) error {
	return x.client.Create(roleBinding, opts...)
}

func (x *roleBindingClient) Delete(roleBinding *RoleBinding, opts ...client.DeleteOption) error {
	return x.client.Delete(roleBinding, opts...)
}

func (x *roleBindingClient) Update(roleBinding *RoleBinding, opts ...client.UpdateOption) error {
	return x.client.Update(roleBinding, opts...)
}

func (x *MsgCreateRole) StateTransition() {}

func (x *MsgCreateRole) New() meta.StateTransition {
	return new(MsgCreateRole)
}

func (x *MsgBindRole) StateTransition() {}

func (x *MsgBindRole) New() meta.StateTransition {
	return new(MsgBindRole)
}

var ParamsSchema = schema.Definition{
	Meta: meta.Meta{
		APIGroup: "tmos.rbac.v1alpha1",
		APIKind:  "Params",
	},
	Singleton: true,
}

var RoleSchema = schema.Definition{
	Meta: meta.Meta{
		APIGroup: "tmos.rbac.v1alpha1",
		APIKind:  "Role",
	},
	PrimaryKey: "id",
}

var RoleBindingSchema = schema.Definition{
	Meta: meta.Meta{
		APIGroup: "tmos.rbac.v1alpha1",
		APIKind:  "RoleBinding",
	},
	PrimaryKey:    "subject",
	SecondaryKeys: []string{"roleRef"},
}

type ClientSet interface {
	Params() ParamsClient
	Roles() RoleClient
	RoleBindings() RoleBindingClient
	ExecMsgCreateRole(msg *MsgCreateRole) error
	ExecMsgBindRole(msg *MsgBindRole) error
}

func NewClientSet(client client.RuntimeClient) ClientSet {
	return &clientSet{
		client:            client,
		paramsClient:      &paramsClient{client: client},
		roleClient:        &roleClient{client: client},
		roleBindingClient: &roleBindingClient{client: client},
	}
}

type clientSet struct {
	client client.RuntimeClient
	// paramsClient is the client used to interact with Params
	paramsClient ParamsClient
	// roleClient is the client used to interact with Role
	roleClient RoleClient
	// roleBindingClient is the client used to interact with RoleBinding
	roleBindingClient RoleBindingClient
}

func (x *clientSet) Params() ParamsClient {
	return x.paramsClient
}

func (x *clientSet) Roles() RoleClient {
	return x.roleClient
}

func (x *clientSet) RoleBindings() RoleBindingClient {
	return x.roleBindingClient
}

func (x *clientSet) ExecMsgCreateRole(msg *MsgCreateRole) error {
	return x.client.Deliver(msg)
}

func (x *clientSet) ExecMsgBindRole(msg *MsgBindRole) error {
	return x.client.Deliver(msg)
}
