package schema

import (
	"fmt"

	"github.com/fdymylja/tmos/runtime/meta"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type ValueEncoderFunc func(value protoreflect.Value) []byte
type InterfaceEncoderFunc func(i interface{}) (value protoreflect.Value, valid bool)

type Schema struct {
	messageType          protoreflect.MessageType
	name                 string
	typePrefix           []byte // TODO should we force copies of this?
	primaryKey           protoreflect.FieldDescriptor
	primaryKeyEncode     ValueEncoderFunc
	secondaryKeys        []*SecondaryKey
	secondaryKeysByField map[string]*SecondaryKey
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

func (s *Schema) Indexer(fieldName string) (*SecondaryKey, error) {
	sk, exists := s.secondaryKeysByField[fieldName]
	if !exists {
		return nil, fmt.Errorf("%w: %s in object %s", ErrSecondaryKey, fieldName, s.name)
	}
	return sk, nil
}

func (s *Schema) Indexes() []*SecondaryKey {
	return s.secondaryKeys
}

func (s *Schema) Singleton() bool {
	return s.singleton
}

type Options struct {
	// Singleton marks if there can exist only one instance of this object
	// it's invalid to use primary key alongside a Singleton
	Singleton bool
	// PrimaryKey indicates the field to use as a primary key
	// it must be the json name of the protobuf object
	PrimaryKey string
	// SecondaryKeys indicates the protobuf json names of fields
	// of the object to use as secondary keys, the ones that can be
	// passed to Store.List
	SecondaryKeys []string
}

func NewSchema(o meta.StateObject, options Options) (*Schema, error) {
	return parseObjectSchema(o, options)
}

func parseObjectSchema(o meta.StateObject, options Options) (*Schema, error) {
	schema := &Schema{}
	fds := o.ProtoReflect().Descriptor().Fields()
	switch options.Singleton {
	case true:
		schema.singleton = true
		if options.PrimaryKey != "" {
			return nil, fmt.Errorf("schema: singleton do not have primary keys")
		}
	case false:
		primaryKey := fds.ByJSONName(options.PrimaryKey)
		if primaryKey == nil {
			return nil, fmt.Errorf("%w: invalid primary key field %s in object %s", ErrRegister, options.PrimaryKey, meta.Name(o))
		}
		primaryKeyEncoder, err := encoderForKind(primaryKey.Kind())
		if err != nil {
			return nil, fmt.Errorf("store: %s has invalid primary key field: %w", meta.Name(o), err)
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

	schema.secondaryKeys = make([]*SecondaryKey, len(options.SecondaryKeys))
	schema.secondaryKeysByField = make(map[string]*SecondaryKey, len(options.SecondaryKeys))
	for i, skName := range options.SecondaryKeys {
		sk, err := NewSecondaryKey(o, skName)
		if err != nil {
			return nil, err
		}

		schema.secondaryKeys[i] = sk
		schema.secondaryKeysByField[skName] = sk
	}
	schema.hasIndexes = true
	return schema, nil
}
