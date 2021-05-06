package rbac

import (
	"fmt"

	"github.com/fdymylja/tmos/module/rbac/v1alpha1"
	runtimev1alpha1 "github.com/fdymylja/tmos/module/runtime/v1alpha1"
	"github.com/fdymylja/tmos/runtime/meta"
	"github.com/fdymylja/tmos/runtime/module"
	"github.com/scylladb/go-set/strset"
)

type Authorizer struct {
	c module.Client
}

func (a Authorizer) Allowed(verb runtimev1alpha1.Verb, resource meta.Type, subjects ...string) error {
	roles, err := a.fetchRoles(subjects...)
	if err != nil {
		return err
	}
	// check if role is allowed to access the resource
	set := a.getResourcesSet(verb, roles)
	if !set.Has(meta.Name(resource)) {
		return fmt.Errorf("no subject in %#v has role %s towards resource %s", subjects, verb, meta.Name(resource))
	}
	return nil
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
