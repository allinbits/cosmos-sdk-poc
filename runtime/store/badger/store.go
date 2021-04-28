package badger

import (
	"errors"
	"fmt"

	"github.com/fdymylja/tmos/runtime/meta"
	"k8s.io/klog/v2"

	"google.golang.org/protobuf/proto"
)

var (
	ErrNotFound      = errors.New("store: object not found")
	ErrAlreadyExists = errors.New("store: object already registered")
)

type kvStore interface {
	has(k []byte) bool
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

// Store defines the state object store
type Store struct {
	objectPrefixes map[string][]byte // objectPrefixes maps object to their prefixes
	kv             kvStore
}

func (s *Store) Get(id meta.ID, object meta.StateObject) error {
	key := s.keyForWithID(id, object)
	o, exists := s.kv.get(key)
	if !exists {
		return fmt.Errorf("%w: %s", ErrNotFound, key)
	}
	v := o
	err := proto.Unmarshal(v, object)
	if err != nil {
		panic(err)
	}
	return nil
}

func (s *Store) Create(object meta.StateObject) error {
	key := s.keyFor(object)
	if s.kv.has(key) {
		return ErrAlreadyExists
	}
	klog.Infof("creating: %s %s", key, object)
	b, err := proto.Marshal(object)
	if err != nil {
		panic(err)
	}
	s.kv.set(key, b)
	return nil
}

func (s *Store) Update(object meta.StateObject) error {
	k := s.keyFor(object)
	if !s.kv.has(k) {
		return fmt.Errorf("%w: %s", ErrNotFound, k)
	}
	b, err := proto.Marshal(object)
	if err != nil {
		panic(err)
	}
	s.kv.set(k, b)
	return nil
}

func (s *Store) Delete(object meta.StateObject) error {
	key := s.keyFor(object)
	s.kv.delete(key)

	return nil
}

func (s *Store) keyForWithID(id meta.ID, object meta.StateObject) []byte {
	// get prefix for object
	name := meta.Name(object)
	pfx, ok := s.objectPrefixes[name]
	if !ok {
		panic("unregistered state object: " + name)
	}
	// TODO(fdymylja): should we panic or return nil ?
	key := id.Bytes()
	if len(key) == 0 {
		panic("object key returned an empty length byte array")
	}
	return append(pfx, key...)
}

func (s *Store) keyFor(object meta.StateObject) []byte {
	return s.keyForWithID(object.GetID(), object)
}

func (s *Store) RegisterStateObject(object meta.StateObject) error {
	name := meta.Name(object)
	// check if registered
	_, exists := s.objectPrefixes[name]
	if exists {
		return fmt.Errorf("%w: %s", ErrAlreadyExists, name)
	}
	// register object
	s.objectPrefixes[name] = []byte(name + "/")
	return nil
}
