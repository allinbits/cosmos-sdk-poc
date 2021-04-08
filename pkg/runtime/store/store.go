package store

import (
	"google.golang.org/protobuf/proto"
	"k8s.io/klog/v2"

	"github.com/fdymylja/tmos/pkg/application"
)

var _ application.StateObjectRegisterer = (*Store)(nil)
var _ application.StateObjectClient = (*Store)(nil)

type kvStore interface {
	get(k []byte) (v []byte, exists bool)
	set(k []byte, v []byte)
	delete(k []byte)
}

func NewStore() *Store {
	db := NewBadger()
	return &Store{
		objectPrefixes: make(map[string][]byte),
		kv:             db,
	}
}

type Store struct {
	objectPrefixes map[string][]byte

	kv kvStore
}

func (s *Store) Get(object application.StateObject) (exists bool) {
	key := s.keyFor(object)
	o, exists := s.kv.get(key)
	if !exists {
		return false
	}
	v := o
	err := proto.Unmarshal(v, object)
	if err != nil {
		panic(err)
	}
	return true
}

func (s *Store) Set(object application.StateObject) {
	key := s.keyFor(object)
	b, err := proto.Marshal(object)
	if err != nil {
		panic(err)
	}
	s.kv.set(key, b)
}

func (s *Store) Delete(object application.StateObject) {
	key := s.keyFor(object)
	s.kv.delete(key)
}

func (s *Store) keyFor(object application.StateObject) []byte {
	// get prefix for object
	name := (string)(object.ProtoReflect().Descriptor().FullName())
	pfx, ok := s.objectPrefixes[name]
	if !ok {
		panic("unregistered state object: " + name)
	}
	meta := object.GetMeta()
	// TODO(fdymylja): should we panic or return nil ?
	if meta == nil {
		panic("state object with nil meta")
	}
	if meta.Id == "" {
		panic("state object with empty meta.ID")
	}
	return append(pfx, []byte(meta.Id)...)
}

func (s *Store) RegisterStateObject(object application.StateObject) {
	name := (string)(object.ProtoReflect().Descriptor().FullName())
	_, exists := s.objectPrefixes[name]
	if exists {
		panic("state object already registered twice: " + name)
	}
	s.objectPrefixes[name] = []byte(name)
	klog.InfoS("registering object", "context", "store", "object", name, "prefix", []byte(name))
}
