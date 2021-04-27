package iavl

import runtime "github.com/fdymylja/tmos/runtime/store"

func NewStore() runtime.Store {
	return store{}
}

type store struct {
}

func (s store) Get(id runtime.ID) error {
	panic("implement me")
}

func (s store) Create(id runtime.ID, object runtime.StateObject) error {
	panic("implement me")
}

func (s store) Update(id runtime.ID, object runtime.StateObject) error {
	panic("implement me")
}

func (s store) List(object runtime.StateObject) error {
	panic("implement me")
}
