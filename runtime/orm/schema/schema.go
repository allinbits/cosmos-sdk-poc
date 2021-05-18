package schema

import (
	"fmt"

	"github.com/fdymylja/tmos/runtime/meta"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type FieldEncoderFunc func(value protoreflect.Value) []byte

type Schema struct {
	Type                 protoreflect.MessageType
	Name                 string
	TypePrefix           []byte // TODO should we force copies of this?
	PrimaryKey           protoreflect.FieldDescriptor
	PrimaryKeyEncode     FieldEncoderFunc
	SecondaryKeys        map[string]protoreflect.FieldDescriptor
	SecondaryKeyEncoders map[string]FieldEncoderFunc
}

// EncodePrimaryKey returns the encoded primary given a meta.StateObject
// NOTE: panics if the field does not belong to the message
func (s *Schema) EncodePrimaryKey(o meta.StateObject) []byte {
	pkValue := o.ProtoReflect().Get(s.PrimaryKey)
	return s.PrimaryKeyEncode(pkValue)
}

// EncodeObjectField returns the encoded secondary key given a meta.StateObject
// and the json name of the field to encode.
// NOTE: panics if the key provided is not part of the schema.
func (s *Schema) EncodeObjectField(key string, object meta.StateObject) ([]byte, error) {
	fd, exists := s.SecondaryKeys[key]
	if !exists {
		panic(fmt.Errorf("%w: object %s is not indexed by secondary key %s", ErrSecondaryKey, meta.Name(object), key))
	}
	encode := s.SecondaryKeyEncoders[key]
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

type Options struct {
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
	primaryKey := fds.ByJSONName(options.PrimaryKey)
	if primaryKey == nil {
		return nil, fmt.Errorf("%w: invalid primary key field %s in object %s", ErrRegister, options.PrimaryKey, meta.Name(o))
	}
	primaryKeyEncoder, err := encoderForKind(primaryKey.Kind())
	if err != nil {
		return nil, fmt.Errorf("store: %s has invalid primary key field: %w", meta.Name(o), err)
	}
	schema.PrimaryKey = primaryKey
	schema.PrimaryKeyEncode = primaryKeyEncoder

	schema.SecondaryKeys = make(map[string]protoreflect.FieldDescriptor, len(options.SecondaryKeys))
	schema.SecondaryKeyEncoders = make(map[string]FieldEncoderFunc, len(options.SecondaryKeys))
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
		schema.SecondaryKeyEncoders[sk] = secondaryKeyEncoder
	}

	schema.TypePrefix = []byte(meta.Name(o))
	schema.Name = meta.Name(o)
	return schema, nil
}
