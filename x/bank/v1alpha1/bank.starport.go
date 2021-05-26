package v1alpha1

import (
	client "github.com/fdymylja/tmos/runtime/client"
	meta "github.com/fdymylja/tmos/runtime/meta"
	schema "github.com/fdymylja/tmos/runtime/orm/schema"
)

func (x *MsgSendCoins) StateTransition() {}

func (x *MsgSendCoins) New() meta.StateTransition {
	return new(MsgSendCoins)
}

func (x *MsgSetBalance) StateTransition() {}

func (x *MsgSetBalance) New() meta.StateTransition {
	return new(MsgSetBalance)
}

func (x *Balance) StateObject() {}

func (x *Balance) New() meta.StateObject {
	return new(Balance)
}

type BalanceClient interface {
	Get(address string, opts ...client.GetOption) (*Balance, error)
	Create(balance *Balance, opts ...client.CreateOption) error
	Delete(balance *Balance, opts ...client.DeleteOption) error
	Update(balance *Balance, opts ...client.UpdateOption) error
}

type balanceClient struct {
	client client.RuntimeClient
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

func (x *balanceClient) Create(balance *Balance, opts ...client.CreateOption) error {
	return x.client.Create(balance, opts...)
}

func (x *balanceClient) Delete(balance *Balance, opts ...client.DeleteOption) error {
	return x.client.Delete(balance, opts...)
}

func (x *balanceClient) Update(balance *Balance, opts ...client.UpdateOption) error {
	return x.client.Update(balance, opts...)
}

var BalanceSchema = schema.Definition{
	Meta: meta.Meta{
		APIGroup: "tmos.x.bank.v1alpha1",
		APIKind:  "Balance",
	},
	PrimaryKey: "address",
}

type ClientSet interface {
	Balances() BalanceClient
	ExecMsgSendCoins(msg *MsgSendCoins) error
	ExecMsgSetBalance(msg *MsgSetBalance) error
}

func NewClientSet(client client.RuntimeClient) ClientSet {
	return clientSet{
		client:        client,
		balanceClient: &balanceClient{client: client},
	}
}

type clientSet struct {
	client client.RuntimeClient
	// balanceClient is the client used to interact with Balance
	balanceClient BalanceClient
}

func (x clientSet) Balances() BalanceClient {
	return x.balanceClient
}

func (x clientSet) ExecMsgSendCoins(msg *MsgSendCoins) error {
	return x.client.Deliver(msg)
}

func (x clientSet) ExecMsgSetBalance(msg *MsgSetBalance) error {
	return x.client.Deliver(msg)
}
