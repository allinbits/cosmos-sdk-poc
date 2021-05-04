package authn

import (
	"encoding/json"

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
		ExtendsAuthentication(extensions.New(c)).
		WithGenesis(genesis{c: c})
}

func (m *Module) GetAuthenticator() authentication.Authenticator {
	return m.authenticator
}

type genesis struct {
	c module.Client
}

func (g genesis) SetDefault() error {
	err := g.c.Create(&v1alpha1.Params{
		MaxMemoCharacters:      1000000,
		TxSigLimit:             1000000,
		TxSizeCostPerByte:      0,
		SigVerifyCostEd25519:   0,
		SigVerifyCostSecp256K1: 0,
	})
	if err != nil {
		return err
	}
	err = g.c.Create(&v1alpha1.CurrentAccountNumber{Number: 0})
	if err != nil {
		return err
	}
	return err
}

func (g genesis) Import(state json.RawMessage) error {
	panic("implement me")
}

func (g genesis) Export(state json.RawMessage) error {
	panic("implement me")
}
