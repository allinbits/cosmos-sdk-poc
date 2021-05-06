package rbac

import (
	"github.com/fdymylja/tmos/module/rbac/v1alpha1"
	"github.com/fdymylja/tmos/runtime/module"
)

func NewModule() Module { return Module{} }

type Module struct {
}

func (m Module) Initialize(client module.Client, builder *module.Builder) {
	builder.
		Named("rbac").
		OwnsStateObject(&v1alpha1.Role{}).
		OwnsStateObject(&v1alpha1.RoleBinding{}).
		HandlesStateTransition(&v1alpha1.MsgCreateRole{}, NewCreateRoleController(client)).
		HandlesAdmission(&v1alpha1.MsgCreateRole{}, NewCreateRoleAdmissionController(client)).
		HandlesStateTransition(&v1alpha1.MsgBindRole{}, NewBindRoleController(client)).
		HandlesAdmission(&v1alpha1.MsgBindRole{}, NewBindRoleAdmission(client))
}
