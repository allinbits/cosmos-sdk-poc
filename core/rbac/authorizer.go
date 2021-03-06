package rbac

import (
	"fmt"

	meta "github.com/fdymylja/tmos/core/meta"
	"github.com/fdymylja/tmos/core/rbac/v1alpha1"
	runtimev1alpha1 "github.com/fdymylja/tmos/core/runtime/v1alpha1"
	"github.com/fdymylja/tmos/runtime/authorization"
	"github.com/fdymylja/tmos/runtime/module"
	"github.com/scylladb/go-set/strset"
)

var _ authorization.Authorizer = Authorizer{}

func NewAuthorizer(c module.Client) authorization.Authorizer {
	return Authorizer{c: c}
}

type Authorizer struct {
	c module.Client
}

func (a Authorizer) Authorize(attributes authorization.Attributes) (authorization.Decision, error) {
	usersStr := make([]string, len(attributes.Users.List()))
	for i, u := range attributes.Users.List() {
		usersStr[i] = u.GetName()
	}

	roles, err := a.fetchRoles(usersStr...)
	if err != nil {
		return authorization.DecisionDeny, err
	}
	// check if role is allowed to access the resource
	set := a.getResourcesSet(attributes.Verb, roles)
	if !set.Has(meta.Name(attributes.Resource)) {
		return authorization.DecisionDeny, fmt.Errorf("no subject in %s has role %s towards resource %s", usersStr, attributes.Verb, meta.Name(attributes.Resource))
	}
	return authorization.DecisionAllow, nil
}

func (a Authorizer) fetchRoles(subjects ...string) ([]*v1alpha1.Role, error) {
	roles := make([]*v1alpha1.Role, len(subjects))
	for i, subject := range subjects {
		// get role binding for subject
		roleBinding := new(v1alpha1.RoleBinding)
		err := a.c.Get(meta.NewStringID(subject), roleBinding)
		if err != nil {
			return nil, err
		}
		// get role
		role := new(v1alpha1.Role)
		err = a.c.Get(meta.NewStringID(roleBinding.RoleRef), role)
		if err != nil {
			return nil, err
		}
		// append role
		roles[i] = role
	}
	return roles, nil
}

// getResourcesSet gets the available resources given a Verb and a slice of roles.
func (a Authorizer) getResourcesSet(verb runtimev1alpha1.Verb, roles []*v1alpha1.Role) *strset.Set {
	var resources []string
	for _, role := range roles {
		res := role.GetResourcesForVerb(verb)
		resources = append(resources, res...)
	}
	return strset.New(resources...)
}
