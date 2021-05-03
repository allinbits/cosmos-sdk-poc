package authentication

func NewSubjects() *Subjects {
	return &Subjects{
		authenticatedSubjects: map[string]struct{}{},
	}
}

// Subjects is a convenience struct for checking
// if a transaction was authenticated by the given
// entities
type Subjects struct {
	authenticatedSubjects map[string]struct{}
}

func (s *Subjects) HasAuthenticated(subject string) bool {
	_, ok := s.authenticatedSubjects[subject]
	return ok
}

func (s *Subjects) Add(subject string) {
	s.authenticatedSubjects[subject] = struct{}{}
}

func (s *Subjects) List() []string {
	l := make([]string, 0, len(s.authenticatedSubjects))
	for k := range s.authenticatedSubjects {
		l = append(l, k)
	}
	return l
}
