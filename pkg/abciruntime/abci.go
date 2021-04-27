package abciruntime

import (
	rtabci "github.com/fdymylja/tmos/apis/abci/v1alpha1"
	"github.com/fdymylja/tmos/pkg/runtime"
	abci "github.com/tendermint/tendermint/abci/types"
)

// ABCIIdentity defines the unique identifier for the user ABCI in the runtime
const ABCIIdentity = "abci"

// ABCIRuntime wraps the Runtime as an ABCI application
type ABCIRuntime struct {
	rt *runtime.Runtime
}

func (r ABCIRuntime) Info(info abci.RequestInfo) abci.ResponseInfo {
	panic("implement me")
}

func (r ABCIRuntime) SetOption(option abci.RequestSetOption) abci.ResponseSetOption {
	panic("implement me")
}

func (r ABCIRuntime) Query(query abci.RequestQuery) abci.ResponseQuery {
	panic("implement me")
}

func (r ABCIRuntime) CheckTx(tx abci.RequestCheckTx) abci.ResponseCheckTx {
	panic("implement me")
}

func (r ABCIRuntime) InitChain(req abci.RequestInitChain) abci.ResponseInitChain {
	panic("implement me")
}

// BeginBlock implements the tendermint abci application interface
func (r ABCIRuntime) BeginBlock(block abci.RequestBeginBlock) abci.ResponseBeginBlock {
	// we deliver a request on behalf of ABCI to the runtime
	// so that begin block information is stored and provided
	// to controllers which wish to access the given data.
	err := r.rt.Deliver([]string{ABCIIdentity}, &rtabci.MsgSetBeginBlockState{BeginBlock: convBeginBlock(block)})
	if err != nil {
		panic(err)
	}
	// TODO handle begin block
	return abci.ResponseBeginBlock{
		Events: nil,
	}
}

func (r ABCIRuntime) DeliverTx(tx abci.RequestDeliverTx) abci.ResponseDeliverTx {
	panic("implement me")
}

func (r ABCIRuntime) EndBlock(block abci.RequestEndBlock) abci.ResponseEndBlock {
	panic("implement me")
}

func (r ABCIRuntime) Commit() abci.ResponseCommit {
	panic("implement me")
}

func (r ABCIRuntime) ListSnapshots(snapshots abci.RequestListSnapshots) abci.ResponseListSnapshots {
	panic("implement me")
}

func (r ABCIRuntime) OfferSnapshot(snapshot abci.RequestOfferSnapshot) abci.ResponseOfferSnapshot {
	panic("implement me")
}

func (r ABCIRuntime) LoadSnapshotChunk(chunk abci.RequestLoadSnapshotChunk) abci.ResponseLoadSnapshotChunk {
	panic("implement me")
}

func (r ABCIRuntime) ApplySnapshotChunk(chunk abci.RequestApplySnapshotChunk) abci.ResponseApplySnapshotChunk {
	panic("implement me")
}
