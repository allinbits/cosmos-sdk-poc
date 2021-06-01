package objects

import (
	"fmt"

	meta "github.com/fdymylja/tmos/core/meta"
	"github.com/fdymylja/tmos/runtime/kv"
	"github.com/fdymylja/tmos/runtime/orm"
	"github.com/fdymylja/tmos/runtime/orm/schema"
	"google.golang.org/protobuf/proto"
)

var marshalOptions = proto.MarshalOptions{
	Deterministic: true,
}

var unmarshalOptions = proto.UnmarshalOptions{}

func marshal(o proto.Message) []byte {
	if len(o.ProtoReflect().GetUnknown()) != 0 {
		panic(fmt.Errorf("object with unknown during marshal: %s", o))
	}
	b, err := marshalOptions.Marshal(o)
	if err != nil {
		panic(err)
	}
	return b
}

func unmarshal(b []byte, target proto.Message) {
	err := unmarshalOptions.Unmarshal(b, target)
	if err != nil {
		panic(err)
	}
	if len(target.ProtoReflect().GetUnknown()) != 0 {
		panic(fmt.Errorf("object unmarshalled with unknowns: %s", target))
	}
}

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
		return fmt.Errorf("%w: %s", orm.ErrAlreadyExists, primaryKey)
	}
	// set
	objBytes := marshal(object)
	s.kv.Set(primaryKey, objBytes)
	return nil
}

func (s *Store) Get(sch *schema.Schema, id meta.ID, object meta.StateObject) error {
	key, err := saveKeyRaw(sch, id.Bytes())
	if err != nil {
		return err
	}
	objectBytes, exists := s.kv.Get(key)
	if !exists {
		return fmt.Errorf("%w: %s", orm.ErrNotFound, key)
	}

	unmarshal(objectBytes, object)
	return nil
}

func (s *Store) Update(sch *schema.Schema, object meta.StateObject) error {
	key, err := saveKey(sch, object)
	if err != nil {
		return err
	}
	if !s.kv.Has(key) {
		return fmt.Errorf("%w %s", orm.ErrNotFound, key)
	}

	// set
	objBytes := marshal(object)
	s.kv.Set(key, objBytes)
	return nil
}

func (s *Store) Delete(sch *schema.Schema, object meta.StateObject) error {
	pk, err := saveKey(sch, object)
	if err != nil {
		return err
	}
	if !s.kv.Has(pk) {
		return fmt.Errorf("%w: %s", orm.ErrNotFound, pk)
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
		return nil, fmt.Errorf("orm: empty primary key for object %s", s.Name())
	}
	pk := make([]byte, 0, len(s.TypePrefix())+1+len(key))
	pk = append(pk, s.TypePrefix()...)
	pk = append(pk, '/')
	pk = append(pk, key...)
	return pk, nil
}

// indexingKey marks the key used to index the object given its type
type indexingKey struct {
	prefix []byte
	key    []byte
}

func (k indexingKey) marshal() []byte {
	panic("impl")
}
