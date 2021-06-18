package client

import (
	"fmt"

	"github.com/fdymylja/tmos/core/meta"
	"github.com/fdymylja/tmos/runtime/orm"
)

var _ RuntimeClient = &ORMClient{}

// NewORMClient instantiates a new ReadOnlyClient from an orm.Store
func NewORMClient(store orm.Store, identifier string) RuntimeClient {
	return &ORMClient{
		identifier: identifier,
		store:      store,
	}
}

// ORMClient is a ReadOnlyClient that can be created from an orm.Store
// it is meant to be used internally, but by extension services that
// have no write interaction with the store
type ORMClient struct {
	identifier string
	store      orm.Store // TODO(fdymylja): this should be a read only store.
}

func (o *ORMClient) Create(_ meta.StateObject, _ ...CreateOption) error {
	return fmt.Errorf("%s is not allowed to Create", o.identifier)
}

func (o *ORMClient) Update(_ meta.StateObject, _ ...UpdateOption) error {
	return fmt.Errorf("%s is not allowed to Update", o.identifier)
}

func (o *ORMClient) Delete(_ meta.StateObject, _ ...DeleteOption) error {
	return fmt.Errorf("%s is not allowed to Delete", o.identifier)
}

func (o *ORMClient) Deliver(_ meta.StateTransition, _ ...DeliverOption) error {
	return fmt.Errorf("%s is not allowed to Deliver", o.identifier)
}

// Get implements RuntimeClient
func (o *ORMClient) Get(id meta.ID, object meta.StateObject, opts ...GetOption) error {
	opt := new(getOptions)
	for _, o := range opts {
		o(opt)
	}

	height := o.store.LatestVersion()
	if opt.Height != 0 {
		height = opt.Height
	}

	store, err := o.store.LoadVersion(height)
	if err != nil {
		return err
	}

	return store.Get(id, object)
}

// List implements RuntimeClient
func (o *ORMClient) List(object meta.StateObject, opts ...ListOption) (ObjectIterator, error) {
	opt := new(listOptions)
	for _, o := range opts {
		o(opt)
	}

	height := o.store.LatestVersion()
	if opt.Height != 0 {
		height = opt.Height
	}

	store, err := o.store.LoadVersion(height)
	if err != nil {
		return ObjectIterator{}, err
	}

	return store.List(object, opt.ORMOptions)
}
