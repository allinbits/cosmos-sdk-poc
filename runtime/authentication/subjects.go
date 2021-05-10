package authentication

import "strings"

type ImmutableSubjects interface {
	List() []string
	Contains(subject ...string) bool
}

type MutableSubjects interface {
	ImmutableSubjects
	Add(subject string)
}

func NewEmptySubjects() *Subjects {
	return &Subjects{
		authenticatedSubjects: map[string]struct{}{},
	}
}

func NewSubjects(ss ...string) *Subjects {
	s := &Subjects{authenticatedSubjects: make(map[string]struct{}, len(ss))}
	for _, sub := range ss {
		s.authenticatedSubjects[sub] = struct{}{}
	}
	return s
}

// Subjects is a convenience struct for checking
// if a transaction was authenticated by the given
// entities
type Subjects struct {
	authenticatedSubjects map[string]struct{}
}

func (x *Subjects) Contains(ss ...string) bool {
	for _, s := range ss {
		if !x.contains(s) {
			return false
		}
	}
	return true
}

func (x *Subjects) contains(subject string) bool {
	_, ok := x.authenticatedSubjects[subject]
	return ok
}

func (x *Subjects) Add(subject string) {
	x.authenticatedSubjects[subject] = struct{}{}
}

func (x *Subjects) List() []string {
	l := make([]string, 0, len(x.authenticatedSubjects))
	for k := range x.authenticatedSubjects {
		l = append(l, k)
	}
	return l
}

func (x *Subjects) String() string {
	return strings.Join(x.List(), ",")
}
