package v1alpha1

import (
	meta "github.com/fdymylja/tmos/core/meta"
	client "github.com/fdymylja/tmos/runtime/client"
	schema "github.com/fdymylja/tmos/runtime/orm/schema"
)

func (x *MsgSendCoins) APIDefinition() *meta.APIDefinition {
	return &meta.APIDefinition{
		Group:   "tmos.x.bank.v1alpha1",
		Kind:    "MsgSendCoins",
		ApiType: meta.APIType_StateTransition,
	}
}

func (x *MsgSendCoins) NewStateTransition() meta.StateTransition {
	return new(MsgSendCoins)
}

func (x *MsgSetBalance) APIDefinition() *meta.APIDefinition {
	return &meta.APIDefinition{
		Group:   "tmos.x.bank.v1alpha1",
		Kind:    "MsgSetBalance",
		ApiType: meta.APIType_StateTransition,
	}
}

func (x *MsgSetBalance) NewStateTransition() meta.StateTransition {
	return new(MsgSetBalance)
}

func (x *Balance) APIDefinition() *meta.APIDefinition {
	return &meta.APIDefinition{
		Group:   "tmos.x.bank.v1alpha1",
		Kind:    "Balance",
		ApiType: meta.APIType_StateObject,
	}
}

func (x *Balance) NewStateObject() meta.StateObject {
	return new(Balance)
}

type BalanceClient interface {
	Get(address string, opts ...client.GetOption) (*Balance, error)
	List(opts ...client.ListOption) (BalanceIterator, error)
	Create(balance *Balance, opts ...client.CreateOption) error
	Delete(balance *Balance, opts ...client.DeleteOption) error
	Update(balance *Balance, opts ...client.UpdateOption) error
}

type balanceClient struct {
	client client.Client
}

func (x *balanceClient) Get(address string, opts ...client.GetOption) (*Balance, error) {
	_spfGenO := new(Balance)
	_spfGenID := meta.NewStringID(address)
	_spfGenErr := x.client.Get(_spfGenID, _spfGenO, opts...)
	if _spfGenErr != nil {
		return nil, _spfGenErr
	}
	return _spfGenO, nil
}

func (x *balanceClient) List(opts ...client.ListOption) (BalanceIterator, error) {
	iter, err := x.client.List(new(Balance), opts...)
	if err != nil {
		return nil, err
	}
	return &balanceIterator{iter: iter}, nil
}

func (x *balanceClient) Create(balance *Balance, opts ...client.CreateOption) error {
	return x.client.Create(balance, opts...)
}

func (x *balanceClient) Delete(balance *Balance, opts ...client.DeleteOption) error {
	return x.client.Delete(balance, opts...)
}

func (x *balanceClient) Update(balance *Balance, opts ...client.UpdateOption) error {
	return x.client.Update(balance, opts...)
}

type BalanceIterator interface {
	Get() (*Balance, error)
	Valid() bool
	Next()
}

type balanceIterator struct {
	iter client.ObjectIterator
}

func (x *balanceIterator) Get() (*Balance, error) {
	obj := new(Balance)
	err := x.iter.Get(obj)
	return obj, err
}
func (x *balanceIterator) Valid() bool {
	return x.iter.Valid()
}

func (x *balanceIterator) Next() {
	x.iter.Next()
}

var BalanceSchema = &schema.Definition{
	PrimaryKey: "address",
}

type ClientSet interface {
	Balances() BalanceClient
	ExecMsgSendCoins(msg *MsgSendCoins) error
	ExecMsgSetBalance(msg *MsgSetBalance) error
}

func NewClientSet(client client.Client) ClientSet {
	return &clientSet{
		client:        client,
		balanceClient: &balanceClient{client: client},
	}
}

type clientSet struct {
	client client.Client
	// balanceClient is the client used to interact with Balance
	balanceClient BalanceClient
}

func (x *clientSet) Balances() BalanceClient {
	return x.balanceClient
}

func (x *clientSet) ExecMsgSendCoins(msg *MsgSendCoins) error {
	return x.client.Deliver(msg)
}

func (x *clientSet) ExecMsgSetBalance(msg *MsgSetBalance) error {
	return x.client.Deliver(msg)
}
