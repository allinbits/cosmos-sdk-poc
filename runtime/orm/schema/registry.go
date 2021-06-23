package schema

import (
	"fmt"

	"github.com/fdymylja/tmos/core/meta"
)

func NewRegistry() *Registry {
	return &Registry{
		schemas:      map[string]*Schema{},
		apiGroups:    map[string][]string{},
		schemaByMeta: map[string]map[string]*Schema{},
	}
}

type Registry struct {
	schemas      map[string]*Schema
	apiGroups    map[string][]string
	schemaByMeta map[string]map[string]*Schema
}

func (s *Registry) Add(sch *Schema) error {
	_, exists := s.schemas[sch.Name()]
	if exists {
		return fmt.Errorf("%w: %s", ErrAlreadyExists, sch.Name())
	}
	s.schemas[sch.Name()] = sch
	s.apiGroups[sch.apiDefinition.Group] = append(s.apiGroups[sch.apiDefinition.Group], sch.apiDefinition.Kind)

	// map by group and kind
	_, exists = s.schemaByMeta[sch.apiDefinition.Group]
	// if group was not set add it
	if !exists {
		s.schemaByMeta[sch.apiDefinition.Group] = map[string]*Schema{}
	}
	s.schemaByMeta[sch.apiDefinition.Group][sch.apiDefinition.Kind] = sch
	return nil
}

func (s *Registry) AddObject(o meta.StateObject, options *Definition) error {
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

func (s *Registry) ListAPIGroups() []string {
	groups := make([]string, 0, len(s.apiGroups))
	for k := range s.apiGroups {
		groups = append(groups, k)
	}
	return groups
}

func (s *Registry) ListKindsInGroup(group string) ([]string, error) {
	kinds, exists := s.apiGroups[group]
	if !exists {
		return nil, fmt.Errorf("%w: API Group not found %s", ErrNotFound, group)
	}
	return kinds, nil
}

func (s *Registry) GetByAPIDefinition(ad *meta.APIDefinition) (*Schema, error) {
	kinds, exist := s.schemaByMeta[ad.Group]
	if !exist {
		return nil, fmt.Errorf("%w: API group does not exist %s", ErrNotFound, ad.Group)
	}
	sch, exist := kinds[ad.Kind]
	if !exist {
		return nil, fmt.Errorf("%w: kind %s not found in API group %s", ErrNotFound, ad.Kind, ad.Group)
	}
	return sch, nil
}
