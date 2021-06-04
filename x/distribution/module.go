package distribution

import (
	"github.com/fdymylja/tmos/runtime/module"
	"github.com/fdymylja/tmos/x/bank/v1alpha1"
)

func NewModule() Module {
	return Module{}
}

type Module struct {
}

func (m Module) Initialize(client module.Client) module.Descriptor {
	return module.NewDescriptorBuilder().
		Named("distribution").
		WithAuthAdmissionHandler(NewFeeChecker()).
		WithPostAuthenticationHandler(NewFeeDeduction(client)).
		NeedsStateTransition(&v1alpha1.MsgSendCoins{}).Build()
}
