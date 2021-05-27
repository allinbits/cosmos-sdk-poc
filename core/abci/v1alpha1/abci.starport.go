package v1alpha1

import (
	client "github.com/fdymylja/tmos/runtime/client"
	meta "github.com/fdymylja/tmos/runtime/meta"
	schema "github.com/fdymylja/tmos/runtime/orm/schema"
)

func (x *CurrentBlock) StateObject() {}

func (x *CurrentBlock) New() meta.StateObject {
	return new(CurrentBlock)
}

type CurrentBlockClient interface {
	Get(opts ...client.GetOption) (*CurrentBlock, error)
	Create(currentBlock *CurrentBlock, opts ...client.CreateOption) error
	Delete(currentBlock *CurrentBlock, opts ...client.DeleteOption) error
	Update(currentBlock *CurrentBlock, opts ...client.UpdateOption) error
}

type currentBlockClient struct {
	client client.RuntimeClient
}

func (x *currentBlockClient) Get(opts ...client.GetOption) (*CurrentBlock, error) {
	_spfGenO := new(CurrentBlock)
	_spfGenErr := x.client.Get(meta.SingletonID, _spfGenO, opts...)
	if _spfGenErr != nil {
		return nil, _spfGenErr
	}
	return _spfGenO, nil
}

func (x *currentBlockClient) Create(currentBlock *CurrentBlock, opts ...client.CreateOption) error {
	return x.client.Create(currentBlock, opts...)
}

func (x *currentBlockClient) Delete(currentBlock *CurrentBlock, opts ...client.DeleteOption) error {
	return x.client.Delete(currentBlock, opts...)
}

func (x *currentBlockClient) Update(currentBlock *CurrentBlock, opts ...client.UpdateOption) error {
	return x.client.Update(currentBlock, opts...)
}

func (x *Stage) StateObject() {}

func (x *Stage) New() meta.StateObject {
	return new(Stage)
}

type StageClient interface {
	Get(opts ...client.GetOption) (*Stage, error)
	Create(stage *Stage, opts ...client.CreateOption) error
	Delete(stage *Stage, opts ...client.DeleteOption) error
	Update(stage *Stage, opts ...client.UpdateOption) error
}

type stageClient struct {
	client client.RuntimeClient
}

func (x *stageClient) Get(opts ...client.GetOption) (*Stage, error) {
	_spfGenO := new(Stage)
	_spfGenErr := x.client.Get(meta.SingletonID, _spfGenO, opts...)
	if _spfGenErr != nil {
		return nil, _spfGenErr
	}
	return _spfGenO, nil
}

func (x *stageClient) Create(stage *Stage, opts ...client.CreateOption) error {
	return x.client.Create(stage, opts...)
}

func (x *stageClient) Delete(stage *Stage, opts ...client.DeleteOption) error {
	return x.client.Delete(stage, opts...)
}

func (x *stageClient) Update(stage *Stage, opts ...client.UpdateOption) error {
	return x.client.Update(stage, opts...)
}

func (x *InitChainInfo) StateObject() {}

func (x *InitChainInfo) New() meta.StateObject {
	return new(InitChainInfo)
}

type InitChainInfoClient interface {
	Get(opts ...client.GetOption) (*InitChainInfo, error)
	Create(initChainInfo *InitChainInfo, opts ...client.CreateOption) error
	Delete(initChainInfo *InitChainInfo, opts ...client.DeleteOption) error
	Update(initChainInfo *InitChainInfo, opts ...client.UpdateOption) error
}

type initChainInfoClient struct {
	client client.RuntimeClient
}

func (x *initChainInfoClient) Get(opts ...client.GetOption) (*InitChainInfo, error) {
	_spfGenO := new(InitChainInfo)
	_spfGenErr := x.client.Get(meta.SingletonID, _spfGenO, opts...)
	if _spfGenErr != nil {
		return nil, _spfGenErr
	}
	return _spfGenO, nil
}

func (x *initChainInfoClient) Create(initChainInfo *InitChainInfo, opts ...client.CreateOption) error {
	return x.client.Create(initChainInfo, opts...)
}

func (x *initChainInfoClient) Delete(initChainInfo *InitChainInfo, opts ...client.DeleteOption) error {
	return x.client.Delete(initChainInfo, opts...)
}

func (x *initChainInfoClient) Update(initChainInfo *InitChainInfo, opts ...client.UpdateOption) error {
	return x.client.Update(initChainInfo, opts...)
}

func (x *BeginBlockState) StateObject() {}

func (x *BeginBlockState) New() meta.StateObject {
	return new(BeginBlockState)
}

type BeginBlockStateClient interface {
	Get(opts ...client.GetOption) (*BeginBlockState, error)
	Create(beginBlockState *BeginBlockState, opts ...client.CreateOption) error
	Delete(beginBlockState *BeginBlockState, opts ...client.DeleteOption) error
	Update(beginBlockState *BeginBlockState, opts ...client.UpdateOption) error
}

type beginBlockStateClient struct {
	client client.RuntimeClient
}

func (x *beginBlockStateClient) Get(opts ...client.GetOption) (*BeginBlockState, error) {
	_spfGenO := new(BeginBlockState)
	_spfGenErr := x.client.Get(meta.SingletonID, _spfGenO, opts...)
	if _spfGenErr != nil {
		return nil, _spfGenErr
	}
	return _spfGenO, nil
}

func (x *beginBlockStateClient) Create(beginBlockState *BeginBlockState, opts ...client.CreateOption) error {
	return x.client.Create(beginBlockState, opts...)
}

func (x *beginBlockStateClient) Delete(beginBlockState *BeginBlockState, opts ...client.DeleteOption) error {
	return x.client.Delete(beginBlockState, opts...)
}

func (x *beginBlockStateClient) Update(beginBlockState *BeginBlockState, opts ...client.UpdateOption) error {
	return x.client.Update(beginBlockState, opts...)
}

func (x *CheckTxState) StateObject() {}

func (x *CheckTxState) New() meta.StateObject {
	return new(CheckTxState)
}

type CheckTxStateClient interface {
	Get(opts ...client.GetOption) (*CheckTxState, error)
	Create(checkTxState *CheckTxState, opts ...client.CreateOption) error
	Delete(checkTxState *CheckTxState, opts ...client.DeleteOption) error
	Update(checkTxState *CheckTxState, opts ...client.UpdateOption) error
}

type checkTxStateClient struct {
	client client.RuntimeClient
}

func (x *checkTxStateClient) Get(opts ...client.GetOption) (*CheckTxState, error) {
	_spfGenO := new(CheckTxState)
	_spfGenErr := x.client.Get(meta.SingletonID, _spfGenO, opts...)
	if _spfGenErr != nil {
		return nil, _spfGenErr
	}
	return _spfGenO, nil
}

func (x *checkTxStateClient) Create(checkTxState *CheckTxState, opts ...client.CreateOption) error {
	return x.client.Create(checkTxState, opts...)
}

func (x *checkTxStateClient) Delete(checkTxState *CheckTxState, opts ...client.DeleteOption) error {
	return x.client.Delete(checkTxState, opts...)
}

func (x *checkTxStateClient) Update(checkTxState *CheckTxState, opts ...client.UpdateOption) error {
	return x.client.Update(checkTxState, opts...)
}

func (x *DeliverTxState) StateObject() {}

func (x *DeliverTxState) New() meta.StateObject {
	return new(DeliverTxState)
}

type DeliverTxStateClient interface {
	Get(opts ...client.GetOption) (*DeliverTxState, error)
	Create(deliverTxState *DeliverTxState, opts ...client.CreateOption) error
	Delete(deliverTxState *DeliverTxState, opts ...client.DeleteOption) error
	Update(deliverTxState *DeliverTxState, opts ...client.UpdateOption) error
}

type deliverTxStateClient struct {
	client client.RuntimeClient
}

func (x *deliverTxStateClient) Get(opts ...client.GetOption) (*DeliverTxState, error) {
	_spfGenO := new(DeliverTxState)
	_spfGenErr := x.client.Get(meta.SingletonID, _spfGenO, opts...)
	if _spfGenErr != nil {
		return nil, _spfGenErr
	}
	return _spfGenO, nil
}

func (x *deliverTxStateClient) Create(deliverTxState *DeliverTxState, opts ...client.CreateOption) error {
	return x.client.Create(deliverTxState, opts...)
}

func (x *deliverTxStateClient) Delete(deliverTxState *DeliverTxState, opts ...client.DeleteOption) error {
	return x.client.Delete(deliverTxState, opts...)
}

func (x *deliverTxStateClient) Update(deliverTxState *DeliverTxState, opts ...client.UpdateOption) error {
	return x.client.Update(deliverTxState, opts...)
}

func (x *ValidatorUpdates) StateObject() {}

func (x *ValidatorUpdates) New() meta.StateObject {
	return new(ValidatorUpdates)
}

type ValidatorUpdatesClient interface {
	Get(opts ...client.GetOption) (*ValidatorUpdates, error)
	Create(validatorUpdates *ValidatorUpdates, opts ...client.CreateOption) error
	Delete(validatorUpdates *ValidatorUpdates, opts ...client.DeleteOption) error
	Update(validatorUpdates *ValidatorUpdates, opts ...client.UpdateOption) error
}

type validatorUpdatesClient struct {
	client client.RuntimeClient
}

func (x *validatorUpdatesClient) Get(opts ...client.GetOption) (*ValidatorUpdates, error) {
	_spfGenO := new(ValidatorUpdates)
	_spfGenErr := x.client.Get(meta.SingletonID, _spfGenO, opts...)
	if _spfGenErr != nil {
		return nil, _spfGenErr
	}
	return _spfGenO, nil
}

func (x *validatorUpdatesClient) Create(validatorUpdates *ValidatorUpdates, opts ...client.CreateOption) error {
	return x.client.Create(validatorUpdates, opts...)
}

func (x *validatorUpdatesClient) Delete(validatorUpdates *ValidatorUpdates, opts ...client.DeleteOption) error {
	return x.client.Delete(validatorUpdates, opts...)
}

func (x *validatorUpdatesClient) Update(validatorUpdates *ValidatorUpdates, opts ...client.UpdateOption) error {
	return x.client.Update(validatorUpdates, opts...)
}

func (x *EndBlockState) StateObject() {}

func (x *EndBlockState) New() meta.StateObject {
	return new(EndBlockState)
}

type EndBlockStateClient interface {
	Get(opts ...client.GetOption) (*EndBlockState, error)
	Create(endBlockState *EndBlockState, opts ...client.CreateOption) error
	Delete(endBlockState *EndBlockState, opts ...client.DeleteOption) error
	Update(endBlockState *EndBlockState, opts ...client.UpdateOption) error
}

type endBlockStateClient struct {
	client client.RuntimeClient
}

func (x *endBlockStateClient) Get(opts ...client.GetOption) (*EndBlockState, error) {
	_spfGenO := new(EndBlockState)
	_spfGenErr := x.client.Get(meta.SingletonID, _spfGenO, opts...)
	if _spfGenErr != nil {
		return nil, _spfGenErr
	}
	return _spfGenO, nil
}

func (x *endBlockStateClient) Create(endBlockState *EndBlockState, opts ...client.CreateOption) error {
	return x.client.Create(endBlockState, opts...)
}

func (x *endBlockStateClient) Delete(endBlockState *EndBlockState, opts ...client.DeleteOption) error {
	return x.client.Delete(endBlockState, opts...)
}

func (x *endBlockStateClient) Update(endBlockState *EndBlockState, opts ...client.UpdateOption) error {
	return x.client.Update(endBlockState, opts...)
}

func (x *MsgSetBeginBlockState) StateTransition() {}

func (x *MsgSetBeginBlockState) New() meta.StateTransition {
	return new(MsgSetBeginBlockState)
}

func (x *MsgSetCheckTxState) StateTransition() {}

func (x *MsgSetCheckTxState) New() meta.StateTransition {
	return new(MsgSetCheckTxState)
}

func (x *MsgSetDeliverTxState) StateTransition() {}

func (x *MsgSetDeliverTxState) New() meta.StateTransition {
	return new(MsgSetDeliverTxState)
}

func (x *MsgSetEndBlockState) StateTransition() {}

func (x *MsgSetEndBlockState) New() meta.StateTransition {
	return new(MsgSetEndBlockState)
}

func (x *MsgSetInitChain) StateTransition() {}

func (x *MsgSetInitChain) New() meta.StateTransition {
	return new(MsgSetInitChain)
}

func (x *MsgSetValidatorUpdates) StateTransition() {}

func (x *MsgSetValidatorUpdates) New() meta.StateTransition {
	return new(MsgSetValidatorUpdates)
}

var CurrentBlockSchema = schema.Definition{
	Meta: meta.Meta{
		APIGroup: "tmos.abci.v1alpha1",
		APIKind:  "CurrentBlock",
	},
	Singleton: true,
}

var StageSchema = schema.Definition{
	Meta: meta.Meta{
		APIGroup: "tmos.abci.v1alpha1",
		APIKind:  "Stage",
	},
	Singleton: true,
}

var InitChainInfoSchema = schema.Definition{
	Meta: meta.Meta{
		APIGroup: "tmos.abci.v1alpha1",
		APIKind:  "InitChainInfo",
	},
	Singleton: true,
}

var BeginBlockStateSchema = schema.Definition{
	Meta: meta.Meta{
		APIGroup: "tmos.abci.v1alpha1",
		APIKind:  "BeginBlockState",
	},
	Singleton: true,
}

var CheckTxStateSchema = schema.Definition{
	Meta: meta.Meta{
		APIGroup: "tmos.abci.v1alpha1",
		APIKind:  "CheckTxState",
	},
	Singleton: true,
}

var DeliverTxStateSchema = schema.Definition{
	Meta: meta.Meta{
		APIGroup: "tmos.abci.v1alpha1",
		APIKind:  "DeliverTxState",
	},
	Singleton: true,
}

var ValidatorUpdatesSchema = schema.Definition{
	Meta: meta.Meta{
		APIGroup: "tmos.abci.v1alpha1",
		APIKind:  "ValidatorUpdates",
	},
	Singleton: true,
}

var EndBlockStateSchema = schema.Definition{
	Meta: meta.Meta{
		APIGroup: "tmos.abci.v1alpha1",
		APIKind:  "EndBlockState",
	},
	Singleton: true,
}

type ClientSet interface {
	CurrentBlocks() CurrentBlockClient
	Stages() StageClient
	InitChainInfos() InitChainInfoClient
	BeginBlockStates() BeginBlockStateClient
	CheckTxStates() CheckTxStateClient
	DeliverTxStates() DeliverTxStateClient
	ValidatorUpdates() ValidatorUpdatesClient
	EndBlockStates() EndBlockStateClient
	ExecMsgSetBeginBlockState(msg *MsgSetBeginBlockState) error
	ExecMsgSetCheckTxState(msg *MsgSetCheckTxState) error
	ExecMsgSetDeliverTxState(msg *MsgSetDeliverTxState) error
	ExecMsgSetEndBlockState(msg *MsgSetEndBlockState) error
	ExecMsgSetInitChain(msg *MsgSetInitChain) error
	ExecMsgSetValidatorUpdates(msg *MsgSetValidatorUpdates) error
}

func NewClientSet(client client.RuntimeClient) ClientSet {
	return &clientSet{
		client:                 client,
		currentBlockClient:     &currentBlockClient{client: client},
		stageClient:            &stageClient{client: client},
		initChainInfoClient:    &initChainInfoClient{client: client},
		beginBlockStateClient:  &beginBlockStateClient{client: client},
		checkTxStateClient:     &checkTxStateClient{client: client},
		deliverTxStateClient:   &deliverTxStateClient{client: client},
		validatorUpdatesClient: &validatorUpdatesClient{client: client},
		endBlockStateClient:    &endBlockStateClient{client: client},
	}
}

type clientSet struct {
	client client.RuntimeClient
	// currentBlockClient is the client used to interact with CurrentBlock
	currentBlockClient CurrentBlockClient
	// stageClient is the client used to interact with Stage
	stageClient StageClient
	// initChainInfoClient is the client used to interact with InitChainInfo
	initChainInfoClient InitChainInfoClient
	// beginBlockStateClient is the client used to interact with BeginBlockState
	beginBlockStateClient BeginBlockStateClient
	// checkTxStateClient is the client used to interact with CheckTxState
	checkTxStateClient CheckTxStateClient
	// deliverTxStateClient is the client used to interact with DeliverTxState
	deliverTxStateClient DeliverTxStateClient
	// validatorUpdatesClient is the client used to interact with ValidatorUpdates
	validatorUpdatesClient ValidatorUpdatesClient
	// endBlockStateClient is the client used to interact with EndBlockState
	endBlockStateClient EndBlockStateClient
}

func (x *clientSet) CurrentBlocks() CurrentBlockClient {
	return x.currentBlockClient
}

func (x *clientSet) Stages() StageClient {
	return x.stageClient
}

func (x *clientSet) InitChainInfos() InitChainInfoClient {
	return x.initChainInfoClient
}

func (x *clientSet) BeginBlockStates() BeginBlockStateClient {
	return x.beginBlockStateClient
}

func (x *clientSet) CheckTxStates() CheckTxStateClient {
	return x.checkTxStateClient
}

func (x *clientSet) DeliverTxStates() DeliverTxStateClient {
	return x.deliverTxStateClient
}

func (x *clientSet) ValidatorUpdates() ValidatorUpdatesClient {
	return x.validatorUpdatesClient
}

func (x *clientSet) EndBlockStates() EndBlockStateClient {
	return x.endBlockStateClient
}

func (x *clientSet) ExecMsgSetBeginBlockState(msg *MsgSetBeginBlockState) error {
	return x.client.Deliver(msg)
}

func (x *clientSet) ExecMsgSetCheckTxState(msg *MsgSetCheckTxState) error {
	return x.client.Deliver(msg)
}

func (x *clientSet) ExecMsgSetDeliverTxState(msg *MsgSetDeliverTxState) error {
	return x.client.Deliver(msg)
}

func (x *clientSet) ExecMsgSetEndBlockState(msg *MsgSetEndBlockState) error {
	return x.client.Deliver(msg)
}

func (x *clientSet) ExecMsgSetInitChain(msg *MsgSetInitChain) error {
	return x.client.Deliver(msg)
}

func (x *clientSet) ExecMsgSetValidatorUpdates(msg *MsgSetValidatorUpdates) error {
	return x.client.Deliver(msg)
}
