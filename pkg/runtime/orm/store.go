package orm

import (
	"errors"
	"fmt"

	"github.com/fdymylja/tmos/apis/meta"
	"google.golang.org/protobuf/proto"
)

var (
	ErrUnauthorized  = errors.New("store: user does not own object")
	ErrAlreadyExists = errors.New("store: object already registered")
)

type kvStore interface {
	get(k []byte) (v []byte, exists bool)
	set(k []byte, v []byte)
	delete(k []byte)
}

func NewStore() *Store {
	db := NewBadger()
	return &Store{
		objectPrefixes: make(map[string][]byte),
		objectOwners:   make(map[string]map[string]struct{}),
		kv:             db,
	}
}

// Store defines the state object store
type Store struct {
	objectPrefixes map[string][]byte              // objectPrefixes maps object to their prefixes
	objectOwners   map[string]map[string]struct{} // objectOwners defines which identities own the given object
	kv             kvStore
}

func (s *Store) Get(object meta.StateObject) (exists bool) {
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

func (s *Store) Set(user string, object meta.StateObject) error {
	if !s.owns(user, object) {
		return fmt.Errorf("%w: %s does not own the object", ErrUnauthorized, user)
	}
	key := s.keyFor(object)
	b, err := proto.Marshal(object)
	if err != nil {
		panic(err)
	}
	s.kv.set(key, b)
	return nil
}

func (s *Store) Delete(user string, object meta.StateObject) error {
	if !s.owns(user, object) {
		return fmt.Errorf("%w: %s does not own the object", ErrUnauthorized, user)
	}
	key := s.keyFor(object)
	s.kv.delete(key)

	return nil
}

func (s *Store) keyFor(object meta.StateObject) []byte {
	// get prefix for object
	name := meta.Name(object)
	pfx, ok := s.objectPrefixes[name]
	if !ok {
		panic("unregistered state object: " + name)
	}
	metadata := object.GetObjectMeta()
	// TODO(fdymylja): should we panic or return nil ?
	if metadata == nil {
		panic("state object with nil meta")
	}
	if metadata.Id == "" {
		panic("state object with empty meta.ID")
	}
	return append(pfx, []byte(metadata.Id)...)
}

func (s *Store) RegisterStateObject(owner string, object meta.StateObject) error {
	name := meta.Name(object)
	// check if registered
	_, exists := s.objectPrefixes[name]
	if exists {
		return fmt.Errorf("%w: %s", ErrAlreadyExists, name)
	}
	// register object
	s.objectPrefixes[name] = []byte(name)
	// register owner
	if _, exists := s.objectOwners[owner]; !exists {
		s.objectOwners[owner] = make(map[string]struct{})
	}
	s.objectOwners[owner][name] = struct{}{}
	return nil
}

// owns asserts if the given user owns the provided object
func (s *Store) owns(user string, object meta.StateObject) bool {
	ownedObjects, exists := s.objectOwners[user]
	if !exists {
		return false
	}
	_, owned := ownedObjects[meta.Name(object)]
	return owned
}
