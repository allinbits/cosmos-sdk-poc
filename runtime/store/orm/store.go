package orm

import (
	"errors"
	"fmt"

	"github.com/fdymylja/tmos/runtime/meta"
	"google.golang.org/protobuf/reflect/protoreflect"
)

var (
	ErrAlreadyRegistered = errors.New("store: object already registered")
)

type FieldEncoder func(value protoreflect.Value) []byte

type KVStore interface {
	Get(key []byte) (value []byte, exists bool)
	Set(key, value []byte)
	Has(key []byte) (exists bool)
	Delete(key []byte) (exists bool)
}

type ObjectSchema struct {
	name              string
	prefix            []byte
	primaryKey        protoreflect.FieldDescriptor
	primaryKeyEncoder func(value protoreflect.Value) []byte

	secondaryKeys        map[string]protoreflect.FieldDescriptor
	secondaryKeysEncoder map[string]func(value protoreflect.Value) []byte
}

type Store struct {
	knownObjects map[string]struct{}
	schemas      map[string]ObjectSchema
}

type RegisterObjectOptions struct {
	// PrimaryKey indicates the field to use as a primary key
	// it must be the json name of the protobuf object
	PrimaryKey string
	// SecondaryKeys indicates the protobuf json names of fields
	// of the object to use as secondary keys, the ones that can be
	// passed to Store.List
	SecondaryKeys []string
}

func (s *Store) RegisterObject(o meta.StateObject, options RegisterObjectOptions) error {
	if !s.knownObject(o) {
		return ErrAlreadyRegistered
	}
	err := assertValidFields(o, options)
	if err != nil {
		return err
	}
	schema, err := getObjectSchema(o, options)
	if err != nil {
		return err
	}

}

func getObjectSchema(o meta.StateObject, options RegisterObjectOptions) (ObjectSchema, error) {
	fds := o.ProtoReflect().Descriptor().Fields()
	primaryKey := fds.ByJSONName(options.PrimaryKey)
	encoder := getEncoderForKind(primaryKey.Kind())
	return ObjectSchema{
		name:                 "",
		prefix:               nil,
		primaryKey:           primaryKey,
		primaryKeyEncoder:    nil,
		secondaryKeys:        nil,
		secondaryKeysEncoder: nil,
	}, nil
}

func (s *Store) knownObject(o meta.StateObject) bool {
	name := objectName(o)
	if _, exists := s.knownObjects[name]; exists {
		return true
	}
	return false
}

func objectName(o meta.StateObject) string {
	return meta.Name(o)
}

func assertValidFields(o meta.StateObject, options RegisterObjectOptions) error {
	fds := o.ProtoReflect().Descriptor().Fields()

	// check if primary key exists
	if fd := fds.ByJSONName(options.PrimaryKey); fd == nil {
		return fmt.Errorf("primary key %s does not exist in object %s, make sure to use the correct json name of the field", options.PrimaryKey, objectName(o))
	}
	// check if every secondary keys exists
	for _, sk := range options.SecondaryKeys {
		if fd := fds.ByJSONName(sk); fd == nil {
			return fmt.Errorf("secondary key %s does not exist in object %s, make sure to use the correct json name of the field", options.PrimaryKey, objectName(o))
		}
	}
	return nil
}
