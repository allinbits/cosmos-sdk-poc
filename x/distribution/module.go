package distribution

import (
	"github.com/fdymylja/tmos/runtime/module"
	"github.com/fdymylja/tmos/x/bank/v1alpha1"
	extensions2 "github.com/fdymylja/tmos/x/distribution/extensions"
)

func NewModule() Module {
	return Module{}
}

type Module struct {
}

func (m Module) Initialize(client module.Client) module.Descriptor {
	return module.NewDescriptorBuilder().
		Named("distribution").
		ExtendsAuthentication(extensions2.NewAuthentication(client)).
		NeedsStateTransition(&v1alpha1.MsgSendCoins{}).Build()
}
