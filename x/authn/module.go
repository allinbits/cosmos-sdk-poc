package authn

import (
	rbacv1alpha1 "github.com/fdymylja/tmos/core/rbac/v1alpha1"
	"github.com/fdymylja/tmos/runtime/authentication"
	"github.com/fdymylja/tmos/runtime/module"
	extensions2 "github.com/fdymylja/tmos/x/authn/extensions"
	v1alpha12 "github.com/fdymylja/tmos/x/authn/v1alpha1"
)

// Module implements the authentication.Module
type Module struct {
	txDecoder authentication.TxDecoder
}

func NewModule() *Module {
	return &Module{}
}

func (m *Module) Initialize(c module.Client) module.Descriptor {
	m.txDecoder = newAuthenticator(c)

	return module.NewDescriptorBuilder().
		Named("authn").
		HandlesStateTransition(&v1alpha12.MsgCreateAccount{}, NewCreateAccountController(c), true).
		HandlesAdmission(&v1alpha12.MsgCreateAccount{}, NewCreateAccountAdmissionController()).
		OwnsStateObject(&v1alpha12.Account{}).
		OwnsStateObject(&v1alpha12.Params{}).
		OwnsStateObject(&v1alpha12.CurrentAccountNumber{}).
		ExtendsAuthentication(extensions2.New(c)).
		WithGenesis(genesis{c: c}).
		NeedsStateTransition(&rbacv1alpha1.MsgBindRole{}).Build()
}

func (m *Module) GetTxDecoder() authentication.TxDecoder {
	return m.txDecoder
}
