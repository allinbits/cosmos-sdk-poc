package orm

import (
	"github.com/fdymylja/tmos/runtime/kv"
	"github.com/fdymylja/tmos/runtime/meta"
	"github.com/fdymylja/tmos/runtime/orm/schema"
)

type ObjectsStore interface {
	// Create creates the object given its schema.Schema
	Create(schema *schema.Schema, o meta.StateObject) error
	// Get fetches the object given its meta.ID and schema.Schema
	// and puts it into target
	Get(schema *schema.Schema, id meta.ID, target meta.StateObject) error
	// Update updates the object given its schema.Schema
	// and the object to update
	Update(schema *schema.Schema, o meta.StateObject) error
	// Delete deletes the object given its schema.Schema
	Delete(schema *schema.Schema, o meta.StateObject) error
}

type Indexer interface {
	// Index indexes the object
	Index(schema *schema.Schema, o meta.StateObject) error
	// ClearIndexes clears the indexes for the object
	ClearIndexes(schema *schema.Schema, o meta.StateObject) error
	// List provides an iterator which returns the primary keys
	// of the object (not-prefixed)
	List(schema *schema.Schema, options ListOptionsRaw) (iter kv.Iterator, err error)
}

func NewStore(objects ObjectsStore, indexer Indexer) Store {
	return Store{
		objects: objects,
		indexes: indexer,
		schemas: schema.NewRegistry(),
	}
}

type Store struct {
	objects ObjectsStore
	indexes Indexer
	schemas *schema.Registry
}

func (s Store) RegisterObject(object meta.StateObject, options schema.Definition) error {
	err := s.schemas.AddObject(object, options)
	if err != nil {
		return err
	}
	return nil
}

func (s Store) Create(object meta.StateObject) error {
	sch, err := s.schemas.Get(object)
	if err != nil {
		return err
	}
	// create object
	err = s.objects.Create(sch, object)
	if err != nil {
		return err
	}
	// index object
	err = s.indexes.Index(sch, object)
	if err != nil {
		return err
	}
	return nil
}

func (s Store) Get(id meta.ID, object meta.StateObject) error {
	sch, err := s.schemas.Get(object)
	if err != nil {
		return err
	}
	return s.objects.Get(sch, id, object)
}

func (s Store) Update(object meta.StateObject) error {
	sch, err := s.schemas.Get(object)
	if err != nil {
		return err
	}
	// update object
	err = s.objects.Update(sch, object)
	if err != nil {
		return err
	}
	// clear indexes
	err = s.indexes.ClearIndexes(sch, object)
	if err != nil {
		return err
	}
	// update indexes
	err = s.indexes.Index(sch, object)
	if err != nil {
		return err
	}

	return nil
}

func (s Store) Delete(object meta.StateObject) error {
	sch, err := s.schemas.Get(object)
	if err != nil {
		return err
	}
	// delete the object
	err = s.objects.Delete(sch, object)
	if err != nil {
		return err
	}
	// clear its indexes
	err = s.indexes.ClearIndexes(sch, object)
	if err != nil {
		return err
	}
	return nil
}

type ListOptionsRaw struct {
	MatchFields []ListMatchField
}

type ListMatchField struct {
	Field string
	Value interface{}
}

func (l ListMatchField) apply(opt *ListOptionsRaw) {
	opt.MatchFields = append(opt.MatchFields, l)
}

type ListOptionApplier interface{ apply(opt *ListOptionsRaw) }

func (s Store) List(object meta.StateObject, options ...ListOptionApplier) (Iterator, error) {
	opt := new(ListOptionsRaw)
	for _, o := range options {
		o.apply(opt)
	}
	sch, err := s.schemas.Get(object)
	if err != nil {
		return Iterator{}, err
	}
	iter, err := s.indexes.List(sch, *opt)
	if err != nil {
		return Iterator{}, err
	}
	return Iterator{
		store: s,
		iter:  iter,
	}, nil
}

func (s Store) ListRegisteredStateObjects() []string {
	return s.schemas.List()
}

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
