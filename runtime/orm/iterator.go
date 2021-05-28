package orm

import (
	"github.com/fdymylja/tmos/runtime/kv"
	"github.com/fdymylja/tmos/runtime/meta"
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

// TODO implement fast delete function currently to delete you need to get object and delete it
// when we can skip the unmarshalling part and go straight to deletion of the key.
