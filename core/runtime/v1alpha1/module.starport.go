package v1alpha1

import (
	meta "github.com/fdymylja/tmos/core/meta"
	client "github.com/fdymylja/tmos/runtime/client"
	schema "github.com/fdymylja/tmos/runtime/orm/schema"
)

func (x *ModuleDescriptors) APIDefinition() *meta.APIDefinition {
	return &meta.APIDefinition{
		Group:   "tmos.runtime.v1alpha1",
		Kind:    "ModuleDescriptors",
		ApiType: meta.APIType_StateObject,
	}
}

func (x *ModuleDescriptors) NewStateObject() meta.StateObject {
	return new(ModuleDescriptors)
}

type ModuleDescriptorsClient interface {
	Get(opts ...client.GetOption) (*ModuleDescriptors, error)
	Create(moduleDescriptors *ModuleDescriptors, opts ...client.CreateOption) error
	Delete(moduleDescriptors *ModuleDescriptors, opts ...client.DeleteOption) error
	Update(moduleDescriptors *ModuleDescriptors, opts ...client.UpdateOption) error
}

type moduleDescriptorsClient struct {
	client client.Client
}

func (x *moduleDescriptorsClient) Get(opts ...client.GetOption) (*ModuleDescriptors, error) {
	_spfGenO := new(ModuleDescriptors)
	_spfGenErr := x.client.Get(meta.SingletonID, _spfGenO, opts...)
	if _spfGenErr != nil {
		return nil, _spfGenErr
	}
	return _spfGenO, nil
}

func (x *moduleDescriptorsClient) Create(moduleDescriptors *ModuleDescriptors, opts ...client.CreateOption) error {
	return x.client.Create(moduleDescriptors, opts...)
}

func (x *moduleDescriptorsClient) Delete(moduleDescriptors *ModuleDescriptors, opts ...client.DeleteOption) error {
	return x.client.Delete(moduleDescriptors, opts...)
}

func (x *moduleDescriptorsClient) Update(moduleDescriptors *ModuleDescriptors, opts ...client.UpdateOption) error {
	return x.client.Update(moduleDescriptors, opts...)
}

func (x *CreateModuleDescriptors) APIDefinition() *meta.APIDefinition {
	return &meta.APIDefinition{
		Group:   "tmos.runtime.v1alpha1",
		Kind:    "CreateModuleDescriptors",
		ApiType: meta.APIType_StateTransition,
	}
}

func (x *CreateModuleDescriptors) NewStateTransition() meta.StateTransition {
	return new(CreateModuleDescriptors)
}

var ModuleDescriptorsSchema = &schema.Definition{
	Singleton: true,
}

type ClientSet interface {
	ModuleDescriptors() ModuleDescriptorsClient
	ExecCreateModuleDescriptors(msg *CreateModuleDescriptors) error
}

func NewClientSet(client client.Client) ClientSet {
	return &clientSet{
		client:                  client,
		moduleDescriptorsClient: &moduleDescriptorsClient{client: client},
	}
}

type clientSet struct {
	client client.Client
	// moduleDescriptorsClient is the client used to interact with ModuleDescriptors
	moduleDescriptorsClient ModuleDescriptorsClient
}

func (x *clientSet) ModuleDescriptors() ModuleDescriptorsClient {
	return x.moduleDescriptorsClient
}

func (x *clientSet) ExecCreateModuleDescriptors(msg *CreateModuleDescriptors) error {
	return x.client.Deliver(msg)
}
