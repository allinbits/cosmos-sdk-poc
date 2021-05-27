package schema

import (
	"fmt"

	"github.com/fdymylja/tmos/runtime/meta"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// Definition defines an object *Schema
type Definition struct {
	// Meta contains the information regarding the type API
	Meta meta.Meta
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

// Verify verifies a Definition
func (d Definition) Verify() error {
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

// ValueEncoderFunc is a function that encodes a protoreflect.Value to bytes
type ValueEncoderFunc func(value protoreflect.Value) []byte

// InterfaceEncoderFunc converts an interface to protoreflect.Value,
// returns false if the interface does not match the correct expected type.
type InterfaceEncoderFunc func(i interface{}) (value protoreflect.Value, valid bool)

// Schema represents how a meta.StateObject is saved and indexed into the store
// and provides all the required functionalities to index the fields of the object
type Schema struct {
	messageType          protoreflect.MessageType
	name                 string
	typePrefix           []byte // TODO should we force copies of this?
	primaryKey           protoreflect.FieldDescriptor
	primaryKeyEncode     ValueEncoderFunc
	secondaryKeys        []*Indexer
	secondaryKeysByField map[string]*Indexer
	singleton            bool
	hasIndexes           bool
}

func (s *Schema) HasIndexes() bool {
	return s.hasIndexes
}

func (s *Schema) TypePrefix() []byte {
	return s.typePrefix
}

func (s *Schema) Name() string {
	return s.name
}

// EncodePrimaryKey returns the encoded primary given a meta.StateObject
// NOTE: panics if the field does not belong to the message
func (s *Schema) EncodePrimaryKey(o meta.StateObject) []byte {
	if s.singleton == true {
		return []byte("unique")
	}
	pkValue := o.ProtoReflect().Get(s.primaryKey)
	return s.primaryKeyEncode(pkValue)
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
	schema := &Schema{}
	fds := o.ProtoReflect().Descriptor().Fields()
	switch options.Singleton {
	case true:
		schema.singleton = true
		if options.PrimaryKey != "" {
			return nil, fmt.Errorf("%w: can not register a singleton with a primary key in object %s", ErrBadOptions, meta.Name(o))
		}
	case false:
		primaryKey := fds.ByJSONName(options.PrimaryKey)
		if primaryKey == nil {
			return nil, fmt.Errorf("%w: invalid primary key field %s in object %s", ErrBadOptions, options.PrimaryKey, meta.Name(o))
		}
		primaryKeyEncoder, err := encoderForKind(primaryKey.Kind())
		if err != nil {
			return nil, fmt.Errorf("%w: %s has invalid primary key field: %s", ErrBadOptions, meta.Name(o), err)
		}
		schema.primaryKey = primaryKey
		schema.primaryKeyEncode = primaryKeyEncoder
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
		return nil, fmt.Errorf("%w: singletons can not have secondary indexes in object %s", ErrBadOptions, meta.Name(o))
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
