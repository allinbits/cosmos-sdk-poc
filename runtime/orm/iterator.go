package orm

import (
	meta "github.com/fdymylja/tmos/core/meta"
	"github.com/fdymylja/tmos/runtime/kv"
	"github.com/fdymylja/tmos/runtime/orm/schema"
)

type Iterator struct {
	store  Store
	iter   kv.Iterator
	schema *schema.Schema
}

func (i Iterator) Get(o meta.StateObject) error {
	key := i.iter.Key()
	return i.store.Get(meta.NewBytesID(key), o)
}

func (i Iterator) Next() {
	i.iter.Next()
}

func (i Iterator) Valid() bool {
	return i.iter.Valid()
}

func (i Iterator) Close() {
	i.iter.Close()
}

// Delete deletes the current iterator object.
// TODO(fdymylja): implement fast deletion path, as this one unmarshals the object which is not needed since we know the key.
func (i Iterator) Delete() error {
	obj := i.schema.NewStateObject()
	err := i.Get(obj)
	if err != nil {
		return err
	}
	return i.store.Delete(obj)
}
