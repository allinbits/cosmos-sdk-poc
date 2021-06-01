package v1alpha1

import (
	"fmt"

	meta "github.com/fdymylja/tmos/core/meta"
	runtimev1alpha1 "github.com/fdymylja/tmos/core/runtime/v1alpha1"
	"github.com/scylladb/go-set/strset"
)

// ExternalAccountRoleID defines the external account role
// which every newly created account has, it gives them
// access to Exec on external handlers.
const ExternalAccountRoleID = "external_account"

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

func (x *Role) appendToVerb(verb runtimev1alpha1.Verb, resource meta.Type) error {
	name := meta.Name(resource)
	switch verb {
	case runtimev1alpha1.Verb_Get:
		x.Gets = append(x.Gets, name)
		return nil
	case runtimev1alpha1.Verb_List:
		x.Lists = append(x.Lists, name)
		return nil
	case runtimev1alpha1.Verb_Create:
		x.Creates = append(x.Creates, name)
		return nil
	case runtimev1alpha1.Verb_Delete:
		x.Deletes = append(x.Deletes, name)
		return nil
	case runtimev1alpha1.Verb_Update:
		x.Updates = append(x.Updates, name)
		return nil
	case runtimev1alpha1.Verb_Deliver:
		x.Delivers = append(x.Delivers, name)
		return nil
	default:
		return fmt.Errorf("unknown verb %s", verb)
	}
}

func (x *Role) Extend(verb runtimev1alpha1.Verb, resource meta.Type) error {
	res := x.GetResourcesForVerb(verb)
	set := strset.New(res...)
	if len(res) != 0 && set.Has(meta.Name(resource)) {
		return fmt.Errorf("role %s has already resource %s", x, meta.Name(resource))
	}
	err := x.appendToVerb(verb, resource)
	if err != nil {
		return err
	}
	return nil
}
