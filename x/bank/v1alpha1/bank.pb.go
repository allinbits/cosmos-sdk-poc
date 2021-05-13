// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.13.0
// source: x/bank/v1alpha1/bank.proto

package v1alpha1

import (
	v1alpha1 "github.com/fdymylja/tmos/core/coin/v1alpha1"
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

type MsgSendCoins struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FromAddress string           `protobuf:"bytes,1,opt,name=from_address,json=fromAddress,proto3" json:"from_address,omitempty"`
	ToAddress   string           `protobuf:"bytes,2,opt,name=to_address,json=toAddress,proto3" json:"to_address,omitempty"`
	Amount      []*v1alpha1.Coin `protobuf:"bytes,3,rep,name=amount,proto3" json:"amount,omitempty"`
}

func (x *MsgSendCoins) Reset() {
	*x = MsgSendCoins{}
	if protoimpl.UnsafeEnabled {
		mi := &file_x_bank_v1alpha1_bank_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgSendCoins) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgSendCoins) ProtoMessage() {}

func (x *MsgSendCoins) ProtoReflect() protoreflect.Message {
	mi := &file_x_bank_v1alpha1_bank_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgSendCoins.ProtoReflect.Descriptor instead.
func (*MsgSendCoins) Descriptor() ([]byte, []int) {
	return file_x_bank_v1alpha1_bank_proto_rawDescGZIP(), []int{0}
}

func (x *MsgSendCoins) GetFromAddress() string {
	if x != nil {
		return x.FromAddress
	}
	return ""
}

func (x *MsgSendCoins) GetToAddress() string {
	if x != nil {
		return x.ToAddress
	}
	return ""
}

func (x *MsgSendCoins) GetAmount() []*v1alpha1.Coin {
	if x != nil {
		return x.Amount
	}
	return nil
}

type MsgSetBalance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string           `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Amount  []*v1alpha1.Coin `protobuf:"bytes,2,rep,name=amount,proto3" json:"amount,omitempty"`
}

func (x *MsgSetBalance) Reset() {
	*x = MsgSetBalance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_x_bank_v1alpha1_bank_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgSetBalance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgSetBalance) ProtoMessage() {}

func (x *MsgSetBalance) ProtoReflect() protoreflect.Message {
	mi := &file_x_bank_v1alpha1_bank_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgSetBalance.ProtoReflect.Descriptor instead.
func (*MsgSetBalance) Descriptor() ([]byte, []int) {
	return file_x_bank_v1alpha1_bank_proto_rawDescGZIP(), []int{1}
}

func (x *MsgSetBalance) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *MsgSetBalance) GetAmount() []*v1alpha1.Coin {
	if x != nil {
		return x.Amount
	}
	return nil
}

type Balance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// address identifies the user who owns balance TODO: i'd rename this user...
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// balance represents the balance of the user
	Balance []*v1alpha1.Coin `protobuf:"bytes,2,rep,name=balance,proto3" json:"balance,omitempty"`
}

func (x *Balance) Reset() {
	*x = Balance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_x_bank_v1alpha1_bank_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Balance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Balance) ProtoMessage() {}

func (x *Balance) ProtoReflect() protoreflect.Message {
	mi := &file_x_bank_v1alpha1_bank_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Balance.ProtoReflect.Descriptor instead.
func (*Balance) Descriptor() ([]byte, []int) {
	return file_x_bank_v1alpha1_bank_proto_rawDescGZIP(), []int{2}
}

func (x *Balance) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Balance) GetBalance() []*v1alpha1.Coin {
	if x != nil {
		return x.Balance
	}
	return nil
}

// GenesisState defines the bank module's genesis state.
type GenesisState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// params defines all the paramaters of the module.
	Params *Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params,omitempty"`
	// balances is an array containing the balances of all the accounts.
	Balances []*Balance `protobuf:"bytes,2,rep,name=balances,proto3" json:"balances,omitempty"`
	// supply represents the total supply. If it is left empty, then supply will be calculated based on the provided
	// balances. Otherwise, it will be used to validate that the sum of the balances equals this amount.
	Balance []*v1alpha1.Coin `protobuf:"bytes,3,rep,name=balance,proto3" json:"balance,omitempty"`
	// denom_metadata defines the metadata of the differents coins.
	DenomMetadata []*Metadata `protobuf:"bytes,4,rep,name=denom_metadata,json=denomMetadata,proto3" json:"denom_metadata,omitempty"`
}

func (x *GenesisState) Reset() {
	*x = GenesisState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_x_bank_v1alpha1_bank_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenesisState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenesisState) ProtoMessage() {}

func (x *GenesisState) ProtoReflect() protoreflect.Message {
	mi := &file_x_bank_v1alpha1_bank_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenesisState.ProtoReflect.Descriptor instead.
func (*GenesisState) Descriptor() ([]byte, []int) {
	return file_x_bank_v1alpha1_bank_proto_rawDescGZIP(), []int{3}
}

func (x *GenesisState) GetParams() *Params {
	if x != nil {
		return x.Params
	}
	return nil
}

func (x *GenesisState) GetBalances() []*Balance {
	if x != nil {
		return x.Balances
	}
	return nil
}

func (x *GenesisState) GetBalance() []*v1alpha1.Coin {
	if x != nil {
		return x.Balance
	}
	return nil
}

func (x *GenesisState) GetDenomMetadata() []*Metadata {
	if x != nil {
		return x.DenomMetadata
	}
	return nil
}

// Params defines the parameters for the bank module.
type Params struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SendEnabled        []*SendEnabled `protobuf:"bytes,1,rep,name=send_enabled,json=sendEnabled,proto3" json:"send_enabled,omitempty"`
	DefaultSendEnabled bool           `protobuf:"varint,2,opt,name=default_send_enabled,json=defaultSendEnabled,proto3" json:"default_send_enabled,omitempty"`
}

func (x *Params) Reset() {
	*x = Params{}
	if protoimpl.UnsafeEnabled {
		mi := &file_x_bank_v1alpha1_bank_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Params) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Params) ProtoMessage() {}

func (x *Params) ProtoReflect() protoreflect.Message {
	mi := &file_x_bank_v1alpha1_bank_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Params.ProtoReflect.Descriptor instead.
func (*Params) Descriptor() ([]byte, []int) {
	return file_x_bank_v1alpha1_bank_proto_rawDescGZIP(), []int{4}
}

func (x *Params) GetSendEnabled() []*SendEnabled {
	if x != nil {
		return x.SendEnabled
	}
	return nil
}

func (x *Params) GetDefaultSendEnabled() bool {
	if x != nil {
		return x.DefaultSendEnabled
	}
	return false
}

// SendEnabled maps coin denom to a send_enabled status (whether a denom is
// sendable).
type SendEnabled struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Denom   string `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty"`
	Enabled bool   `protobuf:"varint,2,opt,name=enabled,proto3" json:"enabled,omitempty"`
}

func (x *SendEnabled) Reset() {
	*x = SendEnabled{}
	if protoimpl.UnsafeEnabled {
		mi := &file_x_bank_v1alpha1_bank_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendEnabled) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendEnabled) ProtoMessage() {}

func (x *SendEnabled) ProtoReflect() protoreflect.Message {
	mi := &file_x_bank_v1alpha1_bank_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendEnabled.ProtoReflect.Descriptor instead.
func (*SendEnabled) Descriptor() ([]byte, []int) {
	return file_x_bank_v1alpha1_bank_proto_rawDescGZIP(), []int{5}
}

func (x *SendEnabled) GetDenom() string {
	if x != nil {
		return x.Denom
	}
	return ""
}

func (x *SendEnabled) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

// Metadata represents a struct that describes
// a basic token.
type Metadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Description string `protobuf:"bytes,1,opt,name=description,proto3" json:"description,omitempty"`
	// denom_units represents the list of DenomUnit's for a given coin
	DenomUnits []*DenomUnit `protobuf:"bytes,2,rep,name=denom_units,json=denomUnits,proto3" json:"denom_units,omitempty"`
	// base represents the base denom (should be the DenomUnit with exponent = 0).
	Base string `protobuf:"bytes,3,opt,name=base,proto3" json:"base,omitempty"`
	// display indicates the suggested denom that should be
	// displayed in clients.
	Display string `protobuf:"bytes,4,opt,name=display,proto3" json:"display,omitempty"`
	// name defines the name of the token (eg: Cosmos Atom)
	Name string `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	// symbol is the token symbol usually shown on exchanges (eg: ATOM). This can
	// be the same as the display.
	Symbol string `protobuf:"bytes,6,opt,name=symbol,proto3" json:"symbol,omitempty"`
}

func (x *Metadata) Reset() {
	*x = Metadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_x_bank_v1alpha1_bank_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Metadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metadata) ProtoMessage() {}

func (x *Metadata) ProtoReflect() protoreflect.Message {
	mi := &file_x_bank_v1alpha1_bank_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Metadata.ProtoReflect.Descriptor instead.
func (*Metadata) Descriptor() ([]byte, []int) {
	return file_x_bank_v1alpha1_bank_proto_rawDescGZIP(), []int{6}
}

func (x *Metadata) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Metadata) GetDenomUnits() []*DenomUnit {
	if x != nil {
		return x.DenomUnits
	}
	return nil
}

func (x *Metadata) GetBase() string {
	if x != nil {
		return x.Base
	}
	return ""
}

func (x *Metadata) GetDisplay() string {
	if x != nil {
		return x.Display
	}
	return ""
}

func (x *Metadata) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Metadata) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

// DenomUnit represents a struct that describes a given
// denomination unit of the basic token.
type DenomUnit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// denom represents the string name of the given denom unit (e.g uatom).
	Denom string `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty"`
	// exponent represents power of 10 exponent that one must
	// raise the base_denom to in order to equal the given DenomUnit's denom
	// 1 denom = 1^exponent base_denom
	// (e.g. with a base_denom of uatom, one can create a DenomUnit of 'atom' with
	// exponent = 6, thus: 1 atom = 10^6 uatom).
	Exponent uint32 `protobuf:"varint,2,opt,name=exponent,proto3" json:"exponent,omitempty"`
	// aliases is a list of string aliases for the given denom
	Aliases []string `protobuf:"bytes,3,rep,name=aliases,proto3" json:"aliases,omitempty"`
}

func (x *DenomUnit) Reset() {
	*x = DenomUnit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_x_bank_v1alpha1_bank_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DenomUnit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DenomUnit) ProtoMessage() {}

func (x *DenomUnit) ProtoReflect() protoreflect.Message {
	mi := &file_x_bank_v1alpha1_bank_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DenomUnit.ProtoReflect.Descriptor instead.
func (*DenomUnit) Descriptor() ([]byte, []int) {
	return file_x_bank_v1alpha1_bank_proto_rawDescGZIP(), []int{7}
}

func (x *DenomUnit) GetDenom() string {
	if x != nil {
		return x.Denom
	}
	return ""
}

func (x *DenomUnit) GetExponent() uint32 {
	if x != nil {
		return x.Exponent
	}
	return 0
}

func (x *DenomUnit) GetAliases() []string {
	if x != nil {
		return x.Aliases
	}
	return nil
}

var File_x_bank_v1alpha1_bank_proto protoreflect.FileDescriptor

var file_x_bank_v1alpha1_bank_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x78, 0x2f, 0x62, 0x61, 0x6e, 0x6b, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2f, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x74, 0x6d,
	0x6f, 0x73, 0x2e, 0x78, 0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x1a, 0x1d, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x63, 0x6f, 0x69, 0x6e, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x87, 0x01, 0x0a, 0x0c, 0x4d, 0x73, 0x67, 0x53, 0x65, 0x6e, 0x64, 0x43, 0x6f, 0x69,
	0x6e, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x66, 0x72, 0x6f, 0x6d, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x6f, 0x5f, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x6f, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x35, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x74, 0x6d, 0x6f, 0x73, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x43,
	0x6f, 0x69, 0x6e, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x60, 0x0a, 0x0d, 0x4d,
	0x73, 0x67, 0x53, 0x65, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x35, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x74, 0x6d, 0x6f, 0x73, 0x2e, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x43, 0x6f, 0x69, 0x6e, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x5c, 0x0a,
	0x07, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x37, 0x0a, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x74, 0x6d, 0x6f, 0x73, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x43, 0x6f,
	0x69, 0x6e, 0x52, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x22, 0xff, 0x01, 0x0a, 0x0c,
	0x47, 0x65, 0x6e, 0x65, 0x73, 0x69, 0x73, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x34, 0x0a, 0x06,
	0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x74,
	0x6d, 0x6f, 0x73, 0x2e, 0x78, 0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x12, 0x39, 0x0a, 0x08, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x74, 0x6d, 0x6f, 0x73, 0x2e, 0x78, 0x2e, 0x62, 0x61,
	0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x42, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x65, 0x52, 0x08, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x12, 0x37, 0x0a,
	0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d,
	0x2e, 0x74, 0x6d, 0x6f, 0x73, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x43, 0x6f, 0x69, 0x6e, 0x52, 0x07, 0x62,
	0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x45, 0x0a, 0x0e, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x5f,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e,
	0x2e, 0x74, 0x6d, 0x6f, 0x73, 0x2e, 0x78, 0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x0d,
	0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x22, 0x80, 0x01,
	0x0a, 0x06, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x44, 0x0a, 0x0c, 0x73, 0x65, 0x6e, 0x64,
	0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21,
	0x2e, 0x74, 0x6d, 0x6f, 0x73, 0x2e, 0x78, 0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x64, 0x52, 0x0b, 0x73, 0x65, 0x6e, 0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x30,
	0x0a, 0x14, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x73, 0x65, 0x6e, 0x64, 0x5f, 0x65,
	0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x64, 0x65,
	0x66, 0x61, 0x75, 0x6c, 0x74, 0x53, 0x65, 0x6e, 0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64,
	0x22, 0x3d, 0x0a, 0x0b, 0x53, 0x65, 0x6e, 0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x22,
	0xc8, 0x01, 0x0a, 0x08, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x20, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x40,
	0x0a, 0x0b, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x5f, 0x75, 0x6e, 0x69, 0x74, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x74, 0x6d, 0x6f, 0x73, 0x2e, 0x78, 0x2e, 0x62, 0x61, 0x6e,
	0x6b, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x44, 0x65, 0x6e, 0x6f, 0x6d,
	0x55, 0x6e, 0x69, 0x74, 0x52, 0x0a, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x55, 0x6e, 0x69, 0x74, 0x73,
	0x12, 0x12, 0x0a, 0x04, 0x62, 0x61, 0x73, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x62, 0x61, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x22, 0x57, 0x0a, 0x09, 0x44, 0x65,
	0x6e, 0x6f, 0x6d, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x65, 0x6e, 0x6f, 0x6d,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x12, 0x1a, 0x0a,
	0x08, 0x65, 0x78, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x08, 0x65, 0x78, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x6c, 0x69,
	0x61, 0x73, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x61, 0x6c, 0x69, 0x61,
	0x73, 0x65, 0x73, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x66, 0x64, 0x79, 0x6d, 0x79, 0x6c, 0x6a, 0x61, 0x2f, 0x74, 0x6d, 0x6f, 0x73, 0x2f,
	0x78, 0x2f, 0x62, 0x61, 0x6e, 0x6b, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_x_bank_v1alpha1_bank_proto_rawDescOnce sync.Once
	file_x_bank_v1alpha1_bank_proto_rawDescData = file_x_bank_v1alpha1_bank_proto_rawDesc
)

func file_x_bank_v1alpha1_bank_proto_rawDescGZIP() []byte {
	file_x_bank_v1alpha1_bank_proto_rawDescOnce.Do(func() {
		file_x_bank_v1alpha1_bank_proto_rawDescData = protoimpl.X.CompressGZIP(file_x_bank_v1alpha1_bank_proto_rawDescData)
	})
	return file_x_bank_v1alpha1_bank_proto_rawDescData
}

var file_x_bank_v1alpha1_bank_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_x_bank_v1alpha1_bank_proto_goTypes = []interface{}{
	(*MsgSendCoins)(nil),  // 0: tmos.x.bank.v1alpha1.MsgSendCoins
	(*MsgSetBalance)(nil), // 1: tmos.x.bank.v1alpha1.MsgSetBalance
	(*Balance)(nil),       // 2: tmos.x.bank.v1alpha1.Balance
	(*GenesisState)(nil),  // 3: tmos.x.bank.v1alpha1.GenesisState
	(*Params)(nil),        // 4: tmos.x.bank.v1alpha1.Params
	(*SendEnabled)(nil),   // 5: tmos.x.bank.v1alpha1.SendEnabled
	(*Metadata)(nil),      // 6: tmos.x.bank.v1alpha1.Metadata
	(*DenomUnit)(nil),     // 7: tmos.x.bank.v1alpha1.DenomUnit
	(*v1alpha1.Coin)(nil), // 8: tmos.core.coin.v1alpha1.Coin
}
var file_x_bank_v1alpha1_bank_proto_depIdxs = []int32{
	8, // 0: tmos.x.bank.v1alpha1.MsgSendCoins.amount:type_name -> tmos.core.coin.v1alpha1.Coin
	8, // 1: tmos.x.bank.v1alpha1.MsgSetBalance.amount:type_name -> tmos.core.coin.v1alpha1.Coin
	8, // 2: tmos.x.bank.v1alpha1.Balance.balance:type_name -> tmos.core.coin.v1alpha1.Coin
	4, // 3: tmos.x.bank.v1alpha1.GenesisState.params:type_name -> tmos.x.bank.v1alpha1.Params
	2, // 4: tmos.x.bank.v1alpha1.GenesisState.balances:type_name -> tmos.x.bank.v1alpha1.Balance
	8, // 5: tmos.x.bank.v1alpha1.GenesisState.balance:type_name -> tmos.core.coin.v1alpha1.Coin
	6, // 6: tmos.x.bank.v1alpha1.GenesisState.denom_metadata:type_name -> tmos.x.bank.v1alpha1.Metadata
	5, // 7: tmos.x.bank.v1alpha1.Params.send_enabled:type_name -> tmos.x.bank.v1alpha1.SendEnabled
	7, // 8: tmos.x.bank.v1alpha1.Metadata.denom_units:type_name -> tmos.x.bank.v1alpha1.DenomUnit
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_x_bank_v1alpha1_bank_proto_init() }
func file_x_bank_v1alpha1_bank_proto_init() {
	if File_x_bank_v1alpha1_bank_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_x_bank_v1alpha1_bank_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgSendCoins); i {
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
		file_x_bank_v1alpha1_bank_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgSetBalance); i {
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
		file_x_bank_v1alpha1_bank_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Balance); i {
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
		file_x_bank_v1alpha1_bank_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenesisState); i {
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
		file_x_bank_v1alpha1_bank_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Params); i {
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
		file_x_bank_v1alpha1_bank_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendEnabled); i {
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
		file_x_bank_v1alpha1_bank_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Metadata); i {
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
		file_x_bank_v1alpha1_bank_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DenomUnit); i {
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
			RawDescriptor: file_x_bank_v1alpha1_bank_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_x_bank_v1alpha1_bank_proto_goTypes,
		DependencyIndexes: file_x_bank_v1alpha1_bank_proto_depIdxs,
		MessageInfos:      file_x_bank_v1alpha1_bank_proto_msgTypes,
	}.Build()
	File_x_bank_v1alpha1_bank_proto = out.File
	file_x_bank_v1alpha1_bank_proto_rawDesc = nil
	file_x_bank_v1alpha1_bank_proto_goTypes = nil
	file_x_bank_v1alpha1_bank_proto_depIdxs = nil
}
