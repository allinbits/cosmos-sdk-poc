package v1alpha1

import (
	client "github.com/fdymylja/tmos/runtime/client"
	meta "github.com/fdymylja/tmos/runtime/meta"
	schema "github.com/fdymylja/tmos/runtime/orm/schema"
)

func (x *StateObjectsList) StateObject() {}

func (x *StateObjectsList) New() meta.StateObject {
	return new(StateObjectsList)
}

type StateObjectsListClient interface {
	Get(opts ...client.GetOption) (*StateObjectsList, error)
	Create(stateObjectsList *StateObjectsList, opts ...client.CreateOption) error
	Delete(stateObjectsList *StateObjectsList, opts ...client.DeleteOption) error
	Update(stateObjectsList *StateObjectsList, opts ...client.UpdateOption) error
}

type stateObjectsListClient struct {
	client client.RuntimeClient
}

func (x *stateObjectsListClient) Get(opts ...client.GetOption) (*StateObjectsList, error) {
	_spfGenO := new(StateObjectsList)
	_spfGenErr := x.client.Get(meta.SingletonID, _spfGenO, opts...)
	if _spfGenErr != nil {
		return nil, _spfGenErr
	}
	return _spfGenO, nil
}

func (x *stateObjectsListClient) Create(stateObjectsList *StateObjectsList, opts ...client.CreateOption) error {
	return x.client.Create(stateObjectsList, opts...)
}

func (x *stateObjectsListClient) Delete(stateObjectsList *StateObjectsList, opts ...client.DeleteOption) error {
	return x.client.Delete(stateObjectsList, opts...)
}

func (x *stateObjectsListClient) Update(stateObjectsList *StateObjectsList, opts ...client.UpdateOption) error {
	return x.client.Update(stateObjectsList, opts...)
}

func (x *StateTransitionsList) StateObject() {}

func (x *StateTransitionsList) New() meta.StateObject {
	return new(StateTransitionsList)
}

type StateTransitionsListClient interface {
	Get(opts ...client.GetOption) (*StateTransitionsList, error)
	Create(stateTransitionsList *StateTransitionsList, opts ...client.CreateOption) error
	Delete(stateTransitionsList *StateTransitionsList, opts ...client.DeleteOption) error
	Update(stateTransitionsList *StateTransitionsList, opts ...client.UpdateOption) error
}

type stateTransitionsListClient struct {
	client client.RuntimeClient
}

func (x *stateTransitionsListClient) Get(opts ...client.GetOption) (*StateTransitionsList, error) {
	_spfGenO := new(StateTransitionsList)
	_spfGenErr := x.client.Get(meta.SingletonID, _spfGenO, opts...)
	if _spfGenErr != nil {
		return nil, _spfGenErr
	}
	return _spfGenO, nil
}

func (x *stateTransitionsListClient) Create(stateTransitionsList *StateTransitionsList, opts ...client.CreateOption) error {
	return x.client.Create(stateTransitionsList, opts...)
}

func (x *stateTransitionsListClient) Delete(stateTransitionsList *StateTransitionsList, opts ...client.DeleteOption) error {
	return x.client.Delete(stateTransitionsList, opts...)
}

func (x *stateTransitionsListClient) Update(stateTransitionsList *StateTransitionsList, opts ...client.UpdateOption) error {
	return x.client.Update(stateTransitionsList, opts...)
}

func (x *CreateStateObjectsList) StateTransition() {}

func (x *CreateStateObjectsList) New() meta.StateTransition {
	return new(CreateStateObjectsList)
}

func (x *CreateStateTransitionsList) StateTransition() {}

func (x *CreateStateTransitionsList) New() meta.StateTransition {
	return new(CreateStateTransitionsList)
}

var StateObjectsListSchema = schema.Definition{
	Meta: meta.Meta{
		APIGroup: "tmos.runtime.v1alpha1",
		APIKind:  "StateObjectsList",
	},
	Singleton: true,
}

var StateTransitionsListSchema = schema.Definition{
	Meta: meta.Meta{
		APIGroup: "tmos.runtime.v1alpha1",
		APIKind:  "StateTransitionsList",
	},
	Singleton: true,
}

type ClientSet interface {
	StateObjectsList() StateObjectsListClient
	StateTransitionsList() StateTransitionsListClient
	ExecCreateStateObjectsList(msg *CreateStateObjectsList) error
	ExecCreateStateTransitionsList(msg *CreateStateTransitionsList) error
}

func NewClientSet(client client.RuntimeClient) ClientSet {
	return &clientSet{
		client:                     client,
		stateObjectsListClient:     &stateObjectsListClient{client: client},
		stateTransitionsListClient: &stateTransitionsListClient{client: client},
	}
}

type clientSet struct {
	client client.RuntimeClient
	// stateObjectsListClient is the client used to interact with StateObjectsList
	stateObjectsListClient StateObjectsListClient
	// stateTransitionsListClient is the client used to interact with StateTransitionsList
	stateTransitionsListClient StateTransitionsListClient
}

func (x *clientSet) StateObjectsList() StateObjectsListClient {
	return x.stateObjectsListClient
}

func (x *clientSet) StateTransitionsList() StateTransitionsListClient {
	return x.stateTransitionsListClient
}

func (x *clientSet) ExecCreateStateObjectsList(msg *CreateStateObjectsList) error {
	return x.client.Deliver(msg)
}

func (x *clientSet) ExecCreateStateTransitionsList(msg *CreateStateTransitionsList) error {
	return x.client.Deliver(msg)
}
