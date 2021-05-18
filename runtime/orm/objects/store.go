package objects

import (
	"fmt"

	kv "github.com/fdymylja/tmos/runtime/kv"
	"github.com/fdymylja/tmos/runtime/meta"
	"github.com/fdymylja/tmos/runtime/orm/schema"
	"google.golang.org/protobuf/proto"
)

var marshal = proto.MarshalOptions{
	Deterministic: true,
}
var unmarshal = proto.UnmarshalOptions{}

func NewStore(kv kv.KV) *Store {
	return &Store{
		kv:      kv,
		schemas: map[string]*schema.Schema{},
	}
}

type Store struct {
	kv      kv.KV
	schemas map[string]*schema.Schema
}

func (s *Store) AddSchema(sch *schema.Schema) error {
	if _, exists := s.schemas[sch.Name]; exists {
		return fmt.Errorf("orm: object %s is already part of the schema", sch.Name)
	}
	s.schemas[sch.Name] = sch
	return nil
}

func (s *Store) unknown(object meta.StateObject) bool {
	_, known := s.schemas[meta.Name(object)]
	if known {
		return false
	}
	return true
}

func (s *Store) Create(object meta.StateObject) error {
	sch, err := s.getSchema(object)
	if err != nil {
		return err
	}
	primaryKey, err := saveKey(sch, object)
	if err != nil {
		return err
	}
	// check if object exists
	if s.kv.Has(primaryKey) {
		return fmt.Errorf("orm: object already exists %s", primaryKey)
	}
	// set
	objBytes, err := marshal.Marshal(object)
	if err != nil {
		return fmt.Errorf("orm: codec error %w", err)
	}
	s.kv.Set(primaryKey, objBytes)
	return nil
}

func (s *Store) Get(id meta.ID, object meta.StateObject) error {
	sch, err := s.getSchema(object)
	if err != nil {
		return err
	}
	key, err := saveKeyRaw(sch, id.Bytes())
	if err != nil {
		return err
	}
	objectBytes, exists := s.kv.Get(key)
	if !exists {
		return fmt.Errorf("orm: object not found %s", key)
	}
	err = unmarshal.Unmarshal(objectBytes, object)
	if err != nil {
		return fmt.Errorf("orm: codec error: %w", err)
	}
	if len(object.ProtoReflect().GetUnknown()) != 0 {
		panic(fmt.Sprintf("orm: object with unknowns during unmarshal %s", object))
	}
	return nil
}

func (s *Store) Update(object meta.StateObject) error {
	sch, err := s.getSchema(object)
	if err != nil {
		return err
	}
	key, err := saveKey(sch, object)
	if err != nil {
		return err
	}
	if !s.kv.Has(key) {
		return fmt.Errorf("orm: object not found %s", key)
	}

	// set
	objBytes, err := marshal.Marshal(object)
	if err != nil {
		return fmt.Errorf("orm: codec error %w", err)
	}
	s.kv.Set(key, objBytes)
	return nil
}

func (s *Store) Delete(object meta.StateObject) error {
	sch, err := s.getSchema(object)
	if err != nil {
		return err
	}
	pk, err := saveKey(sch, object)
	if err != nil {
		return err
	}
	if !s.kv.Has(pk) {
		return fmt.Errorf("orm: object not found %s", pk)
	}
	s.kv.Delete(pk)
	return nil
}

func (s *Store) getSchema(object meta.StateObject) (sch *schema.Schema, err error) {
	sch, exists := s.schemas[meta.Name(object)]
	if !exists {
		return nil, fmt.Errorf("store: unknown object %s", meta.Name(object))
	}
	return sch, nil
}

func saveKey(s *schema.Schema, object meta.StateObject) ([]byte, error) {
	// encode key
	pk := s.EncodePrimaryKey(object)
	return saveKeyRaw(s, pk)
}

func saveKeyRaw(s *schema.Schema, key []byte) ([]byte, error) {
	if len(key) == 0 {
		return nil, fmt.Errorf("orm: empty primary key for object %s", s.Name)
	}
	pk := make([]byte, len(s.Prefix)+len(key))
	pk = append(pk, s.Prefix...)
	pk = append(pk, key...)
	return pk, nil
}
