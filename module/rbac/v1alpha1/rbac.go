package v1alpha1

import (
	runtimev1alpha1 "github.com/fdymylja/tmos/module/runtime/v1alpha1"
	"github.com/fdymylja/tmos/runtime/meta"
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

func (x *RoleBinding) GetID() meta.ID { return meta.NewStringID(x.Subject) }

func (x *MsgBindRole) StateTransition()   {}
func (x *MsgCreateRole) StateTransition() {}
