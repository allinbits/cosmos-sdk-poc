package schema

import (
	"fmt"

	"github.com/fdymylja/tmos/runtime/meta"
)

func NewRegistry() *Registry {
	return &Registry{schemas: map[string]*Schema{}}
}

type Registry struct {
	schemas map[string]*Schema
}

func (s *Registry) Add(sch *Schema) error {
	_, exists := s.schemas[sch.Name()]
	if exists {
		return fmt.Errorf("%w: %s", ErrAlreadyExists, sch.Name())
	}
	s.schemas[sch.Name()] = sch
	return nil
}

func (s *Registry) AddObject(o meta.StateObject, options Definition) error {
	sch, err := NewSchema(o, options)
	if err != nil {
		return err
	}
	return s.Add(sch)
}

func (s *Registry) Get(o meta.StateObject) (*Schema, error) {
	sch, exists := s.schemas[meta.Name(o)]
	if !exists {
		return nil, fmt.Errorf("%w: %s", ErrNotFound, meta.Name(o))
	}
	return sch, nil
}

func (s *Registry) List() []string {
	list := make([]string, 0, len(s.schemas))
	for s := range s.schemas {
		list = append(list, s)
	}
	return list
}
