package v1alpha1

import (
	client "github.com/fdymylja/tmos/runtime/client"
	meta "github.com/fdymylja/tmos/runtime/meta"
	schema "github.com/fdymylja/tmos/runtime/orm/schema"
)

func (x *InvariantHandler) StateObject() {}

func (x *InvariantHandler) New() meta.StateObject {
	return new(InvariantHandler)
}

type InvariantHandlerClient interface {
	Get(stateTransition string, opts ...client.GetOption) (*InvariantHandler, error)
	Create(invariantHandler *InvariantHandler, opts ...client.CreateOption) error
	Delete(invariantHandler *InvariantHandler, opts ...client.DeleteOption) error
	Update(invariantHandler *InvariantHandler, opts ...client.UpdateOption) error
}

type invariantHandlerClient struct {
	client client.RuntimeClient
}

func (x *invariantHandlerClient) Get(stateTransition string, opts ...client.GetOption) (*InvariantHandler, error) {
	_spfGenO := new(InvariantHandler)
	_spfGenID := meta.NewStringID(stateTransition)
	_spfGenErr := x.client.Get(_spfGenID, _spfGenO, opts...)
	if _spfGenErr != nil {
		return nil, _spfGenErr
	}
	return _spfGenO, nil
}

func (x *invariantHandlerClient) Create(invariantHandler *InvariantHandler, opts ...client.CreateOption) error {
	return x.client.Create(invariantHandler, opts...)
}

func (x *invariantHandlerClient) Delete(invariantHandler *InvariantHandler, opts ...client.DeleteOption) error {
	return x.client.Delete(invariantHandler, opts...)
}

func (x *invariantHandlerClient) Update(invariantHandler *InvariantHandler, opts ...client.UpdateOption) error {
	return x.client.Update(invariantHandler, opts...)
}

var InvariantHandlerSchema = schema.Definition{
	Meta: meta.Meta{
		APIGroup: "tmos.x.crisis.v1alpha1",
		APIKind:  "InvariantHandler",
	},
	PrimaryKey:    "stateTransition",
	SecondaryKeys: []string{"module", "route"},
}

type ClientSet interface {
	InvariantHandlers() InvariantHandlerClient
}

func NewClientSet(client client.RuntimeClient) ClientSet {
	return &clientSet{
		client:                 client,
		invariantHandlerClient: &invariantHandlerClient{client: client},
	}
}

type clientSet struct {
	client client.RuntimeClient
	// invariantHandlerClient is the client used to interact with InvariantHandler
	invariantHandlerClient InvariantHandlerClient
}

func (x *clientSet) InvariantHandlers() InvariantHandlerClient {
	return x.invariantHandlerClient
}
