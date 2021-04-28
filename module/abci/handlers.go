package abci

import (
	"fmt"

	"github.com/fdymylja/tmos/module/abci/tendermint/abci"
	"github.com/fdymylja/tmos/module/abci/v1alpha1"
	"github.com/fdymylja/tmos/runtime"
	"github.com/fdymylja/tmos/runtime/controller"
)

func checkTxHandler(client runtime.ModuleClient) controller.StateTransitionFn {
	return func(req controller.StateTransitionRequest) (resp controller.StateTransitionResponse, err error) {
		msg := req.Transition.(*v1alpha1.MsgSetCheckTxState)
		switch msg.CheckTx.Type {
		case abci.CheckTxType_NEW:
			err = client.Update(&v1alpha1.Stage{Stage: v1alpha1.ABCIStage_CheckTx})
			if err != nil {
				return
			}
		case abci.CheckTxType_RECHECK:
			err = client.Update(&v1alpha1.Stage{Stage: v1alpha1.ABCIStage_ReCheckTx})
			if err != nil {
				return
			}
		default:
			panic(fmt.Errorf("unknown checktx type: %s", msg.CheckTx.Type))
		}
		err = client.Update(&v1alpha1.CheckTxState{CheckTx: msg.CheckTx})
		return
	}
}

func beginBlockHandler(client runtime.ModuleClient) controller.StateTransitionFn {
	return func(req controller.StateTransitionRequest) (resp controller.StateTransitionResponse, err error) {
		msg := req.Transition.(*v1alpha1.MsgSetBeginBlockState)
		err = client.Update(&v1alpha1.Stage{Stage: v1alpha1.ABCIStage_BeginBlock})
		if err != nil {
			return
		}
		err = client.Update(&v1alpha1.BeginBlockState{BeginBlock: msg.BeginBlock})
		return
	}
}

func deliverTxHandler(client runtime.ModuleClient) controller.StateTransitionFn {
	return func(req controller.StateTransitionRequest) (resp controller.StateTransitionResponse, err error) {
		msg := req.Transition.(*v1alpha1.MsgSetDeliverTxState)
		err = client.Update(&v1alpha1.Stage{Stage: v1alpha1.ABCIStage_DeliverTx})
		if err != nil {
			return
		}
		err = client.Update(&v1alpha1.DeliverTxState{DeliverTx: msg.DeliverTx})
		return
	}
}
