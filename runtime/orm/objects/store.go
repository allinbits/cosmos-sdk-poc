package objects

import (
	"fmt"

	kv "github.com/fdymylja/tmos/runtime/kv"
	"github.com/fdymylja/tmos/runtime/meta"
	"github.com/fdymylja/tmos/runtime/orm"
	"github.com/fdymylja/tmos/runtime/orm/schema"
	"google.golang.org/protobuf/proto"
)

var marshal = proto.MarshalOptions{
	Deterministic: true,
}
var unmarshal = proto.UnmarshalOptions{}

func NewStore(kv kv.KV) *Store {
	return &Store{
		kv: kv,
	}
}

type Store struct {
	kv kv.KV
}

func (s *Store) Create(sch *schema.Schema, object meta.StateObject) error {
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

func (s *Store) Get(sch *schema.Schema, id meta.ID, object meta.StateObject) error {
	if sch.Singleton() {
		id = meta.NewStringID("unique")
	}
	key, err := saveKeyRaw(sch, id.Bytes())
	if err != nil {
		return err
	}
	objectBytes, exists := s.kv.Get(key)
	if !exists {
		return fmt.Errorf("%w: %s", orm.ErrNotFound, key)
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

func (s *Store) Update(sch *schema.Schema, object meta.StateObject) error {
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

func (s *Store) Delete(sch *schema.Schema, object meta.StateObject) error {
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

func saveKey(s *schema.Schema, object meta.StateObject) ([]byte, error) {
	// encode key
	pk := s.EncodePrimaryKey(object)
	return saveKeyRaw(s, pk)
}

func saveKeyRaw(s *schema.Schema, key []byte) ([]byte, error) {
	if len(key) == 0 {
		return nil, fmt.Errorf("orm: empty primary key for object %s", s.Name)
	}
	pk := make([]byte, 0, len(s.TypePrefix)+1+len(key))
	pk = append(pk, s.TypePrefix...)
	pk = append(pk, '/')
	pk = append(pk, key...)
	return pk, nil
}
