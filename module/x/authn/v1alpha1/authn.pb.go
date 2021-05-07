// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.13.0
// source: module/x/authn/v1alpha1/authn.proto

package v1alpha1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// - STATE TRANSITIONS -
// MsgCreateAccount is used to create a new account
type MsgCreateAccount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Account *Account `protobuf:"bytes,2,opt,name=account,proto3" json:"account,omitempty"`
}

func (x *MsgCreateAccount) Reset() {
	*x = MsgCreateAccount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_module_x_authn_v1alpha1_authn_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgCreateAccount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgCreateAccount) ProtoMessage() {}

func (x *MsgCreateAccount) ProtoReflect() protoreflect.Message {
	mi := &file_module_x_authn_v1alpha1_authn_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgCreateAccount.ProtoReflect.Descriptor instead.
func (*MsgCreateAccount) Descriptor() ([]byte, []int) {
	return file_module_x_authn_v1alpha1_authn_proto_rawDescGZIP(), []int{0}
}

func (x *MsgCreateAccount) GetAccount() *Account {
	if x != nil {
		return x.Account
	}
	return nil
}

// MsgUpdateAccount is used to update an existing account
type MsgUpdateAccount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Account *Account `protobuf:"bytes,2,opt,name=account,proto3" json:"account,omitempty"`
}

func (x *MsgUpdateAccount) Reset() {
	*x = MsgUpdateAccount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_module_x_authn_v1alpha1_authn_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgUpdateAccount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgUpdateAccount) ProtoMessage() {}

func (x *MsgUpdateAccount) ProtoReflect() protoreflect.Message {
	mi := &file_module_x_authn_v1alpha1_authn_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgUpdateAccount.ProtoReflect.Descriptor instead.
func (*MsgUpdateAccount) Descriptor() ([]byte, []int) {
	return file_module_x_authn_v1alpha1_authn_proto_rawDescGZIP(), []int{1}
}

func (x *MsgUpdateAccount) GetAccount() *Account {
	if x != nil {
		return x.Account
	}
	return nil
}

// MsgDeleteAccount is used to remove an account
type MsgDeleteAccount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Account *Account `protobuf:"bytes,2,opt,name=account,proto3" json:"account,omitempty"`
}

func (x *MsgDeleteAccount) Reset() {
	*x = MsgDeleteAccount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_module_x_authn_v1alpha1_authn_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgDeleteAccount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgDeleteAccount) ProtoMessage() {}

func (x *MsgDeleteAccount) ProtoReflect() protoreflect.Message {
	mi := &file_module_x_authn_v1alpha1_authn_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgDeleteAccount.ProtoReflect.Descriptor instead.
func (*MsgDeleteAccount) Descriptor() ([]byte, []int) {
	return file_module_x_authn_v1alpha1_authn_proto_rawDescGZIP(), []int{2}
}

func (x *MsgDeleteAccount) GetAccount() *Account {
	if x != nil {
		return x.Account
	}
	return nil
}

// Account defines an account
type Account struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address       string     `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	PubKey        *anypb.Any `protobuf:"bytes,2,opt,name=pub_key,json=pubKey,proto3" json:"pub_key,omitempty"`
	AccountNumber uint64     `protobuf:"varint,3,opt,name=account_number,json=accountNumber,proto3" json:"account_number,omitempty"`
	Sequence      uint64     `protobuf:"varint,4,opt,name=sequence,proto3" json:"sequence,omitempty"`
}

func (x *Account) Reset() {
	*x = Account{}
	if protoimpl.UnsafeEnabled {
		mi := &file_module_x_authn_v1alpha1_authn_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Account) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Account) ProtoMessage() {}

func (x *Account) ProtoReflect() protoreflect.Message {
	mi := &file_module_x_authn_v1alpha1_authn_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Account.ProtoReflect.Descriptor instead.
func (*Account) Descriptor() ([]byte, []int) {
	return file_module_x_authn_v1alpha1_authn_proto_rawDescGZIP(), []int{3}
}

func (x *Account) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Account) GetPubKey() *anypb.Any {
	if x != nil {
		return x.PubKey
	}
	return nil
}

func (x *Account) GetAccountNumber() uint64 {
	if x != nil {
		return x.AccountNumber
	}
	return 0
}

func (x *Account) GetSequence() uint64 {
	if x != nil {
		return x.Sequence
	}
	return 0
}

// CurrentAccountNumber is the state object containing the current account number
type CurrentAccountNumber struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// number is the current account number
	Number uint64 `protobuf:"varint,2,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *CurrentAccountNumber) Reset() {
	*x = CurrentAccountNumber{}
	if protoimpl.UnsafeEnabled {
		mi := &file_module_x_authn_v1alpha1_authn_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CurrentAccountNumber) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CurrentAccountNumber) ProtoMessage() {}

func (x *CurrentAccountNumber) ProtoReflect() protoreflect.Message {
	mi := &file_module_x_authn_v1alpha1_authn_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CurrentAccountNumber.ProtoReflect.Descriptor instead.
func (*CurrentAccountNumber) Descriptor() ([]byte, []int) {
	return file_module_x_authn_v1alpha1_authn_proto_rawDescGZIP(), []int{4}
}

func (x *CurrentAccountNumber) GetNumber() uint64 {
	if x != nil {
		return x.Number
	}
	return 0
}

// Params defines the parameters for the auth module.
type Params struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MaxMemoCharacters      uint64 `protobuf:"varint,1,opt,name=max_memo_characters,json=maxMemoCharacters,proto3" json:"max_memo_characters,omitempty"`
	TxSigLimit             uint64 `protobuf:"varint,2,opt,name=tx_sig_limit,json=txSigLimit,proto3" json:"tx_sig_limit,omitempty"`
	TxSizeCostPerByte      uint64 `protobuf:"varint,3,opt,name=tx_size_cost_per_byte,json=txSizeCostPerByte,proto3" json:"tx_size_cost_per_byte,omitempty"`
	SigVerifyCostEd25519   uint64 `protobuf:"varint,4,opt,name=sig_verify_cost_ed25519,json=sigVerifyCostEd25519,proto3" json:"sig_verify_cost_ed25519,omitempty"`
	SigVerifyCostSecp256K1 uint64 `protobuf:"varint,5,opt,name=sig_verify_cost_secp256k1,json=sigVerifyCostSecp256k1,proto3" json:"sig_verify_cost_secp256k1,omitempty"`
}

func (x *Params) Reset() {
	*x = Params{}
	if protoimpl.UnsafeEnabled {
		mi := &file_module_x_authn_v1alpha1_authn_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Params) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Params) ProtoMessage() {}

func (x *Params) ProtoReflect() protoreflect.Message {
	mi := &file_module_x_authn_v1alpha1_authn_proto_msgTypes[5]
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
	return file_module_x_authn_v1alpha1_authn_proto_rawDescGZIP(), []int{5}
}

func (x *Params) GetMaxMemoCharacters() uint64 {
	if x != nil {
		return x.MaxMemoCharacters
	}
	return 0
}

func (x *Params) GetTxSigLimit() uint64 {
	if x != nil {
		return x.TxSigLimit
	}
	return 0
}

func (x *Params) GetTxSizeCostPerByte() uint64 {
	if x != nil {
		return x.TxSizeCostPerByte
	}
	return 0
}

func (x *Params) GetSigVerifyCostEd25519() uint64 {
	if x != nil {
		return x.SigVerifyCostEd25519
	}
	return 0
}

func (x *Params) GetSigVerifyCostSecp256K1() uint64 {
	if x != nil {
		return x.SigVerifyCostSecp256K1
	}
	return 0
}

var File_module_x_authn_v1alpha1_authn_proto protoreflect.FileDescriptor

var file_module_x_authn_v1alpha1_authn_proto_rawDesc = []byte{
	0x0a, 0x23, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2f, 0x78, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x6e,
	0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x74, 0x6d, 0x6f, 0x73, 0x2e, 0x78, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x19, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4c, 0x0a, 0x10, 0x4d, 0x73, 0x67, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x38, 0x0a, 0x07, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x74,
	0x6d, 0x6f, 0x73, 0x2e, 0x78, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x07, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x4c, 0x0a, 0x10, 0x4d, 0x73, 0x67, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x38, 0x0a, 0x07, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x74, 0x6d, 0x6f,
	0x73, 0x2e, 0x78, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x22, 0x4c, 0x0a, 0x10, 0x4d, 0x73, 0x67, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x38, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x74, 0x6d, 0x6f, 0x73, 0x2e,
	0x78, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x22, 0x95, 0x01, 0x0a, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x18, 0x0a,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x2d, 0x0a, 0x07, 0x70, 0x75, 0x62, 0x5f, 0x6b,
	0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x06,
	0x70, 0x75, 0x62, 0x4b, 0x65, 0x79, 0x12, 0x25, 0x0a, 0x0e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1a, 0x0a,
	0x08, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x08, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x22, 0x2e, 0x0a, 0x14, 0x43, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0xfe, 0x01, 0x0a, 0x06, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x73, 0x12, 0x2e, 0x0a, 0x13, 0x6d, 0x61, 0x78, 0x5f, 0x6d, 0x65, 0x6d, 0x6f,
	0x5f, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x11, 0x6d, 0x61, 0x78, 0x4d, 0x65, 0x6d, 0x6f, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63,
	0x74, 0x65, 0x72, 0x73, 0x12, 0x20, 0x0a, 0x0c, 0x74, 0x78, 0x5f, 0x73, 0x69, 0x67, 0x5f, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x74, 0x78, 0x53, 0x69,
	0x67, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x30, 0x0a, 0x15, 0x74, 0x78, 0x5f, 0x73, 0x69, 0x7a,
	0x65, 0x5f, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x11, 0x74, 0x78, 0x53, 0x69, 0x7a, 0x65, 0x43, 0x6f, 0x73,
	0x74, 0x50, 0x65, 0x72, 0x42, 0x79, 0x74, 0x65, 0x12, 0x35, 0x0a, 0x17, 0x73, 0x69, 0x67, 0x5f,
	0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x5f, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x65, 0x64, 0x32, 0x35,
	0x35, 0x31, 0x39, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x14, 0x73, 0x69, 0x67, 0x56, 0x65,
	0x72, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x73, 0x74, 0x45, 0x64, 0x32, 0x35, 0x35, 0x31, 0x39, 0x12,
	0x39, 0x0a, 0x19, 0x73, 0x69, 0x67, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x5f, 0x63, 0x6f,
	0x73, 0x74, 0x5f, 0x73, 0x65, 0x63, 0x70, 0x32, 0x35, 0x36, 0x6b, 0x31, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x16, 0x73, 0x69, 0x67, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x73,
	0x74, 0x53, 0x65, 0x63, 0x70, 0x32, 0x35, 0x36, 0x6b, 0x31, 0x42, 0x32, 0x5a, 0x30, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x64, 0x79, 0x6d, 0x79, 0x6c, 0x6a,
	0x61, 0x2f, 0x74, 0x6d, 0x6f, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2f, 0x78, 0x2f,
	0x61, 0x75, 0x74, 0x68, 0x6e, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_module_x_authn_v1alpha1_authn_proto_rawDescOnce sync.Once
	file_module_x_authn_v1alpha1_authn_proto_rawDescData = file_module_x_authn_v1alpha1_authn_proto_rawDesc
)

func file_module_x_authn_v1alpha1_authn_proto_rawDescGZIP() []byte {
	file_module_x_authn_v1alpha1_authn_proto_rawDescOnce.Do(func() {
		file_module_x_authn_v1alpha1_authn_proto_rawDescData = protoimpl.X.CompressGZIP(file_module_x_authn_v1alpha1_authn_proto_rawDescData)
	})
	return file_module_x_authn_v1alpha1_authn_proto_rawDescData
}

var file_module_x_authn_v1alpha1_authn_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_module_x_authn_v1alpha1_authn_proto_goTypes = []interface{}{
	(*MsgCreateAccount)(nil),     // 0: tmos.x.authn.v1alpha1.MsgCreateAccount
	(*MsgUpdateAccount)(nil),     // 1: tmos.x.authn.v1alpha1.MsgUpdateAccount
	(*MsgDeleteAccount)(nil),     // 2: tmos.x.authn.v1alpha1.MsgDeleteAccount
	(*Account)(nil),              // 3: tmos.x.authn.v1alpha1.Account
	(*CurrentAccountNumber)(nil), // 4: tmos.x.authn.v1alpha1.CurrentAccountNumber
	(*Params)(nil),               // 5: tmos.x.authn.v1alpha1.Params
	(*anypb.Any)(nil),            // 6: google.protobuf.Any
}
var file_module_x_authn_v1alpha1_authn_proto_depIdxs = []int32{
	3, // 0: tmos.x.authn.v1alpha1.MsgCreateAccount.account:type_name -> tmos.x.authn.v1alpha1.Account
	3, // 1: tmos.x.authn.v1alpha1.MsgUpdateAccount.account:type_name -> tmos.x.authn.v1alpha1.Account
	3, // 2: tmos.x.authn.v1alpha1.MsgDeleteAccount.account:type_name -> tmos.x.authn.v1alpha1.Account
	6, // 3: tmos.x.authn.v1alpha1.Account.pub_key:type_name -> google.protobuf.Any
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_module_x_authn_v1alpha1_authn_proto_init() }
func file_module_x_authn_v1alpha1_authn_proto_init() {
	if File_module_x_authn_v1alpha1_authn_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_module_x_authn_v1alpha1_authn_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgCreateAccount); i {
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
		file_module_x_authn_v1alpha1_authn_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgUpdateAccount); i {
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
		file_module_x_authn_v1alpha1_authn_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgDeleteAccount); i {
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
		file_module_x_authn_v1alpha1_authn_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Account); i {
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
		file_module_x_authn_v1alpha1_authn_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CurrentAccountNumber); i {
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
		file_module_x_authn_v1alpha1_authn_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_module_x_authn_v1alpha1_authn_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_module_x_authn_v1alpha1_authn_proto_goTypes,
		DependencyIndexes: file_module_x_authn_v1alpha1_authn_proto_depIdxs,
		MessageInfos:      file_module_x_authn_v1alpha1_authn_proto_msgTypes,
	}.Build()
	File_module_x_authn_v1alpha1_authn_proto = out.File
	file_module_x_authn_v1alpha1_authn_proto_rawDesc = nil
	file_module_x_authn_v1alpha1_authn_proto_goTypes = nil
	file_module_x_authn_v1alpha1_authn_proto_depIdxs = nil
}
