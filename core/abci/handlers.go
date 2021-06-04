package abci

import (
	"fmt"

	client2 "github.com/fdymylja/tmos/runtime/client"

	"github.com/fdymylja/tmos/core/abci/tendermint/abci"
	"github.com/fdymylja/tmos/core/abci/v1alpha1"
	meta "github.com/fdymylja/tmos/core/meta"
	"github.com/fdymylja/tmos/runtime/statetransition"
)

func setInitChainInfo() statetransition.ExecutionHandlerFunc {
	return func(client client2.RuntimeClient, req statetransition.ExecutionRequest) (resp statetransition.ExecutionResponse, err error) {
		msg := req.Transition.(*v1alpha1.MsgSetInitChain)
		// set init chain info
		err = client.Update(&v1alpha1.InitChainInfo{ChainId: msg.InitChainInfo.ChainId})
		if err != nil {
			return statetransition.ExecutionResponse{}, err
		}
		// set stage init chain
		err = client.Update(&v1alpha1.Stage{Stage: v1alpha1.ABCIStage_InitChain})
		if err != nil {
			return statetransition.ExecutionResponse{}, err
		}
		return statetransition.ExecutionResponse{}, nil
	}
}

func checkTxHandler() statetransition.ExecutionHandlerFunc {
	return func(client client2.RuntimeClient, req statetransition.ExecutionRequest) (resp statetransition.ExecutionResponse, err error) {
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

func beginBlockHandler() statetransition.ExecutionHandlerFunc {
	return func(client client2.RuntimeClient, req statetransition.ExecutionRequest) (resp statetransition.ExecutionResponse, err error) {
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

func endBlockHandler() statetransition.ExecutionHandlerFunc {
	return func(client client2.RuntimeClient, req statetransition.ExecutionRequest) (resp statetransition.ExecutionResponse, err error) {
		msg := req.Transition.(*v1alpha1.MsgSetEndBlockState)
		err = client.Update(&v1alpha1.Stage{Stage: v1alpha1.ABCIStage_EndBlock})
		if err != nil {
			return statetransition.ExecutionResponse{}, err
		}
		err = client.Update(&v1alpha1.EndBlockState{
			EndBlock: msg.EndBlock,
		})
		if err != nil {
			return statetransition.ExecutionResponse{}, err
		}
		return
	}
}

func deliverTxHandler() statetransition.ExecutionHandlerFunc {
	return func(client client2.RuntimeClient, req statetransition.ExecutionRequest) (resp statetransition.ExecutionResponse, err error) {
		msg := req.Transition.(*v1alpha1.MsgSetDeliverTxState)
		err = client.Update(&v1alpha1.Stage{Stage: v1alpha1.ABCIStage_DeliverTx})
		if err != nil {
			return
		}
		err = client.Update(&v1alpha1.DeliverTxState{DeliverTx: msg.DeliverTx})
		return
	}
}

func validatorUpdatesHandler() statetransition.ExecutionHandlerFunc {
	return func(client client2.RuntimeClient, req statetransition.ExecutionRequest) (resp statetransition.ExecutionResponse, err error) {
		msg := req.Transition.(*v1alpha1.MsgSetValidatorUpdates)
		stage := new(v1alpha1.Stage)
		err = client.Get(meta.SingletonID, stage)
		if err != nil {
			return statetransition.ExecutionResponse{}, err
		}
		// this handler can only be executed during begin and endblock stages
		if stage.Stage != v1alpha1.ABCIStage_InitChain && stage.Stage != v1alpha1.ABCIStage_EndBlock {
			return statetransition.ExecutionResponse{}, fmt.Errorf("validator updates can be done only during EndBlock and InitChain, stage is %s", stage.Stage)
		}
		// update validator set
		err = client.Update(&v1alpha1.ValidatorUpdates{ValidatorUpdates: msg.ValidatorUpdates})
		if err != nil {
			return statetransition.ExecutionResponse{}, err
		}
		return
	}
}
