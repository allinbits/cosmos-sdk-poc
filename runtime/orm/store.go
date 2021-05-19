package orm

import (
	"github.com/fdymylja/tmos/runtime/meta"
	"github.com/fdymylja/tmos/runtime/orm/schema"
)

type ObjectsStore interface {
	Create(schema *schema.Schema, o meta.StateObject) error
	Get(schema *schema.Schema, id meta.ID, o meta.StateObject) error
	Update(schema *schema.Schema, o meta.StateObject) error
	Delete(schema *schema.Schema, o meta.StateObject) error
}

type Indexer interface {
	// Index indexes the object
	Index(schema *schema.Schema, o meta.StateObject) error
	// ClearIndexes clears the indexes for the object
	ClearIndexes(schema *schema.Schema, o meta.StateObject) error
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

func (s Store) RegisterObject(object meta.StateObject, options schema.Options) error {
	err := s.schemas.AddObject(object, options)
	if err != nil {
		return err
	}
	return nil
}
