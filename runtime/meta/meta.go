package meta

import "fmt"

// APIGroup marks a group of APIs which belong to one module
type APIGroup string

func (g APIGroup) String() string {
	return (string)(g)
}

// APIKind marks the single type, which belongs to a group.
type APIKind string

func (g APIKind) String() string {
	return (string)(g)
}

// Meta defines an object belonging to the Runtime
type Meta struct {
	// APIGroup is the APIGroup of a Type
	APIGroup APIGroup
	// APIKind is the unique Type name of an object belonging to an APIGroup
	APIKind APIKind
}

// Validate asserts if the Meta object is correctly formed
func (m Meta) Validate() error {
	if m.APIGroup == "" || m.APIKind == "" {
		return fmt.Errorf("meta: bad Meta definition")
	}
	return nil
}

func (m Meta) Fullname() string {
	return fmt.Sprintf("%s.%s", m.APIGroup, m.APIKind)
}
