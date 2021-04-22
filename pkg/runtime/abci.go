package runtime

import (
	"github.com/tendermint/tendermint/abci/types"
)

// ABCIRuntime wraps the Runtime as an ABCI basic
type ABCIRuntime struct {
	rt *Runtime
}

func (r ABCIRuntime) Info(info types.RequestInfo) types.ResponseInfo {
	panic("implement me")
}

func (r ABCIRuntime) SetOption(option types.RequestSetOption) types.ResponseSetOption {
	panic("implement me")
}

func (r ABCIRuntime) Query(query types.RequestQuery) types.ResponseQuery {
	panic("implement me")
}

func (r ABCIRuntime) CheckTx(tx types.RequestCheckTx) types.ResponseCheckTx {
	panic("implement me")
}

func (r ABCIRuntime) InitChain(chain types.RequestInitChain) types.ResponseInitChain {
	panic("implement me")
}

func (r ABCIRuntime) BeginBlock(block types.RequestBeginBlock) types.ResponseBeginBlock {
	panic("implement me")
}

func (r ABCIRuntime) DeliverTx(tx types.RequestDeliverTx) types.ResponseDeliverTx {
	panic("implement me")
}

func (r ABCIRuntime) EndBlock(block types.RequestEndBlock) types.ResponseEndBlock {
	panic("implement me")
}

func (r ABCIRuntime) Commit() types.ResponseCommit {
	panic("implement me")
}

func (r ABCIRuntime) ListSnapshots(snapshots types.RequestListSnapshots) types.ResponseListSnapshots {
	panic("implement me")
}

func (r ABCIRuntime) OfferSnapshot(snapshot types.RequestOfferSnapshot) types.ResponseOfferSnapshot {
	panic("implement me")
}

func (r ABCIRuntime) LoadSnapshotChunk(chunk types.RequestLoadSnapshotChunk) types.ResponseLoadSnapshotChunk {
	panic("implement me")
}

func (r ABCIRuntime) ApplySnapshotChunk(chunk types.RequestApplySnapshotChunk) types.ResponseApplySnapshotChunk {
	panic("implement me")
}
