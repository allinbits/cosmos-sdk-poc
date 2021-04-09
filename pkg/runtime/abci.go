package runtime

import (
	"github.com/tendermint/tendermint/abci/types"
	"google.golang.org/protobuf/proto"

	runtime "github.com/fdymylja/tmos/apis/core/runtime/v1alpha1"
)

func (r *Runtime) Info(info types.RequestInfo) types.ResponseInfo {
	panic("implement me")
}

func (r *Runtime) SetOption(option types.RequestSetOption) types.ResponseSetOption {
	panic("implement me")
}

func (r *Runtime) Query(query types.RequestQuery) types.ResponseQuery {
	panic("implement me")
}

func (r *Runtime) CheckTx(tx types.RequestCheckTx) types.ResponseCheckTx {
	panic("implement me")
}

func (r *Runtime) InitChain(chain types.RequestInitChain) types.ResponseInitChain {
	panic("implement me")
}

func (r *Runtime) BeginBlock(block types.RequestBeginBlock) types.ResponseBeginBlock {
	panic("implement me")
}

func (r *Runtime) DeliverTx(tmTx types.RequestDeliverTx) types.ResponseDeliverTx {
	rawTx := new(runtime.RawTx)
	err := proto.Unmarshal(tmTx.Tx, rawTx)
	if err != nil {
		panic(err)
	}
	// after we decode the tx we need to find the concrete type of this tx
	stateTransitions, err := r.decodeTx(rawTx)
	if err != nil {
		panic(err)
	}
	// after we get the state transitions we route them to the correct handlers
	for _, stateTransition := range stateTransitions {
		err = r.routeTransition(stateTransition)
		if err != nil {
			panic(err)
		}
	}
	return types.ResponseDeliverTx{}
}

func (r *Runtime) EndBlock(block types.RequestEndBlock) types.ResponseEndBlock {
	panic("implement me")
}

func (r *Runtime) Commit() types.ResponseCommit {
	panic("implement me")
}

func (r Runtime) ListSnapshots(snapshots types.RequestListSnapshots) types.ResponseListSnapshots {
	panic("implement me")
}

func (r Runtime) OfferSnapshot(snapshot types.RequestOfferSnapshot) types.ResponseOfferSnapshot {
	panic("implement me")
}

func (r Runtime) LoadSnapshotChunk(chunk types.RequestLoadSnapshotChunk) types.ResponseLoadSnapshotChunk {
	panic("implement me")
}

func (r Runtime) ApplySnapshotChunk(chunk types.RequestApplySnapshotChunk) types.ResponseApplySnapshotChunk {
	panic("implement me")
}
