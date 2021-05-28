package schema

import (
	"fmt"

	"github.com/fdymylja/tmos/runtime/meta"
)

func NewRegistry() *Registry {
	return &Registry{
		schemas:      map[string]*Schema{},
		apiGroups:    map[meta.APIGroup][]meta.APIKind{},
		schemaByMeta: map[meta.APIGroup]map[meta.APIKind]*Schema{},
	}
}

type Registry struct {
	schemas      map[string]*Schema
	apiGroups    map[meta.APIGroup][]meta.APIKind
	schemaByMeta map[meta.APIGroup]map[meta.APIKind]*Schema
}

func (s *Registry) Add(sch *Schema) error {
	_, exists := s.schemas[sch.Name()]
	if exists {
		return fmt.Errorf("%w: %s", ErrAlreadyExists, sch.Name())
	}
	s.schemas[sch.Name()] = sch
	s.apiGroups[sch.meta.APIGroup] = append(s.apiGroups[sch.meta.APIGroup], sch.meta.APIKind)

	// map by group and kind
	_, exists = s.schemaByMeta[sch.meta.APIGroup]
	// if group was not set add it
	if !exists {
		s.schemaByMeta[sch.meta.APIGroup] = map[meta.APIKind]*Schema{}
	}
	s.schemaByMeta[sch.meta.APIGroup][sch.meta.APIKind] = sch
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

func (s *Registry) ListAPIGroups() []meta.APIGroup {
	groups := make([]meta.APIGroup, 0, len(s.apiGroups))
	for k := range s.apiGroups {
		groups = append(groups, k)
	}
	return groups
}

func (s *Registry) ListKindsInGroup(group meta.APIGroup) ([]meta.APIKind, error) {
	kinds, exists := s.apiGroups[group]
	if !exists {
		return nil, fmt.Errorf("%w: API Group not found %s", ErrNotFound, group)
	}
	return kinds, nil
}

func (s *Registry) GetByMeta(m meta.Meta) (*Schema, error) {
	kinds, exist := s.schemaByMeta[m.APIGroup]
	if !exist {
		return nil, fmt.Errorf("%w: API group does not exist %s", ErrNotFound, m.APIGroup)
	}
	sch, exist := kinds[m.APIKind]
	if !exist {
		return nil, fmt.Errorf("%w: kind %s not found in API group %s", ErrNotFound, m.APIKind, m.APIGroup)
	}
	return sch, nil
}
