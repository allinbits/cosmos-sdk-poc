package abci

import (
	"github.com/fdymylja/tmos/module/abci/v1alpha1"
	"github.com/fdymylja/tmos/runtime/module"
)

func NewModule() module.Module {
	return Module{}
}

type Module struct {
}

func (m Module) Initialize(client module.Client, builder *module.Builder) {
	builder.
		Named("abci").
		OwnsStateObject(&v1alpha1.Stage{}).
		OwnsStateObject(&v1alpha1.BeginBlockState{}).
		OwnsStateObject(&v1alpha1.DeliverTxState{}).
		OwnsStateObject(&v1alpha1.CheckTxState{}).
		OwnsStateObject(&v1alpha1.CurrentBlock{}).
		OwnsStateObject(&v1alpha1.InitChainInfo{}).
		HandlesStateTransition(&v1alpha1.MsgSetInitChain{}, setInitChainInfo(client)).
		HandlesStateTransition(&v1alpha1.MsgSetCheckTxState{}, checkTxHandler(client)).
		HandlesStateTransition(&v1alpha1.MsgSetBeginBlockState{}, beginBlockHandler(client)).
		HandlesStateTransition(&v1alpha1.MsgSetDeliverTxState{}, deliverTxHandler(client)).
		WithGenesis(newGenesisHandler(client))
}
