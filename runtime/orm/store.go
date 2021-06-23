package orm

import (
	"fmt"

	meta "github.com/fdymylja/tmos/core/meta"
	"github.com/fdymylja/tmos/runtime/kv"
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
	List(schema *schema.Schema, options ListOptions) (iter kv.Iterator, err error)
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

func (s Store) RegisterObject(object meta.StateObject, options *schema.Definition) error {
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

type ListOptions struct {
	MatchFieldInterface []ListMatchFieldInterface
	MatchFieldString    []ListMatchFieldString
	Start, End          uint64
}

func (l ListOptions) String() string {
	return fmt.Sprintf("%#v", l)
}

// ListMatchFieldInterface can be used to match a field
// with an interface{} value which is converted to the field's
// concrete value.
type ListMatchFieldInterface struct {
	Field string
	Value interface{}
}

// ListMatchFieldString can be used to match a field
// with a string value which is converted to the field's
// concrete value.
type ListMatchFieldString struct {
	Field string
	Value string
}

func (s Store) List(object meta.StateObject, listOptions ListOptions) (Iterator, error) {
	sch, err := s.schemas.Get(object)
	if err != nil {
		return Iterator{}, err
	}
	iter, err := s.indexes.List(sch, listOptions)
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

// SchemaRegistry returns the schema.Registry of the store
// TODO(fdymylja): I'm not sure if the store should be returning this
// or if the runtime should own this object.
func (s Store) SchemaRegistry() *schema.Registry {
	return s.schemas
}

func (s Store) LatestVersion() uint64 {
	return 0
}

// LoadVersion returns an instance of the Store at the given height
// returns an error if the version cannot be loaded
// TODO load correct version
func (s Store) LoadVersion(height uint64) (Store, error) {
	return s, nil
}
