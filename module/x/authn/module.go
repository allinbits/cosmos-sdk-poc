package authn

import (
	"github.com/fdymylja/tmos/module/x/authn/extensions"
	"github.com/fdymylja/tmos/module/x/authn/v1alpha1"
	"github.com/fdymylja/tmos/runtime/authentication"
	"github.com/fdymylja/tmos/runtime/module"
)

// Module implements the authentication.Module
type Module struct {
	authenticator authentication.Authenticator
}

func NewModule() *Module {
	return &Module{}
}

func (m *Module) Initialize(c module.Client, builder *module.Builder) {
	m.authenticator = newAuthenticator(c)
	builder.
		Named("authn").
		HandlesStateTransition(&v1alpha1.MsgCreateAccount{}, NewCreateAccountController(c)).
		HandlesAdmission(&v1alpha1.MsgCreateAccount{}, NewCreateAccountAdmissionController()).
		OwnsStateObject(&v1alpha1.Account{}).
		OwnsStateObject(&v1alpha1.Params{}).
		OwnsStateObject(&v1alpha1.CurrentAccountNumber{}).
		ExtendsAuthentication(extensions.New(c))
}

func (m *Module) GetAuthenticator() authentication.Authenticator {
	return m.authenticator
}
