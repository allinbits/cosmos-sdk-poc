package runtime

import (
	"strings"

	"github.com/fdymylja/tmos/module/abci/tendermint/abci"
	abcictrl "github.com/fdymylja/tmos/module/abci/v1alpha1"
	runtimev1alpha1 "github.com/fdymylja/tmos/module/runtime/v1alpha1"
	"github.com/fdymylja/tmos/runtime/meta"
	"github.com/tendermint/tendermint/abci/types"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
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
	// TODO get store based on height
	splitted := strings.Split(query.Path, "/")
	switch len(splitted) {
	case 3:
	case 4:
		splitted = splitted[1:]
	default:
		return types.ResponseQuery{Code: CodeBadRequest, Log: "invalid path"}
	}
	verb := splitted[0]
	stateObjectName := splitted[1] // TODO check if store knows of the existence of this stateObjectName
	key := splitted[2]
	switch verb {
	case runtimev1alpha1.Verb_Get.String():
		objType, err := protoregistry.GlobalTypes.FindMessageByName(protoreflect.FullName(stateObjectName))
		if err != nil {
			return types.ResponseQuery{Code: CodeNotFound}
		}
		object := objType.New().Interface().(meta.StateObject)
		err = a.rt.Get(meta.NewStringID(key), object)
		if err != nil {
			return types.ResponseQuery{Code: CodeUnknown, Log: err.Error()}
		}
		jsonObject, err := protojson.Marshal(object)
		if err != nil {
			return types.ResponseQuery{Code: CodeUnknown, Log: err.Error()}
		}
		return types.ResponseQuery{
			Value: jsonObject,
		}
	default:
		return types.ResponseQuery{Code: CodeBadRequest, Log: "unsupported verb" + verb}
	}
}

func (a ABCIApplication) CheckTx(tmTx types.RequestCheckTx) types.ResponseCheckTx {
	// decode tx
	tx, err := a.rt.authn.DecodeTx(tmTx.Tx)
	if err != nil {
		return types.ResponseCheckTx(ToABCIResponse(0, 0, err))
	}
	// run admission checks on the transaction
	err = a.rt.runTxAdmissionChain(tx)
	if err != nil {
		return types.ResponseCheckTx(ToABCIResponse(0, 0, err))
	}
	// run admission checks on single state transitions
	for _, transition := range tx.StateTransitions() {
		err = a.rt.runAdmissionChain(transition)
		if err != nil {
			return types.ResponseCheckTx(ToABCIResponse(0, 0, err))
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
	// do authentication
	err = a.rt.authn.Authenticate(tx)
	if err != nil {
		return ToABCIResponse(0, 0, err)
	}
	// TODO cache the store
	// todo run authentication chain
	err = a.rt.runTxPostAuthenticationChain(tx)
	if err != nil {
		return ToABCIResponse(0, 0, err)
	}
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
	return types.ResponseDeliverTx{
		Code:      0,
		Data:      nil,
		Log:       "",
		Info:      "",
		GasWanted: 0,
		GasUsed:   0,
		Events:    nil,
		Codespace: "",
	}
}

func (a ABCIApplication) EndBlock(block types.RequestEndBlock) types.ResponseEndBlock {
	// we set the abci endblcok state given us by tendermint so other modules can access it
	err := a.rt.Deliver(nil, &abcictrl.MsgSetEndBlockState{
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
	err = a.rt.Get(abcictrl.ValidatorUpdatesID, valUpdates)
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
