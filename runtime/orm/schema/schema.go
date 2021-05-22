package schema

import (
	"fmt"

	"github.com/fdymylja/tmos/runtime/meta"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type FieldEncoderFunc func(value protoreflect.Value) []byte

type Schema struct {
	oType                 protoreflect.MessageType
	Name                  string
	TypePrefix            []byte // TODO should we force copies of this?
	primaryKey            protoreflect.FieldDescriptor
	primaryKeyEncode      FieldEncoderFunc
	SecondaryKeys         map[string]protoreflect.FieldDescriptor
	secondaryKeysEncoders map[string]FieldEncoderFunc

	singleton bool
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

// EncodeObjectField returns the encoded secondary key given a meta.StateObject
// and the json name of the field to encode.
// NOTE: panics if the key provided is not part of the schema.
func (s *Schema) EncodeObjectField(key string, object meta.StateObject) ([]byte, error) {
	fd, exists := s.SecondaryKeys[key]
	if !exists {
		panic(fmt.Errorf("%w: object %s is not indexed by secondary key %s", ErrSecondaryKey, meta.Name(object), key))
	}
	encode := s.secondaryKeysEncoders[key]
	return encode(object.ProtoReflect().Get(fd)), nil
}

func (s *Schema) MustEncodeObjectField(key string, object meta.StateObject) []byte {
	k, err := s.EncodeObjectField(key, object)
	if err != nil {
		panic(err)
	}
	return k
}

func (s *Schema) EncodeFieldInterface(fieldName string, i interface{}) ([]byte, error) {
	fd, exists := s.SecondaryKeys[fieldName]
	if !exists {
		return nil, fmt.Errorf("%w: field not found %s in object %s", ErrSecondaryKey, fieldName, s.Name)
	}
	encodedBytes, err := safeFieldEncodeInterface(fd, i)
	if err != nil {
		return nil, fmt.Errorf("%w: %s %s", ErrFieldEncode, fieldName, err)
	}
	return encodedBytes, nil
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
	return getObjectSchema(o, options)
}

func getObjectSchema(o meta.StateObject, options Options) (*Schema, error) {
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

	schema.SecondaryKeys = make(map[string]protoreflect.FieldDescriptor, len(options.SecondaryKeys))
	schema.secondaryKeysEncoders = make(map[string]FieldEncoderFunc, len(options.SecondaryKeys))
	for _, sk := range options.SecondaryKeys {
		secondaryKey := fds.ByJSONName(sk)
		if secondaryKey == nil {
			return nil, fmt.Errorf("%w: invalid secondary key field %s in object %s", ErrRegister, sk, meta.Name(o))
		}
		secondaryKeyEncoder, err := encoderForKind(secondaryKey.Kind())
		if err != nil {
			return nil, fmt.Errorf("store: %s has invalid secondary key field: %w", meta.Name(o), err)
		}
		schema.SecondaryKeys[sk] = secondaryKey
		schema.secondaryKeysEncoders[sk] = secondaryKeyEncoder
	}

	schema.TypePrefix = []byte(meta.Name(o))
	schema.Name = meta.Name(o)
	return schema, nil
}
