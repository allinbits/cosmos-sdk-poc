package v1alpha1

import "github.com/fdymylja/tmos/runtime/meta"

var (
	StageID            = meta.NewStringID("stage")
	InitChainInfoID    = meta.NewStringID("chain_id")
	CurrentBlockID     = meta.NewStringID("current_block")
	BeginBlockStateID  = meta.NewStringID("begin_block")
	CheckTxStateID     = meta.NewStringID("check_tx")
	DeliverTxStateID   = meta.NewStringID("deliver_tx")
	EndBlockStateID    = meta.NewStringID("end_block")
	ValidatorUpdatesID = meta.NewStringID("validator_updates")
)

func (x *ValidatorUpdates) GetID() meta.ID {
	return ValidatorUpdatesID
}

func (x *InitChainInfo) GetID() meta.ID {
	return InitChainInfoID
}

func (x *EndBlockState) GetID() meta.ID {
	return EndBlockStateID
}

func (x *Stage) GetID() meta.ID {
	return StageID
}

func (x *BeginBlockState) GetID() meta.ID {
	return BeginBlockStateID
}

func (x *DeliverTxState) GetID() meta.ID {
	return DeliverTxStateID
}

func (x *CheckTxState) GetID() meta.ID {
	return CheckTxStateID
}

func (x *CurrentBlock) GetID() meta.ID {
	return CurrentBlockID
}

func (x *MsgSetCheckTxState) StateTransition()     {}
func (x *MsgSetBeginBlockState) StateTransition()  {}
func (x *MsgSetDeliverTxState) StateTransition()   {}
func (x *MsgSetInitChain) StateTransition()        {}
func (x *MsgSetEndBlockState) StateTransition()    {}
func (x *MsgSetValidatorUpdates) StateTransition() {}
