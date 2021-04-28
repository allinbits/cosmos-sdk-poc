package abci

import (
	"github.com/fdymylja/tmos/module/abci/v1alpha1"
	"github.com/fdymylja/tmos/runtime"
)

func NewModule() runtime.Module {
	return Module{}
}

type Module struct {
}

func (m Module) Initialize(client runtime.ModuleClient, builder *runtime.ModuleBuilder) {
	builder.
		Named("abci").
		OwnsStateObject(&v1alpha1.Stage{}).
		OwnsStateObject(&v1alpha1.BeginBlockState{}).
		OwnsStateObject(&v1alpha1.DeliverTxState{}).
		OwnsStateObject(&v1alpha1.CheckTxState{}).
		HandlesStateTransition(&v1alpha1.MsgSetCheckTxState{}, checkTxHandler(client)).
		HandlesStateTransition(&v1alpha1.MsgSetBeginBlockState{}, beginBlockHandler(client)).
		HandlesStateTransition(&v1alpha1.MsgSetDeliverTxState{}, deliverTxHandler(client)).
		WithGenesis(newGenesisHandler(client))
}
