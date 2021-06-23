package v1alpha1

import (
	meta "github.com/fdymylja/tmos/core/meta"
	client "github.com/fdymylja/tmos/runtime/client"
	schema "github.com/fdymylja/tmos/runtime/orm/schema"
)

func (x *MsgCreateAccount) APIDefinition() *meta.APIDefinition {
	return &meta.APIDefinition{
		Group:   "tmos.x.authn.v1alpha1",
		Kind:    "MsgCreateAccount",
		ApiType: meta.APIType_StateTransition,
	}
}

func (x *MsgCreateAccount) NewStateTransition() meta.StateTransition {
	return new(MsgCreateAccount)
}

func (x *MsgUpdateAccount) APIDefinition() *meta.APIDefinition {
	return &meta.APIDefinition{
		Group:   "tmos.x.authn.v1alpha1",
		Kind:    "MsgUpdateAccount",
		ApiType: meta.APIType_StateTransition,
	}
}

func (x *MsgUpdateAccount) NewStateTransition() meta.StateTransition {
	return new(MsgUpdateAccount)
}

func (x *MsgDeleteAccount) APIDefinition() *meta.APIDefinition {
	return &meta.APIDefinition{
		Group:   "tmos.x.authn.v1alpha1",
		Kind:    "MsgDeleteAccount",
		ApiType: meta.APIType_StateTransition,
	}
}

func (x *MsgDeleteAccount) NewStateTransition() meta.StateTransition {
	return new(MsgDeleteAccount)
}

func (x *Account) APIDefinition() *meta.APIDefinition {
	return &meta.APIDefinition{
		Group:   "tmos.x.authn.v1alpha1",
		Kind:    "Account",
		ApiType: meta.APIType_StateObject,
	}
}

func (x *Account) NewStateObject() meta.StateObject {
	return new(Account)
}

type AccountClient interface {
	Get(address string, opts ...client.GetOption) (*Account, error)
	List(opts ...client.ListOption) (AccountIterator, error)
	Create(account *Account, opts ...client.CreateOption) error
	Delete(account *Account, opts ...client.DeleteOption) error
	Update(account *Account, opts ...client.UpdateOption) error
}

type accountClient struct {
	client client.RuntimeClient
}

func (x *accountClient) Get(address string, opts ...client.GetOption) (*Account, error) {
	_spfGenO := new(Account)
	_spfGenID := meta.NewStringID(address)
	_spfGenErr := x.client.Get(_spfGenID, _spfGenO, opts...)
	if _spfGenErr != nil {
		return nil, _spfGenErr
	}
	return _spfGenO, nil
}

func (x *accountClient) List(opts ...client.ListOption) (AccountIterator, error) {
	iter, err := x.client.List(new(Account), opts...)
	if err != nil {
		return nil, err
	}
	return &accountIterator{iter: iter}, nil
}

func (x *accountClient) Create(account *Account, opts ...client.CreateOption) error {
	return x.client.Create(account, opts...)
}

func (x *accountClient) Delete(account *Account, opts ...client.DeleteOption) error {
	return x.client.Delete(account, opts...)
}

func (x *accountClient) Update(account *Account, opts ...client.UpdateOption) error {
	return x.client.Update(account, opts...)
}

type AccountIterator interface {
	Get() (*Account, error)
	Valid() bool
	Next()
}

type accountIterator struct {
	iter client.ObjectIterator
}

func (x *accountIterator) Get() (*Account, error) {
	obj := new(Account)
	err := x.iter.Get(obj)
	return obj, err
}
func (x *accountIterator) Valid() bool {
	return x.iter.Valid()
}

func (x *accountIterator) Next() {
	x.iter.Next()
}

func (x *CurrentAccountNumber) APIDefinition() *meta.APIDefinition {
	return &meta.APIDefinition{
		Group:   "tmos.x.authn.v1alpha1",
		Kind:    "CurrentAccountNumber",
		ApiType: meta.APIType_StateObject,
	}
}

func (x *CurrentAccountNumber) NewStateObject() meta.StateObject {
	return new(CurrentAccountNumber)
}

type CurrentAccountNumberClient interface {
	Get(opts ...client.GetOption) (*CurrentAccountNumber, error)
	Create(currentAccountNumber *CurrentAccountNumber, opts ...client.CreateOption) error
	Delete(currentAccountNumber *CurrentAccountNumber, opts ...client.DeleteOption) error
	Update(currentAccountNumber *CurrentAccountNumber, opts ...client.UpdateOption) error
}

type currentAccountNumberClient struct {
	client client.RuntimeClient
}

func (x *currentAccountNumberClient) Get(opts ...client.GetOption) (*CurrentAccountNumber, error) {
	_spfGenO := new(CurrentAccountNumber)
	_spfGenErr := x.client.Get(meta.SingletonID, _spfGenO, opts...)
	if _spfGenErr != nil {
		return nil, _spfGenErr
	}
	return _spfGenO, nil
}

func (x *currentAccountNumberClient) Create(currentAccountNumber *CurrentAccountNumber, opts ...client.CreateOption) error {
	return x.client.Create(currentAccountNumber, opts...)
}

func (x *currentAccountNumberClient) Delete(currentAccountNumber *CurrentAccountNumber, opts ...client.DeleteOption) error {
	return x.client.Delete(currentAccountNumber, opts...)
}

func (x *currentAccountNumberClient) Update(currentAccountNumber *CurrentAccountNumber, opts ...client.UpdateOption) error {
	return x.client.Update(currentAccountNumber, opts...)
}

func (x *Params) APIDefinition() *meta.APIDefinition {
	return &meta.APIDefinition{
		Group:   "tmos.x.authn.v1alpha1",
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

var AccountSchema = schema.Definition{
	PrimaryKey:    "address",
	SecondaryKeys: []string{"accountNumber"},
}

var CurrentAccountNumberSchema = schema.Definition{
	Singleton: true,
}

var ParamsSchema = schema.Definition{
	Singleton: true,
}

type ClientSet interface {
	Accounts() AccountClient
	CurrentAccountNumber() CurrentAccountNumberClient
	Params() ParamsClient
	ExecMsgCreateAccount(msg *MsgCreateAccount) error
	ExecMsgUpdateAccount(msg *MsgUpdateAccount) error
	ExecMsgDeleteAccount(msg *MsgDeleteAccount) error
}

func NewClientSet(client client.RuntimeClient) ClientSet {
	return &clientSet{
		client:                     client,
		accountClient:              &accountClient{client: client},
		currentAccountNumberClient: &currentAccountNumberClient{client: client},
		paramsClient:               &paramsClient{client: client},
	}
}

type clientSet struct {
	client client.RuntimeClient
	// accountClient is the client used to interact with Account
	accountClient AccountClient
	// currentAccountNumberClient is the client used to interact with CurrentAccountNumber
	currentAccountNumberClient CurrentAccountNumberClient
	// paramsClient is the client used to interact with Params
	paramsClient ParamsClient
}

func (x *clientSet) Accounts() AccountClient {
	return x.accountClient
}

func (x *clientSet) CurrentAccountNumber() CurrentAccountNumberClient {
	return x.currentAccountNumberClient
}

func (x *clientSet) Params() ParamsClient {
	return x.paramsClient
}

func (x *clientSet) ExecMsgCreateAccount(msg *MsgCreateAccount) error {
	return x.client.Deliver(msg)
}

func (x *clientSet) ExecMsgUpdateAccount(msg *MsgUpdateAccount) error {
	return x.client.Deliver(msg)
}

func (x *clientSet) ExecMsgDeleteAccount(msg *MsgDeleteAccount) error {
	return x.client.Deliver(msg)
}
