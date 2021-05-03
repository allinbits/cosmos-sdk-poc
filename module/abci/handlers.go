package abci

import (
	"fmt"

	"github.com/fdymylja/tmos/module/abci/tendermint/abci"
	"github.com/fdymylja/tmos/module/abci/v1alpha1"
	"github.com/fdymylja/tmos/runtime/controller"
	"github.com/fdymylja/tmos/runtime/module"
)

func setInitChainInfo(client module.Client) controller.StateTransitionFn {
	return func(req controller.StateTransitionRequest) (resp controller.StateTransitionResponse, err error) {
		msg := req.Transition.(*v1alpha1.MsgSetInitChain)
		err = client.Update(&v1alpha1.InitChainInfo{ChainId: msg.InitChainInfo.ChainId})
		if err != nil {
			return controller.StateTransitionResponse{}, err
		}
		return controller.StateTransitionResponse{}, nil
	}
}

func checkTxHandler(client module.Client) controller.StateTransitionFn {
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

func beginBlockHandler(client module.Client) controller.StateTransitionFn {
	return func(req controller.StateTransitionRequest) (resp controller.StateTransitionResponse, err error) {
		msg := req.Transition.(*v1alpha1.MsgSetBeginBlockState)
		err = client.Update(&v1alpha1.Stage{Stage: v1alpha1.ABCIStage_BeginBlock})
		if err != nil {
			return
		}
		err = client.Update(&v1alpha1.BeginBlockState{BeginBlock: msg.BeginBlock})
		if err != nil {
			return
		}
		err = client.Update(&v1alpha1.CurrentBlock{BlockNumber: uint64(msg.BeginBlock.Header.Height)})
		return
	}
}

func deliverTxHandler(client module.Client) controller.StateTransitionFn {
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
