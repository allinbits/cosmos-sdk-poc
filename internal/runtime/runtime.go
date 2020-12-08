package runtime

import (
	store2 "github.com/cosmos/cosmos-sdk/store"
	"github.com/cosmos/cosmos-sdk/store/iavl"
	"github.com/fdymylja/cosmos-os/pkg/apis/core/v1alpha1"
	"github.com/fdymylja/cosmos-os/pkg/application"
	"github.com/fdymylja/cosmos-os/pkg/codec"
	"github.com/tendermint/tendermint/abci/types"
	dbm "github.com/tendermint/tm-db"
	"log"
)

func NewRuntime() Runtime {
	store, err := iavl.LoadStore(dbm.NewMemDB(), store2.CommitID{}, true)
	if err != nil {
		panic(err)
	}
	return Runtime{store: store.(*iavl.Store)}
}

type Runtime struct {
	store *iavl.Store
	cdc codec.Codec
	deliverRouter map[string]deliverer
	checkRouter map[string]checker
	queryRouter map[string]querier
}

func (r Runtime) Info(info types.RequestInfo) types.ResponseInfo {
	panic("implement me")
}

func (r Runtime) SetOption(option types.RequestSetOption) types.ResponseSetOption {
	panic("implement me")
}

func (r Runtime) Query(rawQuery types.RequestQuery) types.ResponseQuery {
	if !r.store.VersionExists(rawQuery.Height) {
		panic("store version does not exist")
	}
	hStore, err := r.store.GetImmutable(rawQuery.Height)
	if err != nil {
		panic(err)
	}
	query := new(v1alpha1.Query)
	err = r.cdc.Unmarshal(rawQuery.Data, query)
	if err != nil {
		panic(err)
	}
	name := query.Query.TypeUrl
	querier, exists := r.queryRouter[name]
	if !exists {
		panic("not exist duh")
	}
	store := newAppStore(r.cdc, hStore, querier.applicationID)
	resp, err := querier.do(application.QueryRequest{
		Request: query.Query.Value,
		Client:  newAppClient(r.cdc, hStore, r.deliverRouter, r.queryRouter),
		DB:      store,
	})
	if err != nil {
		panic(err)
	}
	return types.ResponseQuery{
		Value:     resp.Response,
	}
}

func (r Runtime) CheckTx(tx types.RequestCheckTx) types.ResponseCheckTx {
	panic("implement me")
}

func (r Runtime) InitChain(chain types.RequestInitChain) types.ResponseInitChain {
	panic("implement me")
}

func (r Runtime) BeginBlock(block types.RequestBeginBlock) types.ResponseBeginBlock {
	panic("implement me")
}

func (r Runtime) DeliverTx(tmTx types.RequestDeliverTx) types.ResponseDeliverTx {
	tx := new(v1alpha1.Transaction)
	err := r.cdc.Unmarshal(tmTx.Tx, tx)
	if err != nil {
		panic(err)
	}

	objectName := tx.Messages.TypeUrl
	deliverer, exists := r.deliverRouter[objectName]
	if !exists {
		panic(err)
	}
	resp, err := deliverer.do(application.DeliverRequest{
		Request: tx.Messages.Value,
		Client:  newAppClient(r.cdc, r.store, r.deliverRouter, r.queryRouter),
		Store:   newAppStore(r.cdc, r.store, deliverer.applicationID),
	})
	log.Printf("%#v", resp)
	if err != nil {
		panic(err)
	}
	return types.ResponseDeliverTx{}
}

func (r Runtime) EndBlock(block types.RequestEndBlock) types.ResponseEndBlock {
	panic("implement me")
}

func (r Runtime) Commit() types.ResponseCommit {
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

