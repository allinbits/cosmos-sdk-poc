package bank

import (
	"github.com/fdymylja/tmos/module/x/bank/v1alpha1"
	"github.com/fdymylja/tmos/runtime/module"
)

func NewModule() Module {
	return Module{}
}

type Module struct {
}

func (m Module) Initialize(client module.Client, builder *module.Builder) {
	builder.
		Named("bank").
		OwnsStateObject(&v1alpha1.Balance{}).
		HandlesStateTransition(&v1alpha1.MsgSendCoins{}, NewSendCoinsHandler(client))
}
