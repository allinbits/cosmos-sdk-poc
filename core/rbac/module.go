package rbac

import (
	"encoding/json"

	"github.com/fdymylja/tmos/core/rbac/v1alpha1"
	"github.com/fdymylja/tmos/runtime/authorization"
	"github.com/fdymylja/tmos/runtime/module"
)

func NewModule() *Module { return &Module{} }

type Module struct {
	authorizer authorization.Authorizer
	genesis    *genesis
}

func (m *Module) AsAuthorizer() authorization.Authorizer {
	return m.authorizer
}

func (m *Module) AddInitialRole(role *v1alpha1.Role, binding *v1alpha1.RoleBinding) {
	m.genesis.addInitialRole(role, binding)
}

func (m *Module) Initialize(client module.Client) module.Descriptor {
	m.authorizer = NewAuthorizer(client)
	m.genesis = newGenesis()

	return module.NewDescriptorBuilder().
		Named("rbac").
		OwnsStateObject(&v1alpha1.Role{}, v1alpha1.RoleSchema).
		OwnsStateObject(&v1alpha1.RoleBinding{}, v1alpha1.RoleBindingSchema).
		OwnsStateObject(&v1alpha1.Params{}, v1alpha1.ParamsSchema).
		HandlesStateTransition(&v1alpha1.MsgCreateRole{}, NewCreateRoleHandler(), false).
		HandlesAdmission(&v1alpha1.MsgCreateRole{}, NewCreateRoleAdmissionHandler(client)).
		HandlesStateTransition(&v1alpha1.MsgBindRole{}, NewBindRoleHandler(), false).
		HandlesAdmission(&v1alpha1.MsgBindRole{}, NewBindRoleAdmission(client)).
		WithGenesis(m.genesis).Build()
}

func newGenesis() *genesis {
	return &genesis{}
}

type genesis struct {
	roles    []*v1alpha1.Role
	bindings []*v1alpha1.RoleBinding
}

func (g *genesis) addInitialRole(role *v1alpha1.Role, binding *v1alpha1.RoleBinding) {
	g.roles = append(g.roles, role)
	g.bindings = append(g.bindings, binding)
}

func (g *genesis) Default(client module.Client) error {
	// we create the initial roles and bindings of the associated roles
	// they are core roles created at runtime.DescriptorBuilder.Build() level
	// we do it via deliver because we want to make sure the creation
	// goes through proper checks.
	for i, r := range g.roles {
		err := client.Deliver(&v1alpha1.MsgCreateRole{NewRole: r})
		if err != nil {
			return err
		}
		binding := g.bindings[i]
		if binding == nil {
			continue
		}
		err = client.Deliver(&v1alpha1.MsgBindRole{
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
