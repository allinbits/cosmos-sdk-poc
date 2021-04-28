package iavl

import (
	runtime2 "github.com/fdymylja/tmos/runtime"
	runtime "github.com/fdymylja/tmos/runtime/store"
)

func NewStore() runtime.Store {
	return store{}
}

type store struct {
}

func (s store) Get(id runtime2.ID) error {
	panic("implement me")
}

func (s store) Create(id runtime2.ID, object runtime2.StateObject) error {
	panic("implement me")
}

func (s store) Update(id runtime2.ID, object runtime2.StateObject) error {
	panic("implement me")
}

func (s store) List(object runtime2.StateObject) error {
	panic("implement me")
}
