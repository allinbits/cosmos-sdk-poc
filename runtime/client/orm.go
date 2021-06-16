package client

import (
	"github.com/fdymylja/tmos/core/meta"
	"github.com/fdymylja/tmos/runtime/orm"
)

var _ ReadOnlyClient = &ORMClient{}

// NewORMClient instantiates a new ReadOnlyClient from an orm.Store
func NewORMClient(store orm.Store) ReadOnlyClient {
	return &ORMClient{store: store}
}

// ORMClient is a ReadOnlyClient that can be created from an orm.Store
// it is meant to be used internally, but by extension services that
// have no write interaction with the store
type ORMClient struct {
	store orm.Store // TODO(fdymylja): this should be a read only store.
}

// Get implements ReadOnlyClient
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

// List implements ReadOnlyClient
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
