package bank

import (
	"github.com/fdymylja/tmos/runtime/module"
	"github.com/fdymylja/tmos/x/authn/v1alpha1"
	v1alpha12 "github.com/fdymylja/tmos/x/bank/v1alpha1"
)

func NewModule() Module {
	return Module{}
}

type Module struct {
}

func (m Module) Initialize(client module.Client) module.Descriptor {
	return module.NewDescriptorBuilder().
		Named("bank").
		OwnsStateObject(&v1alpha12.Balance{}).
		HandlesStateTransition(&v1alpha12.MsgSendCoins{}, NewSendCoinsHandler(v1alpha12.NewClient(client)), true).
		HandlesStateTransition(&v1alpha12.MsgSetBalance{}, NewSetCoinsHandler(client), false).
		NeedsStateTransition(&v1alpha1.MsgCreateAccount{}).
		WithGenesis(newGenesisHandler()).Build()
}
