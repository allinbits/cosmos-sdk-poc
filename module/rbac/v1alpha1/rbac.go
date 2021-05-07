package v1alpha1

import (
	runtimev1alpha1 "github.com/fdymylja/tmos/module/runtime/v1alpha1"
	"github.com/fdymylja/tmos/runtime/meta"
	"github.com/scylladb/go-set/strset"
)

func (x *Role) GetID() meta.ID { return meta.NewStringID(x.Id) }

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

func (x *Role) appendToVerb(verb runtimev1alpha1.Verb, resources ...meta.Type) {

}

func (x *Role) Extend(verb runtimev1alpha1.Verb, resources ...meta.Type) error {
	res := x.GetResourcesForVerb(verb)
	// in case its empty simply append
	if len(res) == 0 {
		x.appendToVerb(verb, resources...)
	}
	set := strset.New(res...)
	for _, r := range resources {
		if !set.Has(meta.)
	}
}

func (x *RoleBinding) GetID() meta.ID { return meta.NewStringID(x.Subject) }

func (x *MsgBindRole) StateTransition()   {}
func (x *MsgCreateRole) StateTransition() {}
