package iavl

import (
	"github.com/fdymylja/tmos/runtime/meta"
	store2 "github.com/fdymylja/tmos/runtime/store"
)

func NewStore() store2.Store {
	return store{}
}

type store struct {
}

func (s store) Get(id meta.ID) error {
	panic("implement me")
}

func (s store) Create(id meta.ID, object meta.StateObject) error {
	panic("implement me")
}

func (s store) Update(id meta.ID, object meta.StateObject) error {
	panic("implement me")
}

func (s store) List(object meta.StateObject) error {
	panic("implement me")
}
