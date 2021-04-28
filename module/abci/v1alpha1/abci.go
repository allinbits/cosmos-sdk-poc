package v1alpha1

import "github.com/fdymylja/tmos/runtime/meta"

var (
	StageID           = meta.NewStringID("stage")
	BeginBlockStateID = meta.NewStringID("begin_block")
	CheckTxStateID    = meta.NewStringID("check_tx")
	DeliverTxStateID  = meta.NewStringID("deliver_tx")
)

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

func (x *MsgSetCheckTxState) StateTransition()    {}
func (x *MsgSetBeginBlockState) StateTransition() {}
func (x *MsgSetDeliverTxState) StateTransition()  {}
