package abci

import (
	"github.com/fdymylja/tmos/core/abci/v1alpha1"
	"github.com/fdymylja/tmos/runtime/authentication/user"
	"github.com/fdymylja/tmos/runtime/module"
)

func NewModule() module.Module {
	return Module{}
}

type Module struct {
}

func (m Module) Initialize(client module.Client) module.Descriptor {
	return module.NewDescriptorBuilder().
		Named(user.ABCI).
		OwnsStateObject(&v1alpha1.Stage{}).
		OwnsStateObject(&v1alpha1.BeginBlockState{}).
		OwnsStateObject(&v1alpha1.DeliverTxState{}).
		OwnsStateObject(&v1alpha1.CheckTxState{}).
		OwnsStateObject(&v1alpha1.InitChainInfo{}).
		OwnsStateObject(&v1alpha1.EndBlockState{}).
		OwnsStateObject(&v1alpha1.CurrentBlock{}).
		OwnsStateObject(&v1alpha1.ValidatorUpdates{}).
		HandlesStateTransition(&v1alpha1.MsgSetInitChain{}, setInitChainInfo(client), false).
		HandlesStateTransition(&v1alpha1.MsgSetCheckTxState{}, checkTxHandler(client), false).
		HandlesStateTransition(&v1alpha1.MsgSetBeginBlockState{}, beginBlockHandler(client), false).
		HandlesStateTransition(&v1alpha1.MsgSetDeliverTxState{}, deliverTxHandler(client), false).
		HandlesStateTransition(&v1alpha1.MsgSetEndBlockState{}, endBlockHandler(client), false).
		HandlesStateTransition(&v1alpha1.MsgSetValidatorUpdates{}, validatorUpdatesHandler(client), false).
		WithGenesis(newGenesisHandler(client)).Build()
}