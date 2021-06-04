package schema

import (
	"fmt"

	meta "github.com/fdymylja/tmos/core/meta"
	"github.com/fdymylja/tmos/pkg/protoutils/kindencoder"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// Definition defines an object *Schema
type Definition struct {
	// Singleton marks if there can exist only one instance of this object
	// it's invalid to use primary key alongside a Singleton
	Singleton bool
	// PrimaryKey indicates the field to use as a primary key
	// it must be the json name of the protobuf object
	// NOTE: PrimaryKey must not be set if Singleton is true
	PrimaryKey string
	// SecondaryKeys indicates the protobuf json names of fields
	// of the object to use as secondary keys, the ones that can be
	// passed to the List() endpoints
	// NOTE: SecondaryKeys must not be set if Singleton is true
	SecondaryKeys []string
}

// Validate verifies a Definition
func (d Definition) Validate() error {
	if d.Singleton && d.PrimaryKey != "" {
		return fmt.Errorf("a StateObject can not be singleton and have a primary key at the same time")
	}
	if d.Singleton && len(d.SecondaryKeys) != 0 {
		return fmt.Errorf("a StateObject can not be singleton and have secondary keys at the same time")
	}
	if !d.Singleton && d.PrimaryKey == "" {
		return fmt.Errorf("a StateObject must be a singleton or have a primary key")
	}
	return nil
}

// Schema represents how a meta.StateObject is saved and indexed into the store
// and provides all the required functionalities to index the fields of the object
type Schema struct {
	apiDefinition         *meta.APIDefinition
	mType                 meta.StateObject
	name                  string
	typePrefix            []byte // TODO should we force copies of this?
	primaryKey            protoreflect.FieldDescriptor
	primaryKeyKindEncoder kindencoder.KindEncoder
	secondaryKeys         []*Indexer
	secondaryKeysByField  map[string]*Indexer
	singleton             bool
	hasIndexes            bool
}

func (s *Schema) NewStateObject() meta.StateObject {
	return s.mType.NewStateObject()
}

func (s *Schema) HasIndexes() bool {
	return s.hasIndexes
}

func (s *Schema) TypePrefix() []byte {
	return s.typePrefix
}

func (s *Schema) Name() string {
	return meta.Name(s.mType)
}

// EncodePrimaryKey returns the encoded primary given a meta.StateObject
// NOTE: panics if the field does not belong to the message
func (s *Schema) EncodePrimaryKey(o meta.StateObject) []byte {
	if s.singleton == true {
		return []byte("unique")
	}
	pkValue := o.ProtoReflect().Get(s.primaryKey)
	return s.primaryKeyKindEncoder.EncodeBytes(pkValue)
}

func (s *Schema) Indexer(fieldName string) (*Indexer, error) {
	sk, exists := s.secondaryKeysByField[fieldName]
	if !exists {
		return nil, fmt.Errorf("%w: %s in object %s", ErrSecondaryKey, fieldName, s.name)
	}
	return sk, nil
}

func (s *Schema) Indexes() []*Indexer {
	return s.secondaryKeys
}

func NewSchema(o meta.StateObject, options Definition) (*Schema, error) {
	return parseObjectSchema(o, options)
}

func parseObjectSchema(o meta.StateObject, options Definition) (*Schema, error) {
	if err := options.Validate(); err != nil {
		return nil, fmt.Errorf("%w: %s", ErrBadDefinition, err)
	}
	schema := &Schema{mType: o, apiDefinition: o.APIDefinition()}
	fds := o.ProtoReflect().Descriptor().Fields()
	switch options.Singleton {
	case true:
		schema.singleton = true
		if options.PrimaryKey != "" {
			return nil, fmt.Errorf("%w: can not register a singleton with a primary key in object %s", ErrBadDefinition, meta.Name(o))
		}
	case false:
		primaryKey := fds.ByJSONName(options.PrimaryKey)
		if primaryKey == nil {
			return nil, fmt.Errorf("%w: invalid primary key field %s in object %s", ErrBadDefinition, options.PrimaryKey, meta.Name(o))
		}
		primaryKeyEncoder, err := kindencoder.NewKindEncoder(primaryKey.Kind())
		if err != nil {
			return nil, fmt.Errorf("%w: %s has invalid primary key field: %s", ErrBadDefinition, meta.Name(o), err)
		}
		schema.primaryKey = primaryKey
		schema.primaryKeyKindEncoder = primaryKeyEncoder
	}
	// add prefix and name to schema
	schema.typePrefix = []byte(meta.Name(o))
	schema.name = meta.Name(o)
	// if there are no secondary keys then just skip this part
	if len(options.SecondaryKeys) == 0 {
		schema.hasIndexes = false
		return schema, nil
	}
	// singletons cannot be indexed
	if options.Singleton && len(options.SecondaryKeys) != 0 {
		return nil, fmt.Errorf("%w: singletons can not have secondary indexes in object %s", ErrBadDefinition, meta.Name(o))
	}
	schema.secondaryKeys = make([]*Indexer, len(options.SecondaryKeys))
	schema.secondaryKeysByField = make(map[string]*Indexer, len(options.SecondaryKeys))
	for i, skName := range options.SecondaryKeys {
		sk, err := NewIndexer(o, skName)
		if err != nil {
			return nil, err
		}

		schema.secondaryKeys[i] = sk
		schema.secondaryKeysByField[skName] = sk
	}
	schema.hasIndexes = true
	return schema, nil
}
