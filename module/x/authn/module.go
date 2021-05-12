package authn

import (
	rbacv1alpha1 "github.com/fdymylja/tmos/module/rbac/v1alpha1"
	"github.com/fdymylja/tmos/module/x/authn/extensions"
	"github.com/fdymylja/tmos/module/x/authn/v1alpha1"
	"github.com/fdymylja/tmos/runtime/authentication"
	"github.com/fdymylja/tmos/runtime/module"
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
		HandlesStateTransition(&v1alpha1.MsgCreateAccount{}, NewCreateAccountController(c), true).
		HandlesAdmission(&v1alpha1.MsgCreateAccount{}, NewCreateAccountAdmissionController()).
		OwnsStateObject(&v1alpha1.Account{}).
		OwnsStateObject(&v1alpha1.Params{}).
		OwnsStateObject(&v1alpha1.CurrentAccountNumber{}).
		ExtendsAuthentication(extensions.New(c)).
		WithGenesis(genesis{c: c}).
		NeedsStateTransition(&rbacv1alpha1.MsgBindRole{}).Build()
}

func (m *Module) GetTxDecoder() authentication.TxDecoder {
	return m.txDecoder
}
