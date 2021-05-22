package bank

import (
	"github.com/fdymylja/tmos/runtime/module"
	"github.com/fdymylja/tmos/runtime/orm"
	authnv1alpha1 "github.com/fdymylja/tmos/x/authn/v1alpha1"
	"github.com/fdymylja/tmos/x/bank/v1alpha1"
)

func NewModule() Module {
	return Module{}
}

type Module struct {
}

func (m Module) Initialize(client module.Client) module.Descriptor {
	return module.NewDescriptorBuilder().
		Named("bank").
		OwnsStateObject(&v1alpha1.Balance{}, orm.RegisterOptions{PrimaryKey: "address"}).
		HandlesStateTransition(&v1alpha1.MsgSendCoins{}, NewSendCoinsHandler(v1alpha1.NewClient(client)), true).
		HandlesStateTransition(&v1alpha1.MsgSetBalance{}, NewSetCoinsHandler(client), false).
		NeedsStateTransition(&authnv1alpha1.MsgCreateAccount{}).
		WithGenesis(newGenesisHandler()).Build()
}
