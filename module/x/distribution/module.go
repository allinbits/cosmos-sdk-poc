package distribution

import (
	bank "github.com/fdymylja/tmos/module/x/bank/v1alpha1"
	"github.com/fdymylja/tmos/module/x/distribution/extensions"
	"github.com/fdymylja/tmos/runtime/module"
)

func NewModule() Module {
	return Module{}
}

type Module struct {
}

func (m Module) Initialize(client module.Client) module.Descriptor {
	return module.NewDescriptorBuilder().
		Named("distribution").
		ExtendsAuthentication(extensions.NewAuthentication(client)).
		NeedsStateTransition(&bank.MsgSendCoins{}).Build()
}
