package v1alpha1

import (
	"fmt"

	"github.com/fdymylja/tmos/core/meta"
	runtimev1alpha1 "github.com/fdymylja/tmos/core/runtime/v1alpha1"
	"github.com/scylladb/go-set/strset"
)

// ExternalAccountRoleID defines the external account role
// which every newly created account has, it gives them
// access to Exec on external handlers.
const ExternalAccountRoleID = "external_account"

func NewEmptyRole(name string) *Role {
	return &Role{Id: roleNameForModule(name)}
}

func (x *Role) GetResourcesForVerb(verb runtimev1alpha1.Verb) []string {
	switch verb {
	case runtimev1alpha1.Verb_Get:
		return x.Gets
	case runtimev1alpha1.Verb_List:
		return x.Lists
	case runtimev1alpha1.Verb_Create:
		return x.Creates
	case runtimev1alpha1.Verb_Delete:
		return x.Deletes
	case runtimev1alpha1.Verb_Update:
		return x.Updates
	case runtimev1alpha1.Verb_Deliver:
		return x.Delivers
	default:
		return nil
	}
}

func (x *Role) ExtendRaw(verb runtimev1alpha1.Verb, apiDefinition *meta.APIDefinition) error {
	apiDef := apiDefinition.Name()

	if x.hasResourceForVerb(verb, apiDef) {
		return fmt.Errorf("role %s has already resource %s", x, apiDef)
	}

	err := x.appendResourceToVerb(verb, apiDef)
	if err != nil {
		return err
	}

	return nil
}

func (x *Role) Extend(verb runtimev1alpha1.Verb, resource meta.APIObject) error {
	return x.ExtendRaw(verb, resource.APIDefinition())
}

func (x *Role) hasResourceForVerb(verb runtimev1alpha1.Verb, resourceRaw string) bool {
	res := x.GetResourcesForVerb(verb)
	set := strset.New(res...)
	if len(res) != 0 && set.Has(resourceRaw) {
		return true
	}

	return false
}

func (x *Role) appendResourceToVerb(verb runtimev1alpha1.Verb, resourceRaw string) error {
	switch verb {
	case runtimev1alpha1.Verb_Get:
		x.Gets = append(x.Gets, resourceRaw)
		return nil
	case runtimev1alpha1.Verb_List:
		x.Lists = append(x.Lists, resourceRaw)
		return nil
	case runtimev1alpha1.Verb_Create:
		x.Creates = append(x.Creates, resourceRaw)
		return nil
	case runtimev1alpha1.Verb_Delete:
		x.Deletes = append(x.Deletes, resourceRaw)
		return nil
	case runtimev1alpha1.Verb_Update:
		x.Updates = append(x.Updates, resourceRaw)
		return nil
	case runtimev1alpha1.Verb_Deliver:
		x.Delivers = append(x.Delivers, resourceRaw)
		return nil
	default:
		return fmt.Errorf("unknown verb %s", verb)
	}
}

func roleNameForModule(name string) string {
	const moduleRoleSuffix = "role"
	return fmt.Sprintf("%s-%s", name, moduleRoleSuffix)
}
