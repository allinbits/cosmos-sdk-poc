package rbac

import (
	"encoding/json"

	"github.com/fdymylja/tmos/core/rbac/v1alpha1"
	runtimev1alpha1 "github.com/fdymylja/tmos/core/runtime/v1alpha1"
	"github.com/fdymylja/tmos/runtime/module"
)

func NewModule() *Module { return &Module{} }

type Module struct{}

func (m *Module) Initialize(client module.Client) module.Descriptor {

	return module.NewDescriptorBuilder().
		Named("rbac").
		OwnsStateObject(&v1alpha1.Role{}, v1alpha1.RoleSchema).
		OwnsStateObject(&v1alpha1.RoleBinding{}, v1alpha1.RoleBindingSchema).
		OwnsStateObject(&v1alpha1.Params{}, v1alpha1.ParamsSchema).
		HandlesStateTransition(&v1alpha1.MsgCreateRole{}, NewCreateRoleHandler(client)).
		HandlesAdmission(&v1alpha1.MsgCreateRole{}, NewCreateRoleAdmissionHandler(client)).
		HandlesStateTransition(&v1alpha1.MsgBindRole{}, NewBindRoleHandler(client)).
		HandlesAdmission(&v1alpha1.MsgBindRole{}, NewBindRoleAdmission(client)).
		WithPostStateTransitionHandler(&runtimev1alpha1.CreateModuleDescriptors{}, NewInitRoleCreator(client)).
		WithGenesis(newGenesis(client)).
		IsAuthorizer(NewAuthorizer(client)).
		Build()
}

func newGenesis(client module.Client) *genesis {
	return &genesis{c: client}
}

type genesis struct {
	c module.Client

	roles    []*v1alpha1.Role
	bindings []*v1alpha1.RoleBinding
}

func (g *genesis) addInitialRole(role *v1alpha1.Role, binding *v1alpha1.RoleBinding) {
	g.roles = append(g.roles, role)
	g.bindings = append(g.bindings, binding)
}

func (g *genesis) Default() error {
	// we create the initial roles and bindings of the associated roles
	// they are core roles created at runtime.DescriptorBuilder.Build() level
	// we do it via deliver because we want to make sure the creation
	// goes through proper checks.
	for i, r := range g.roles {
		err := g.c.Deliver(&v1alpha1.MsgCreateRole{NewRole: r})
		if err != nil {
			return err
		}
		binding := g.bindings[i]
		if binding == nil {
			continue
		}
		err = g.c.Deliver(&v1alpha1.MsgBindRole{
			RoleId:  binding.RoleRef,
			Subject: binding.Subject,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func (g *genesis) Import(state json.RawMessage) error {
	panic("implement me")
}

func (g *genesis) Export() (json.RawMessage, error) {
	panic("implement me")
}
