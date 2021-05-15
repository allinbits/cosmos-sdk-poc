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

type KVStore interface {
	Get(key []byte) (value []byte, exists bool)
	Set(key, value []byte)
	Has(key []byte) (exists bool)
	Delete(key []byte) (exists bool)
	Iterate(start, end []byte) KVStoreIterator
}

type KVStoreIterator interface {
	Next()
	Key() []byte
	Value() []byte
	Valid() bool
	Close()
}

type ObjectSchema struct {
	name   string
	prefix []byte

	primaryKey        protoreflect.FieldDescriptor
	primaryKeyEncoder FieldEncoderFunc

	secondaryKeys        map[string]protoreflect.FieldDescriptor
	secondaryKeysEncoder map[string]FieldEncoderFunc
}

func NewStore(kv KVStore) *Store {
	return &Store{
		knownObjects: map[string]struct{}{},
		schemas:      map[string]*ObjectSchema{},
		kv:           kv,
	}
}

type Store struct {
	knownObjects map[string]struct{}
	schemas      map[string]*ObjectSchema
	kv           KVStore
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
	if s.knownObject(o) {
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
	s.schemas[objectName(o)] = schema
	return nil
}

func getObjectSchema(o meta.StateObject, options RegisterObjectOptions) (*ObjectSchema, error) {
	schema := &ObjectSchema{}
	fds := o.ProtoReflect().Descriptor().Fields()
	primaryKey := fds.ByJSONName(options.PrimaryKey)
	primaryKeyEncoder, err := EncoderForKind(primaryKey.Kind())
	if err != nil {
		return nil, fmt.Errorf("store: %s has invalid primary key field: %w", objectName(o), err)
	}
	schema.primaryKey = primaryKey
	schema.primaryKeyEncoder = primaryKeyEncoder

	schema.secondaryKeys = make(map[string]protoreflect.FieldDescriptor, len(options.SecondaryKeys))
	schema.secondaryKeysEncoder = make(map[string]FieldEncoderFunc, len(options.SecondaryKeys))
	for _, sk := range options.SecondaryKeys {
		secondaryKey := fds.ByJSONName(sk)
		secondaryKeyEncoder, err := EncoderForKind(secondaryKey.Kind())
		if err != nil {
			return nil, fmt.Errorf("store: %s has invalid secondary key field: %w", objectName(o), err)
		}
		schema.secondaryKeys[sk] = secondaryKey
		schema.secondaryKeysEncoder[sk] = secondaryKeyEncoder
	}

	return schema, nil
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
			return fmt.Errorf("secondary key %s does not exist in object %s, make sure to use the correct json name of the field", sk, objectName(o))
		}
	}
	return nil
}
