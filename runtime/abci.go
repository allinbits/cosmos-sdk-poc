package runtime

import (
	"github.com/fdymylja/tmos/module/abci/tendermint/abci"
	abcictrl "github.com/fdymylja/tmos/module/abci/v1alpha1"
	"github.com/tendermint/tendermint/abci/types"
)

func NewABCIApplication(rt *Runtime) ABCIApplication {
	return ABCIApplication{rt: rt}
}

// ABCIApplication is a Runtime orchestrated by Tendermint
type ABCIApplication struct {
	rt *Runtime
}

func (a ABCIApplication) Info(info types.RequestInfo) types.ResponseInfo {
	return types.ResponseInfo{
		Data:             "runtime",
		Version:          "0.0.0",
		AppVersion:       0,
		LastBlockHeight:  0,
		LastBlockAppHash: []byte("unknown"),
	}
}

func (a ABCIApplication) SetOption(option types.RequestSetOption) types.ResponseSetOption {
	panic("implement me")
}

func (a ABCIApplication) Query(query types.RequestQuery) types.ResponseQuery {
	panic("implement me")
}

func (a ABCIApplication) CheckTx(tmTx types.RequestCheckTx) types.ResponseCheckTx {
	// we decode the tx and run admission checks
	tx, err := a.rt.authn.DecodeTx(tmTx.Tx)
	if err != nil {
		return types.ResponseCheckTx{Code: 1}
	}
	// run admission checks on transitions
	for _, transition := range tx.StateTransitions() {
		err = a.rt.runAdmissionChain(transition)
		if err != nil {
			return types.ResponseCheckTx{Code: 1}
		}
	}
	return types.ResponseCheckTx{}
}

func (a ABCIApplication) InitChain(chain types.RequestInitChain) types.ResponseInitChain {
	// if app state bytes is nil we initialize with default genesis states
	switch len(chain.AppStateBytes) {
	case 0:
		err := a.rt.Initialize()
		if err != nil {
			panic(err)
		}
	// otherwise we import the chain from the state bytes
	default:
		err := a.rt.Import(chain.AppStateBytes)
		if err != nil {
			panic(err)
		}
	}
	// set init chain info
	err := a.rt.Deliver(nil, &abcictrl.MsgSetInitChain{InitChainInfo: &abcictrl.InitChainInfo{ChainId: chain.ChainId}})
	if err != nil {
		panic(err)
	}
	//
	return types.ResponseInitChain{
		ConsensusParams: nil,
		Validators:      chain.Validators,
		AppHash:         nil,
	}
}

func (a ABCIApplication) BeginBlock(tmBlock types.RequestBeginBlock) types.ResponseBeginBlock {
	block := new(abci.RequestBeginBlock)
	block.FromLegacyProto(&tmBlock)
	err := a.rt.Deliver([]string{"abci"}, &abcictrl.MsgSetBeginBlockState{BeginBlock: block})
	if err != nil {
		panic(err)
	}
	// return
	return types.ResponseBeginBlock{
		Events: nil,
	}
}

func (a ABCIApplication) DeliverTx(tmTx types.RequestDeliverTx) types.ResponseDeliverTx {
	// decode tx
	tx, err := a.rt.authn.DecodeTx(tmTx.Tx)
	if err != nil {
		return ToABCIResponse(0, 0, err)
	}
	// run admission checks on tx
	err = a.rt.runTxAdmissionChain(tx)
	if err != nil {
		return ToABCIResponse(0, 0, err)
	}
	// here we cache the store
	// todo authenticate
	err = a.rt.authn.Authenticate(tx)
	if err != nil {
		return ToABCIResponse(0, 0, err)
	}
	// todo run authentication chain
	// write the cache
	// cache again
	// start delivering transitions
	for _, transition := range tx.StateTransitions() {
		err = a.rt.Deliver(nil, transition)
		if err != nil {
			return ToABCIResponse(0, 0, err)
		}
	}
	// write cache
	// success!
	return types.ResponseDeliverTx{}
}

func (a ABCIApplication) EndBlock(block types.RequestEndBlock) types.ResponseEndBlock {
	panic("implement me")
}

func (a ABCIApplication) Commit() types.ResponseCommit {
	panic("implement me")
}

func (a ABCIApplication) ListSnapshots(snapshots types.RequestListSnapshots) types.ResponseListSnapshots {
	panic("implement me")
}

func (a ABCIApplication) OfferSnapshot(snapshot types.RequestOfferSnapshot) types.ResponseOfferSnapshot {
	panic("implement me")
}

func (a ABCIApplication) LoadSnapshotChunk(chunk types.RequestLoadSnapshotChunk) types.ResponseLoadSnapshotChunk {
	panic("implement me")
}

func (a ABCIApplication) ApplySnapshotChunk(chunk types.RequestApplySnapshotChunk) types.ResponseApplySnapshotChunk {
	panic("implement me")
}
