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
		OwnsStateObject(&v1alpha1.Stage{}, v1alpha1.StageSchema).
		OwnsStateObject(&v1alpha1.BeginBlockState{}, v1alpha1.BeginBlockStateSchema).
		OwnsStateObject(&v1alpha1.DeliverTxState{}, v1alpha1.DeliverTxStateSchema).
		OwnsStateObject(&v1alpha1.CheckTxState{}, v1alpha1.CheckTxStateSchema).
		OwnsStateObject(&v1alpha1.InitChainInfo{}, v1alpha1.InitChainInfoSchema).
		OwnsStateObject(&v1alpha1.EndBlockState{}, v1alpha1.EndBlockStateSchema).
		OwnsStateObject(&v1alpha1.CurrentBlock{}, v1alpha1.CurrentBlockSchema).
		OwnsStateObject(&v1alpha1.ValidatorUpdates{}, v1alpha1.ValidatorUpdatesSchema).
		HandlesStateTransition(&v1alpha1.MsgSetInitChain{}, setInitChainInfo(), false).
		HandlesStateTransition(&v1alpha1.MsgSetCheckTxState{}, checkTxHandler(), false).
		HandlesStateTransition(&v1alpha1.MsgSetBeginBlockState{}, beginBlockHandler(), false).
		HandlesStateTransition(&v1alpha1.MsgSetDeliverTxState{}, deliverTxHandler(), false).
		HandlesStateTransition(&v1alpha1.MsgSetEndBlockState{}, endBlockHandler(), false).
		HandlesStateTransition(&v1alpha1.MsgSetValidatorUpdates{}, validatorUpdatesHandler(), false).
		WithGenesis(newGenesisHandler()).Build()
}
