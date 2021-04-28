package iavl

import (
	runtime2 "github.com/fdymylja/tmos/runtime"
	"github.com/fdymylja/tmos/runtime/meta"
)

func NewStore() runtime2.Store {
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
