package runtime

import (
	"github.com/fdymylja/tmos/core/abci/tendermint/abci"
	abcictrl "github.com/fdymylja/tmos/core/abci/v1alpha1"
	meta "github.com/fdymylja/tmos/core/meta"
	"github.com/fdymylja/tmos/runtime/authentication/user"
	"github.com/fdymylja/tmos/runtime/errors"
	"github.com/tendermint/tendermint/abci/types"
)

func NewABCIApplication(rt *Runtime) ABCIApplication {
	return ABCIApplication{
		rt:    rt,
		user:  user.NewUsersFromString(user.ABCI),
		query: NewQuerier(rt.store),
	}
}

// ABCIApplication is a Runtime orchestrated by Tendermint
type ABCIApplication struct {
	rt    *Runtime
	user  user.Users
	query *Querier
}

func (a ABCIApplication) Info(info types.RequestInfo) types.ResponseInfo {
	return types.ResponseInfo{
		Data:             "runtime",
		Version:          "v0.0.0",
		AppVersion:       0,
		LastBlockHeight:  0,
		LastBlockAppHash: []byte("unknown"),
	}
}

func (a ABCIApplication) SetOption(option types.RequestSetOption) types.ResponseSetOption {
	panic("implement me")
}

func (a ABCIApplication) Query(query types.RequestQuery) types.ResponseQuery {
	return a.query.Handle(query)
}

func (a ABCIApplication) CheckTx(tmTx types.RequestCheckTx) types.ResponseCheckTx {
	// decode tx
	tx, err := a.rt.txDecoder.DecodeTx(tmTx.Tx)
	if err != nil {
		return types.ResponseCheckTx(errors.ToABCIResponse(0, 0, err))
	}
	// run admission checks on the transaction
	err = a.rt.runTxAdmissionChain(tx)
	if err != nil {
		return types.ResponseCheckTx(errors.ToABCIResponse(0, 0, err))
	}
	// run admission checks on single state transitions
	for _, transition := range tx.StateTransitions() {
		err = a.rt.runAdmissionChain(tx.Users(), transition)
		if err != nil {
			return types.ResponseCheckTx(errors.ToABCIResponse(0, 0, err))
		}
	}
	return types.ResponseCheckTx{}
}

func (a ABCIApplication) InitChain(chain types.RequestInitChain) types.ResponseInitChain {
	// if app state bytes is nil we initialize with default genesis states
	switch len(chain.AppStateBytes) {
	case 0:
		err := a.rt.InitGenesis()
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
	err := a.rt.Deliver(a.user, &abcictrl.MsgSetInitChain{InitChainInfo: &abcictrl.InitChainInfo{ChainId: chain.ChainId}})
	if err != nil {
		panic(err)
	}
	// enable role based access control
	a.rt.EnableRBAC()
	return types.ResponseInitChain{
		ConsensusParams: nil,
		Validators:      chain.Validators,
		AppHash:         nil,
	}
}

func (a ABCIApplication) BeginBlock(tmBlock types.RequestBeginBlock) types.ResponseBeginBlock {
	block := new(abci.RequestBeginBlock)
	block.FromLegacyProto(&tmBlock)
	err := a.rt.Deliver(a.user, &abcictrl.MsgSetBeginBlockState{BeginBlock: block})
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
	tx, err := a.rt.txDecoder.DecodeTx(tmTx.Tx)
	if err != nil {
		return errors.ToABCIResponse(0, 0, err)
	}
	// run admission checks on tx
	err = a.rt.runTxAdmissionChain(tx)
	if err != nil {
		return errors.ToABCIResponse(0, 0, err)
	}
	// TODO cache the store
	err = a.rt.runTxPostAuthenticationChain(tx)
	if err != nil {
		return errors.ToABCIResponse(0, 0, err)
	}
	// TODO write the cache
	// TODO cache again
	// start delivering transitions
	for _, transition := range tx.StateTransitions() {
		err = a.rt.Deliver(tx.Users(), transition, DeliverSkipAdmissionHandlers())
		if err != nil {
			return errors.ToABCIResponse(0, 0, err)
		}
	}
	// TODO write cache
	// success!
	return types.ResponseDeliverTx{}
}

func (a ABCIApplication) EndBlock(block types.RequestEndBlock) types.ResponseEndBlock {
	// we set the abci endblcok state given us by tendermint so other modules can access it
	err := a.rt.Deliver(a.user, &abcictrl.MsgSetEndBlockState{
		EndBlock: &abci.RequestEndBlock{
			Height: block.Height,
		},
	})
	if err != nil {
		panic(err)
	}
	// TODO real endblock for modules

	// we check if there was any val updates changes done by modules
	valUpdates := new(abcictrl.ValidatorUpdates)
	err = a.rt.Get(meta.SingletonID, valUpdates)
	if err != nil {
		panic(err)
	}
	// convert from protov2 to legacy protov1
	updates := make([]types.ValidatorUpdate, len(valUpdates.ValidatorUpdates))
	for i, val := range valUpdates.ValidatorUpdates {
		updates[i] = val.ToLegacy()
	}
	// updates..
	return types.ResponseEndBlock{
		ValidatorUpdates:      updates,
		ConsensusParamUpdates: nil,
		Events:                nil,
	}
}

func (a ABCIApplication) Commit() types.ResponseCommit {
	// TODO in commit we should do clearing of old ABCI controller state :)
	return types.ResponseCommit{
		Data:         []byte("constant"),
		RetainHeight: 0,
	}
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
