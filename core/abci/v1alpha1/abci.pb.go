// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.13.0
// source: core/abci/v1alpha1/abci.proto

package v1alpha1

import (
	abci "github.com/fdymylja/tmos/core/abci/tendermint/abci"
	_ "github.com/fdymylja/tmos/core/modulegen"
	_ "github.com/fdymylja/tmos/core/rbac/xt"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ABCIStage int32

const (
	ABCIStage_Unknown    ABCIStage = 0
	ABCIStage_InitChain  ABCIStage = 1
	ABCIStage_BeginBlock ABCIStage = 2
	ABCIStage_CheckTx    ABCIStage = 3
	ABCIStage_ReCheckTx  ABCIStage = 4
	ABCIStage_DeliverTx  ABCIStage = 5
	ABCIStage_EndBlock   ABCIStage = 6
	ABCIStage_Commit     ABCIStage = 7
)

// Enum value maps for ABCIStage.
var (
	ABCIStage_name = map[int32]string{
		0: "Unknown",
		1: "InitChain",
		2: "BeginBlock",
		3: "CheckTx",
		4: "ReCheckTx",
		5: "DeliverTx",
		6: "EndBlock",
		7: "Commit",
	}
	ABCIStage_value = map[string]int32{
		"Unknown":    0,
		"InitChain":  1,
		"BeginBlock": 2,
		"CheckTx":    3,
		"ReCheckTx":  4,
		"DeliverTx":  5,
		"EndBlock":   6,
		"Commit":     7,
	}
)

func (x ABCIStage) Enum() *ABCIStage {
	p := new(ABCIStage)
	*p = x
	return p
}

func (x ABCIStage) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ABCIStage) Descriptor() protoreflect.EnumDescriptor {
	return file_core_abci_v1alpha1_abci_proto_enumTypes[0].Descriptor()
}

func (ABCIStage) Type() protoreflect.EnumType {
	return &file_core_abci_v1alpha1_abci_proto_enumTypes[0]
}

func (x ABCIStage) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ABCIStage.Descriptor instead.
func (ABCIStage) EnumDescriptor() ([]byte, []int) {
	return file_core_abci_v1alpha1_abci_proto_rawDescGZIP(), []int{0}
}

type CurrentBlock struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockNumber uint64 `protobuf:"varint,1,opt,name=block_number,json=blockNumber,proto3" json:"block_number,omitempty"`
}

func (x *CurrentBlock) Reset() {
	*x = CurrentBlock{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_abci_v1alpha1_abci_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CurrentBlock) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CurrentBlock) ProtoMessage() {}

func (x *CurrentBlock) ProtoReflect() protoreflect.Message {
	mi := &file_core_abci_v1alpha1_abci_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CurrentBlock.ProtoReflect.Descriptor instead.
func (*CurrentBlock) Descriptor() ([]byte, []int) {
	return file_core_abci_v1alpha1_abci_proto_rawDescGZIP(), []int{0}
}

func (x *CurrentBlock) GetBlockNumber() uint64 {
	if x != nil {
		return x.BlockNumber
	}
	return 0
}

type Stage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stage ABCIStage `protobuf:"varint,1,opt,name=stage,proto3,enum=tmos.abci.v1alpha1.ABCIStage" json:"stage,omitempty"`
}

func (x *Stage) Reset() {
	*x = Stage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_abci_v1alpha1_abci_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Stage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Stage) ProtoMessage() {}

func (x *Stage) ProtoReflect() protoreflect.Message {
	mi := &file_core_abci_v1alpha1_abci_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Stage.ProtoReflect.Descriptor instead.
func (*Stage) Descriptor() ([]byte, []int) {
	return file_core_abci_v1alpha1_abci_proto_rawDescGZIP(), []int{1}
}

func (x *Stage) GetStage() ABCIStage {
	if x != nil {
		return x.Stage
	}
	return ABCIStage_Unknown
}

type InitChainInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChainId string `protobuf:"bytes,1,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
}

func (x *InitChainInfo) Reset() {
	*x = InitChainInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_abci_v1alpha1_abci_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitChainInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitChainInfo) ProtoMessage() {}

func (x *InitChainInfo) ProtoReflect() protoreflect.Message {
	mi := &file_core_abci_v1alpha1_abci_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitChainInfo.ProtoReflect.Descriptor instead.
func (*InitChainInfo) Descriptor() ([]byte, []int) {
	return file_core_abci_v1alpha1_abci_proto_rawDescGZIP(), []int{2}
}

func (x *InitChainInfo) GetChainId() string {
	if x != nil {
		return x.ChainId
	}
	return ""
}

type BeginBlockState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BeginBlock *abci.RequestBeginBlock `protobuf:"bytes,1,opt,name=begin_block,json=beginBlock,proto3" json:"begin_block,omitempty"`
}

func (x *BeginBlockState) Reset() {
	*x = BeginBlockState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_abci_v1alpha1_abci_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BeginBlockState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BeginBlockState) ProtoMessage() {}

func (x *BeginBlockState) ProtoReflect() protoreflect.Message {
	mi := &file_core_abci_v1alpha1_abci_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BeginBlockState.ProtoReflect.Descriptor instead.
func (*BeginBlockState) Descriptor() ([]byte, []int) {
	return file_core_abci_v1alpha1_abci_proto_rawDescGZIP(), []int{3}
}

func (x *BeginBlockState) GetBeginBlock() *abci.RequestBeginBlock {
	if x != nil {
		return x.BeginBlock
	}
	return nil
}

type CheckTxState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CheckTx *abci.RequestCheckTx `protobuf:"bytes,1,opt,name=check_tx,json=checkTx,proto3" json:"check_tx,omitempty"`
}

func (x *CheckTxState) Reset() {
	*x = CheckTxState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_abci_v1alpha1_abci_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckTxState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckTxState) ProtoMessage() {}

func (x *CheckTxState) ProtoReflect() protoreflect.Message {
	mi := &file_core_abci_v1alpha1_abci_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckTxState.ProtoReflect.Descriptor instead.
func (*CheckTxState) Descriptor() ([]byte, []int) {
	return file_core_abci_v1alpha1_abci_proto_rawDescGZIP(), []int{4}
}

func (x *CheckTxState) GetCheckTx() *abci.RequestCheckTx {
	if x != nil {
		return x.CheckTx
	}
	return nil
}

type DeliverTxState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeliverTx *abci.RequestDeliverTx `protobuf:"bytes,1,opt,name=deliver_tx,json=deliverTx,proto3" json:"deliver_tx,omitempty"`
}

func (x *DeliverTxState) Reset() {
	*x = DeliverTxState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_abci_v1alpha1_abci_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeliverTxState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeliverTxState) ProtoMessage() {}

func (x *DeliverTxState) ProtoReflect() protoreflect.Message {
	mi := &file_core_abci_v1alpha1_abci_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeliverTxState.ProtoReflect.Descriptor instead.
func (*DeliverTxState) Descriptor() ([]byte, []int) {
	return file_core_abci_v1alpha1_abci_proto_rawDescGZIP(), []int{5}
}

func (x *DeliverTxState) GetDeliverTx() *abci.RequestDeliverTx {
	if x != nil {
		return x.DeliverTx
	}
	return nil
}

type ValidatorUpdates struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ValidatorUpdates []*abci.ValidatorUpdate `protobuf:"bytes,1,rep,name=validator_updates,json=validatorUpdates,proto3" json:"validator_updates,omitempty"`
}

func (x *ValidatorUpdates) Reset() {
	*x = ValidatorUpdates{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_abci_v1alpha1_abci_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidatorUpdates) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidatorUpdates) ProtoMessage() {}

func (x *ValidatorUpdates) ProtoReflect() protoreflect.Message {
	mi := &file_core_abci_v1alpha1_abci_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidatorUpdates.ProtoReflect.Descriptor instead.
func (*ValidatorUpdates) Descriptor() ([]byte, []int) {
	return file_core_abci_v1alpha1_abci_proto_rawDescGZIP(), []int{6}
}

func (x *ValidatorUpdates) GetValidatorUpdates() []*abci.ValidatorUpdate {
	if x != nil {
		return x.ValidatorUpdates
	}
	return nil
}

type EndBlockState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EndBlock *abci.RequestEndBlock `protobuf:"bytes,1,opt,name=end_block,json=endBlock,proto3" json:"end_block,omitempty"`
}

func (x *EndBlockState) Reset() {
	*x = EndBlockState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_abci_v1alpha1_abci_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EndBlockState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EndBlockState) ProtoMessage() {}

func (x *EndBlockState) ProtoReflect() protoreflect.Message {
	mi := &file_core_abci_v1alpha1_abci_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EndBlockState.ProtoReflect.Descriptor instead.
func (*EndBlockState) Descriptor() ([]byte, []int) {
	return file_core_abci_v1alpha1_abci_proto_rawDescGZIP(), []int{7}
}

func (x *EndBlockState) GetEndBlock() *abci.RequestEndBlock {
	if x != nil {
		return x.EndBlock
	}
	return nil
}

type MsgSetBeginBlockState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BeginBlock *abci.RequestBeginBlock `protobuf:"bytes,1,opt,name=begin_block,json=beginBlock,proto3" json:"begin_block,omitempty"`
}

func (x *MsgSetBeginBlockState) Reset() {
	*x = MsgSetBeginBlockState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_abci_v1alpha1_abci_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgSetBeginBlockState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgSetBeginBlockState) ProtoMessage() {}

func (x *MsgSetBeginBlockState) ProtoReflect() protoreflect.Message {
	mi := &file_core_abci_v1alpha1_abci_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgSetBeginBlockState.ProtoReflect.Descriptor instead.
func (*MsgSetBeginBlockState) Descriptor() ([]byte, []int) {
	return file_core_abci_v1alpha1_abci_proto_rawDescGZIP(), []int{8}
}

func (x *MsgSetBeginBlockState) GetBeginBlock() *abci.RequestBeginBlock {
	if x != nil {
		return x.BeginBlock
	}
	return nil
}

type MsgSetCheckTxState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CheckTx *abci.RequestCheckTx `protobuf:"bytes,1,opt,name=check_tx,json=checkTx,proto3" json:"check_tx,omitempty"`
}

func (x *MsgSetCheckTxState) Reset() {
	*x = MsgSetCheckTxState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_abci_v1alpha1_abci_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgSetCheckTxState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgSetCheckTxState) ProtoMessage() {}

func (x *MsgSetCheckTxState) ProtoReflect() protoreflect.Message {
	mi := &file_core_abci_v1alpha1_abci_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgSetCheckTxState.ProtoReflect.Descriptor instead.
func (*MsgSetCheckTxState) Descriptor() ([]byte, []int) {
	return file_core_abci_v1alpha1_abci_proto_rawDescGZIP(), []int{9}
}

func (x *MsgSetCheckTxState) GetCheckTx() *abci.RequestCheckTx {
	if x != nil {
		return x.CheckTx
	}
	return nil
}

type MsgSetDeliverTxState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeliverTx *abci.RequestDeliverTx `protobuf:"bytes,1,opt,name=deliver_tx,json=deliverTx,proto3" json:"deliver_tx,omitempty"`
}

func (x *MsgSetDeliverTxState) Reset() {
	*x = MsgSetDeliverTxState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_abci_v1alpha1_abci_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgSetDeliverTxState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgSetDeliverTxState) ProtoMessage() {}

func (x *MsgSetDeliverTxState) ProtoReflect() protoreflect.Message {
	mi := &file_core_abci_v1alpha1_abci_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgSetDeliverTxState.ProtoReflect.Descriptor instead.
func (*MsgSetDeliverTxState) Descriptor() ([]byte, []int) {
	return file_core_abci_v1alpha1_abci_proto_rawDescGZIP(), []int{10}
}

func (x *MsgSetDeliverTxState) GetDeliverTx() *abci.RequestDeliverTx {
	if x != nil {
		return x.DeliverTx
	}
	return nil
}

type MsgSetEndBlockState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EndBlock *abci.RequestEndBlock `protobuf:"bytes,1,opt,name=end_block,json=endBlock,proto3" json:"end_block,omitempty"`
}

func (x *MsgSetEndBlockState) Reset() {
	*x = MsgSetEndBlockState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_abci_v1alpha1_abci_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgSetEndBlockState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgSetEndBlockState) ProtoMessage() {}

func (x *MsgSetEndBlockState) ProtoReflect() protoreflect.Message {
	mi := &file_core_abci_v1alpha1_abci_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgSetEndBlockState.ProtoReflect.Descriptor instead.
func (*MsgSetEndBlockState) Descriptor() ([]byte, []int) {
	return file_core_abci_v1alpha1_abci_proto_rawDescGZIP(), []int{11}
}

func (x *MsgSetEndBlockState) GetEndBlock() *abci.RequestEndBlock {
	if x != nil {
		return x.EndBlock
	}
	return nil
}

type MsgSetInitChain struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InitChainInfo *InitChainInfo `protobuf:"bytes,1,opt,name=init_chain_info,json=initChainInfo,proto3" json:"init_chain_info,omitempty"`
}

func (x *MsgSetInitChain) Reset() {
	*x = MsgSetInitChain{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_abci_v1alpha1_abci_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgSetInitChain) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgSetInitChain) ProtoMessage() {}

func (x *MsgSetInitChain) ProtoReflect() protoreflect.Message {
	mi := &file_core_abci_v1alpha1_abci_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgSetInitChain.ProtoReflect.Descriptor instead.
func (*MsgSetInitChain) Descriptor() ([]byte, []int) {
	return file_core_abci_v1alpha1_abci_proto_rawDescGZIP(), []int{12}
}

func (x *MsgSetInitChain) GetInitChainInfo() *InitChainInfo {
	if x != nil {
		return x.InitChainInfo
	}
	return nil
}

type MsgSetValidatorUpdates struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ValidatorUpdates []*abci.ValidatorUpdate `protobuf:"bytes,1,rep,name=validator_updates,json=validatorUpdates,proto3" json:"validator_updates,omitempty"`
}

func (x *MsgSetValidatorUpdates) Reset() {
	*x = MsgSetValidatorUpdates{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_abci_v1alpha1_abci_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgSetValidatorUpdates) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgSetValidatorUpdates) ProtoMessage() {}

func (x *MsgSetValidatorUpdates) ProtoReflect() protoreflect.Message {
	mi := &file_core_abci_v1alpha1_abci_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgSetValidatorUpdates.ProtoReflect.Descriptor instead.
func (*MsgSetValidatorUpdates) Descriptor() ([]byte, []int) {
	return file_core_abci_v1alpha1_abci_proto_rawDescGZIP(), []int{13}
}

func (x *MsgSetValidatorUpdates) GetValidatorUpdates() []*abci.ValidatorUpdate {
	if x != nil {
		return x.ValidatorUpdates
	}
	return nil
}

var File_core_abci_v1alpha1_abci_proto protoreflect.FileDescriptor

var file_core_abci_v1alpha1_abci_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x61, 0x62, 0x63, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2f, 0x61, 0x62, 0x63, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x12, 0x74, 0x6d, 0x6f, 0x73, 0x2e, 0x61, 0x62, 0x63, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x1a, 0x25, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x61, 0x62, 0x63, 0x69, 0x2f, 0x74,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2f, 0x61, 0x62, 0x63, 0x69, 0x2f, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x63, 0x6f, 0x72, 0x65,
	0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x67, 0x65, 0x6e, 0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c,
	0x65, 0x67, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x63, 0x6f, 0x72, 0x65,
	0x2f, 0x72, 0x62, 0x61, 0x63, 0x2f, 0x78, 0x74, 0x2f, 0x72, 0x62, 0x61, 0x63, 0x5f, 0x65, 0x78,
	0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3a, 0x0a,
	0x0c, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x21, 0x0a,
	0x0c, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x3a, 0x07, 0x92, 0xff, 0x9d, 0x04, 0x02, 0x08, 0x01, 0x22, 0x45, 0x0a, 0x05, 0x53, 0x74, 0x61,
	0x67, 0x65, 0x12, 0x33, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x1d, 0x2e, 0x74, 0x6d, 0x6f, 0x73, 0x2e, 0x61, 0x62, 0x63, 0x69, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x41, 0x42, 0x43, 0x49, 0x53, 0x74, 0x61, 0x67, 0x65,
	0x52, 0x05, 0x73, 0x74, 0x61, 0x67, 0x65, 0x3a, 0x07, 0x92, 0xff, 0x9d, 0x04, 0x02, 0x08, 0x01,
	0x22, 0x33, 0x0a, 0x0d, 0x49, 0x6e, 0x69, 0x74, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x3a, 0x07, 0x92, 0xff,
	0x9d, 0x04, 0x02, 0x08, 0x01, 0x22, 0x5f, 0x0a, 0x0f, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x43, 0x0a, 0x0b, 0x62, 0x65, 0x67, 0x69,
	0x6e, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e,
	0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2e, 0x61, 0x62, 0x63, 0x69, 0x2e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x52, 0x0a, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x3a, 0x07, 0x92,
	0xff, 0x9d, 0x04, 0x02, 0x08, 0x01, 0x22, 0x53, 0x0a, 0x0c, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x54,
	0x78, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x3a, 0x0a, 0x08, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f,
	0x74, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x74, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2e, 0x61, 0x62, 0x63, 0x69, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x54, 0x78, 0x52, 0x07, 0x63, 0x68, 0x65, 0x63, 0x6b,
	0x54, 0x78, 0x3a, 0x07, 0x92, 0xff, 0x9d, 0x04, 0x02, 0x08, 0x01, 0x22, 0x5b, 0x0a, 0x0e, 0x44,
	0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x54, 0x78, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x40, 0x0a,
	0x0a, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x5f, 0x74, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x21, 0x2e, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2e, 0x61,
	0x62, 0x63, 0x69, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x65, 0x6c, 0x69, 0x76,
	0x65, 0x72, 0x54, 0x78, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x54, 0x78, 0x3a,
	0x07, 0x92, 0xff, 0x9d, 0x04, 0x02, 0x08, 0x01, 0x22, 0x6a, 0x0a, 0x10, 0x56, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x6f, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x12, 0x4d, 0x0a, 0x11,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x6d, 0x69, 0x6e, 0x74, 0x2e, 0x61, 0x62, 0x63, 0x69, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x6f, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x10, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x6f, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x3a, 0x07, 0x92, 0xff, 0x9d,
	0x04, 0x02, 0x08, 0x01, 0x22, 0x57, 0x0a, 0x0d, 0x45, 0x6e, 0x64, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x3d, 0x0a, 0x09, 0x65, 0x6e, 0x64, 0x5f, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x74, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2e, 0x61, 0x62, 0x63, 0x69, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x45, 0x6e, 0x64, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x3a, 0x07, 0x92, 0xff, 0x9d, 0x04, 0x02, 0x08, 0x01, 0x22, 0x68, 0x0a,
	0x15, 0x4d, 0x73, 0x67, 0x53, 0x65, 0x74, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x43, 0x0a, 0x0b, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x5f,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x74, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2e, 0x61, 0x62, 0x63, 0x69, 0x2e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52,
	0x0a, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x3a, 0x0a, 0x8a, 0xff, 0x9d,
	0x04, 0x00, 0x82, 0xb0, 0xe3, 0x2d, 0x00, 0x22, 0x5c, 0x0a, 0x12, 0x4d, 0x73, 0x67, 0x53, 0x65,
	0x74, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x54, 0x78, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x3a, 0x0a,
	0x08, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f, 0x74, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1f, 0x2e, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2e, 0x61, 0x62, 0x63,
	0x69, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x54, 0x78,
	0x52, 0x07, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x54, 0x78, 0x3a, 0x0a, 0x8a, 0xff, 0x9d, 0x04, 0x00,
	0x82, 0xb0, 0xe3, 0x2d, 0x00, 0x22, 0x64, 0x0a, 0x14, 0x4d, 0x73, 0x67, 0x53, 0x65, 0x74, 0x44,
	0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x54, 0x78, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x40, 0x0a,
	0x0a, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x5f, 0x74, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x21, 0x2e, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2e, 0x61,
	0x62, 0x63, 0x69, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x65, 0x6c, 0x69, 0x76,
	0x65, 0x72, 0x54, 0x78, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x54, 0x78, 0x3a,
	0x0a, 0x8a, 0xff, 0x9d, 0x04, 0x00, 0x82, 0xb0, 0xe3, 0x2d, 0x00, 0x22, 0x60, 0x0a, 0x13, 0x4d,
	0x73, 0x67, 0x53, 0x65, 0x74, 0x45, 0x6e, 0x64, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x12, 0x3d, 0x0a, 0x09, 0x65, 0x6e, 0x64, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69,
	0x6e, 0x74, 0x2e, 0x61, 0x62, 0x63, 0x69, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x45,
	0x6e, 0x64, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x3a, 0x0a, 0x8a, 0xff, 0x9d, 0x04, 0x00, 0x82, 0xb0, 0xe3, 0x2d, 0x00, 0x22, 0x68, 0x0a,
	0x0f, 0x4d, 0x73, 0x67, 0x53, 0x65, 0x74, 0x49, 0x6e, 0x69, 0x74, 0x43, 0x68, 0x61, 0x69, 0x6e,
	0x12, 0x49, 0x0a, 0x0f, 0x69, 0x6e, 0x69, 0x74, 0x5f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x69,
	0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x74, 0x6d, 0x6f, 0x73,
	0x2e, 0x61, 0x62, 0x63, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x49,
	0x6e, 0x69, 0x74, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0d, 0x69, 0x6e,
	0x69, 0x74, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x3a, 0x0a, 0x8a, 0xff, 0x9d,
	0x04, 0x00, 0x82, 0xb0, 0xe3, 0x2d, 0x00, 0x22, 0x73, 0x0a, 0x16, 0x4d, 0x73, 0x67, 0x53, 0x65,
	0x74, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x73, 0x12, 0x4d, 0x0a, 0x11, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x74,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2e, 0x61, 0x62, 0x63, 0x69, 0x2e, 0x56,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x10,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73,
	0x3a, 0x0a, 0x8a, 0xff, 0x9d, 0x04, 0x00, 0x82, 0xb0, 0xe3, 0x2d, 0x00, 0x2a, 0x7c, 0x0a, 0x09,
	0x41, 0x42, 0x43, 0x49, 0x53, 0x74, 0x61, 0x67, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x6e, 0x6b,
	0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x49, 0x6e, 0x69, 0x74, 0x43, 0x68,
	0x61, 0x69, 0x6e, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x54, 0x78,
	0x10, 0x03, 0x12, 0x0d, 0x0a, 0x09, 0x52, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x54, 0x78, 0x10,
	0x04, 0x12, 0x0d, 0x0a, 0x09, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x54, 0x78, 0x10, 0x05,
	0x12, 0x0c, 0x0a, 0x08, 0x45, 0x6e, 0x64, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x10, 0x06, 0x12, 0x0a,
	0x0a, 0x06, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x10, 0x07, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x64, 0x79, 0x6d, 0x79, 0x6c, 0x6a,
	0x61, 0x2f, 0x74, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x61, 0x62, 0x63, 0x69,
	0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_core_abci_v1alpha1_abci_proto_rawDescOnce sync.Once
	file_core_abci_v1alpha1_abci_proto_rawDescData = file_core_abci_v1alpha1_abci_proto_rawDesc
)

func file_core_abci_v1alpha1_abci_proto_rawDescGZIP() []byte {
	file_core_abci_v1alpha1_abci_proto_rawDescOnce.Do(func() {
		file_core_abci_v1alpha1_abci_proto_rawDescData = protoimpl.X.CompressGZIP(file_core_abci_v1alpha1_abci_proto_rawDescData)
	})
	return file_core_abci_v1alpha1_abci_proto_rawDescData
}

var file_core_abci_v1alpha1_abci_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_core_abci_v1alpha1_abci_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_core_abci_v1alpha1_abci_proto_goTypes = []interface{}{
	(ABCIStage)(0),                 // 0: tmos.abci.v1alpha1.ABCIStage
	(*CurrentBlock)(nil),           // 1: tmos.abci.v1alpha1.CurrentBlock
	(*Stage)(nil),                  // 2: tmos.abci.v1alpha1.Stage
	(*InitChainInfo)(nil),          // 3: tmos.abci.v1alpha1.InitChainInfo
	(*BeginBlockState)(nil),        // 4: tmos.abci.v1alpha1.BeginBlockState
	(*CheckTxState)(nil),           // 5: tmos.abci.v1alpha1.CheckTxState
	(*DeliverTxState)(nil),         // 6: tmos.abci.v1alpha1.DeliverTxState
	(*ValidatorUpdates)(nil),       // 7: tmos.abci.v1alpha1.ValidatorUpdates
	(*EndBlockState)(nil),          // 8: tmos.abci.v1alpha1.EndBlockState
	(*MsgSetBeginBlockState)(nil),  // 9: tmos.abci.v1alpha1.MsgSetBeginBlockState
	(*MsgSetCheckTxState)(nil),     // 10: tmos.abci.v1alpha1.MsgSetCheckTxState
	(*MsgSetDeliverTxState)(nil),   // 11: tmos.abci.v1alpha1.MsgSetDeliverTxState
	(*MsgSetEndBlockState)(nil),    // 12: tmos.abci.v1alpha1.MsgSetEndBlockState
	(*MsgSetInitChain)(nil),        // 13: tmos.abci.v1alpha1.MsgSetInitChain
	(*MsgSetValidatorUpdates)(nil), // 14: tmos.abci.v1alpha1.MsgSetValidatorUpdates
	(*abci.RequestBeginBlock)(nil), // 15: tendermint.abci.RequestBeginBlock
	(*abci.RequestCheckTx)(nil),    // 16: tendermint.abci.RequestCheckTx
	(*abci.RequestDeliverTx)(nil),  // 17: tendermint.abci.RequestDeliverTx
	(*abci.ValidatorUpdate)(nil),   // 18: tendermint.abci.ValidatorUpdate
	(*abci.RequestEndBlock)(nil),   // 19: tendermint.abci.RequestEndBlock
}
var file_core_abci_v1alpha1_abci_proto_depIdxs = []int32{
	0,  // 0: tmos.abci.v1alpha1.Stage.stage:type_name -> tmos.abci.v1alpha1.ABCIStage
	15, // 1: tmos.abci.v1alpha1.BeginBlockState.begin_block:type_name -> tendermint.abci.RequestBeginBlock
	16, // 2: tmos.abci.v1alpha1.CheckTxState.check_tx:type_name -> tendermint.abci.RequestCheckTx
	17, // 3: tmos.abci.v1alpha1.DeliverTxState.deliver_tx:type_name -> tendermint.abci.RequestDeliverTx
	18, // 4: tmos.abci.v1alpha1.ValidatorUpdates.validator_updates:type_name -> tendermint.abci.ValidatorUpdate
	19, // 5: tmos.abci.v1alpha1.EndBlockState.end_block:type_name -> tendermint.abci.RequestEndBlock
	15, // 6: tmos.abci.v1alpha1.MsgSetBeginBlockState.begin_block:type_name -> tendermint.abci.RequestBeginBlock
	16, // 7: tmos.abci.v1alpha1.MsgSetCheckTxState.check_tx:type_name -> tendermint.abci.RequestCheckTx
	17, // 8: tmos.abci.v1alpha1.MsgSetDeliverTxState.deliver_tx:type_name -> tendermint.abci.RequestDeliverTx
	19, // 9: tmos.abci.v1alpha1.MsgSetEndBlockState.end_block:type_name -> tendermint.abci.RequestEndBlock
	3,  // 10: tmos.abci.v1alpha1.MsgSetInitChain.init_chain_info:type_name -> tmos.abci.v1alpha1.InitChainInfo
	18, // 11: tmos.abci.v1alpha1.MsgSetValidatorUpdates.validator_updates:type_name -> tendermint.abci.ValidatorUpdate
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_core_abci_v1alpha1_abci_proto_init() }
func file_core_abci_v1alpha1_abci_proto_init() {
	if File_core_abci_v1alpha1_abci_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_core_abci_v1alpha1_abci_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CurrentBlock); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_core_abci_v1alpha1_abci_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Stage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_core_abci_v1alpha1_abci_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitChainInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_core_abci_v1alpha1_abci_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BeginBlockState); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_core_abci_v1alpha1_abci_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckTxState); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_core_abci_v1alpha1_abci_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeliverTxState); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_core_abci_v1alpha1_abci_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidatorUpdates); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_core_abci_v1alpha1_abci_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EndBlockState); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_core_abci_v1alpha1_abci_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgSetBeginBlockState); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_core_abci_v1alpha1_abci_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgSetCheckTxState); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_core_abci_v1alpha1_abci_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgSetDeliverTxState); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_core_abci_v1alpha1_abci_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgSetEndBlockState); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_core_abci_v1alpha1_abci_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgSetInitChain); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_core_abci_v1alpha1_abci_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgSetValidatorUpdates); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_core_abci_v1alpha1_abci_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_core_abci_v1alpha1_abci_proto_goTypes,
		DependencyIndexes: file_core_abci_v1alpha1_abci_proto_depIdxs,
		EnumInfos:         file_core_abci_v1alpha1_abci_proto_enumTypes,
		MessageInfos:      file_core_abci_v1alpha1_abci_proto_msgTypes,
	}.Build()
	File_core_abci_v1alpha1_abci_proto = out.File
	file_core_abci_v1alpha1_abci_proto_rawDesc = nil
	file_core_abci_v1alpha1_abci_proto_goTypes = nil
	file_core_abci_v1alpha1_abci_proto_depIdxs = nil
}
