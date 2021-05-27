package authn

import (
	rbacv1alpha1 "github.com/fdymylja/tmos/core/rbac/v1alpha1"
	"github.com/fdymylja/tmos/runtime/authentication"
	"github.com/fdymylja/tmos/runtime/module"
	extensions2 "github.com/fdymylja/tmos/x/authn/extensions"
	"github.com/fdymylja/tmos/x/authn/v1alpha1"
)

// Module implements the authentication.Module
type Module struct {
	txDecoder authentication.TxDecoder
}

func NewModule() *Module {
	return &Module{}
}

func (m *Module) Initialize(c module.Client) module.Descriptor {
	m.txDecoder = newTxDecoder()

	return module.NewDescriptorBuilder().
		Named("authn").
		HandlesStateTransition(&v1alpha1.MsgCreateAccount{}, NewCreateAccountController(c), true).
		HandlesAdmission(&v1alpha1.MsgCreateAccount{}, NewCreateAccountAdmissionController()).
		OwnsStateObject(&v1alpha1.Account{}, v1alpha1.AccountSchema).
		OwnsStateObject(&v1alpha1.Params{}, v1alpha1.ParamsSchema).
		OwnsStateObject(&v1alpha1.CurrentAccountNumber{}, v1alpha1.CurrentAccountNumberSchema).
		ExtendsAuthentication(extensions2.New(c)).
		WithGenesis(genesis{c: c}).
		NeedsStateTransition(&rbacv1alpha1.MsgBindRole{}).Build()
}

func (m *Module) GetTxDecoder() authentication.TxDecoder {
	return m.txDecoder
}
