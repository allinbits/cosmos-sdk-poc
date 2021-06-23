package v1alpha1

import (
	meta "github.com/fdymylja/tmos/core/meta"
	client "github.com/fdymylja/tmos/runtime/client"
	schema "github.com/fdymylja/tmos/runtime/orm/schema"
)

func (x *Params) APIDefinition() *meta.APIDefinition {
	return &meta.APIDefinition{
		Group:   "tmos.core.rbac.v1alpha1",
		Kind:    "Params",
		ApiType: meta.APIType_StateObject,
	}
}

func (x *Params) NewStateObject() meta.StateObject {
	return new(Params)
}

type ParamsClient interface {
	Get(opts ...client.GetOption) (*Params, error)
	Create(params *Params, opts ...client.CreateOption) error
	Delete(params *Params, opts ...client.DeleteOption) error
	Update(params *Params, opts ...client.UpdateOption) error
}

type paramsClient struct {
	client client.Client
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

func (x *Role) APIDefinition() *meta.APIDefinition {
	return &meta.APIDefinition{
		Group:   "tmos.core.rbac.v1alpha1",
		Kind:    "Role",
		ApiType: meta.APIType_StateObject,
	}
}

func (x *Role) NewStateObject() meta.StateObject {
	return new(Role)
}

type RoleClient interface {
	Get(id string, opts ...client.GetOption) (*Role, error)
	List(opts ...client.ListOption) (RoleIterator, error)
	Create(role *Role, opts ...client.CreateOption) error
	Delete(role *Role, opts ...client.DeleteOption) error
	Update(role *Role, opts ...client.UpdateOption) error
}

type roleClient struct {
	client client.Client
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

func (x *roleClient) List(opts ...client.ListOption) (RoleIterator, error) {
	iter, err := x.client.List(new(Role), opts...)
	if err != nil {
		return nil, err
	}
	return &roleIterator{iter: iter}, nil
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

type RoleIterator interface {
	Get() (*Role, error)
	Valid() bool
	Next()
}

type roleIterator struct {
	iter client.ObjectIterator
}

func (x *roleIterator) Get() (*Role, error) {
	obj := new(Role)
	err := x.iter.Get(obj)
	return obj, err
}
func (x *roleIterator) Valid() bool {
	return x.iter.Valid()
}

func (x *roleIterator) Next() {
	x.iter.Next()
}

func (x *RoleBinding) APIDefinition() *meta.APIDefinition {
	return &meta.APIDefinition{
		Group:   "tmos.core.rbac.v1alpha1",
		Kind:    "RoleBinding",
		ApiType: meta.APIType_StateObject,
	}
}

func (x *RoleBinding) NewStateObject() meta.StateObject {
	return new(RoleBinding)
}

type RoleBindingClient interface {
	Get(subject string, opts ...client.GetOption) (*RoleBinding, error)
	List(opts ...client.ListOption) (RoleBindingIterator, error)
	Create(roleBinding *RoleBinding, opts ...client.CreateOption) error
	Delete(roleBinding *RoleBinding, opts ...client.DeleteOption) error
	Update(roleBinding *RoleBinding, opts ...client.UpdateOption) error
}

type roleBindingClient struct {
	client client.Client
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

func (x *roleBindingClient) List(opts ...client.ListOption) (RoleBindingIterator, error) {
	iter, err := x.client.List(new(RoleBinding), opts...)
	if err != nil {
		return nil, err
	}
	return &roleBindingIterator{iter: iter}, nil
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

type RoleBindingIterator interface {
	Get() (*RoleBinding, error)
	Valid() bool
	Next()
}

type roleBindingIterator struct {
	iter client.ObjectIterator
}

func (x *roleBindingIterator) Get() (*RoleBinding, error) {
	obj := new(RoleBinding)
	err := x.iter.Get(obj)
	return obj, err
}
func (x *roleBindingIterator) Valid() bool {
	return x.iter.Valid()
}

func (x *roleBindingIterator) Next() {
	x.iter.Next()
}

func (x *MsgCreateRole) APIDefinition() *meta.APIDefinition {
	return &meta.APIDefinition{
		Group:   "tmos.core.rbac.v1alpha1",
		Kind:    "MsgCreateRole",
		ApiType: meta.APIType_StateTransition,
	}
}

func (x *MsgCreateRole) NewStateTransition() meta.StateTransition {
	return new(MsgCreateRole)
}

func (x *MsgBindRole) APIDefinition() *meta.APIDefinition {
	return &meta.APIDefinition{
		Group:   "tmos.core.rbac.v1alpha1",
		Kind:    "MsgBindRole",
		ApiType: meta.APIType_StateTransition,
	}
}

func (x *MsgBindRole) NewStateTransition() meta.StateTransition {
	return new(MsgBindRole)
}

var ParamsSchema = &schema.Definition{
	Singleton: true,
}

var RoleSchema = &schema.Definition{
	PrimaryKey: "id",
}

var RoleBindingSchema = &schema.Definition{
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

func NewClientSet(client client.Client) ClientSet {
	return &clientSet{
		client:            client,
		paramsClient:      &paramsClient{client: client},
		roleClient:        &roleClient{client: client},
		roleBindingClient: &roleBindingClient{client: client},
	}
}

type clientSet struct {
	client client.Client
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
