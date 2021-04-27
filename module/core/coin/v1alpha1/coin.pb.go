// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.13.0
// source: module/core/coin/v1alpha1/coin.proto

package v1alpha1

import (
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

// Coin defines a token with a denomination and an amount.
type Coin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Denom  string `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty"`
	Amount string `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *Coin) Reset() {
	*x = Coin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apis_core_coin_v1alpha1_coin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Coin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Coin) ProtoMessage() {}

func (x *Coin) ProtoReflect() protoreflect.Message {
	mi := &file_apis_core_coin_v1alpha1_coin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Coin.ProtoReflect.Descriptor instead.
func (*Coin) Descriptor() ([]byte, []int) {
	return file_apis_core_coin_v1alpha1_coin_proto_rawDescGZIP(), []int{0}
}

func (x *Coin) GetDenom() string {
	if x != nil {
		return x.Denom
	}
	return ""
}

func (x *Coin) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

// DecCoin defines a token with a denomination and a decimal amount.
type DecCoin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Denom  string `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty"`
	Amount string `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *DecCoin) Reset() {
	*x = DecCoin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apis_core_coin_v1alpha1_coin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DecCoin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DecCoin) ProtoMessage() {}

func (x *DecCoin) ProtoReflect() protoreflect.Message {
	mi := &file_apis_core_coin_v1alpha1_coin_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DecCoin.ProtoReflect.Descriptor instead.
func (*DecCoin) Descriptor() ([]byte, []int) {
	return file_apis_core_coin_v1alpha1_coin_proto_rawDescGZIP(), []int{1}
}

func (x *DecCoin) GetDenom() string {
	if x != nil {
		return x.Denom
	}
	return ""
}

func (x *DecCoin) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

var File_apis_core_coin_v1alpha1_coin_proto protoreflect.FileDescriptor

var file_apis_core_coin_v1alpha1_coin_proto_rawDesc = []byte{
	0x0a, 0x22, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x63, 0x6f, 0x69, 0x6e,
	0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x74, 0x6d, 0x6f, 0x73, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x22, 0x34, 0x0a,
	0x04, 0x43, 0x6f, 0x69, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x22, 0x37, 0x0a, 0x07, 0x44, 0x65, 0x63, 0x43, 0x6f, 0x69, 0x6e, 0x12, 0x14,
	0x0a, 0x05, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x64,
	0x65, 0x6e, 0x6f, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x32, 0x5a, 0x30,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x64, 0x79, 0x6d, 0x79,
	0x6c, 0x6a, 0x61, 0x2f, 0x74, 0x6d, 0x6f, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f,
	0x72, 0x65, 0x2f, 0x63, 0x6f, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_apis_core_coin_v1alpha1_coin_proto_rawDescOnce sync.Once
	file_apis_core_coin_v1alpha1_coin_proto_rawDescData = file_apis_core_coin_v1alpha1_coin_proto_rawDesc
)

func file_apis_core_coin_v1alpha1_coin_proto_rawDescGZIP() []byte {
	file_apis_core_coin_v1alpha1_coin_proto_rawDescOnce.Do(func() {
		file_apis_core_coin_v1alpha1_coin_proto_rawDescData = protoimpl.X.CompressGZIP(file_apis_core_coin_v1alpha1_coin_proto_rawDescData)
	})
	return file_apis_core_coin_v1alpha1_coin_proto_rawDescData
}

var file_apis_core_coin_v1alpha1_coin_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_apis_core_coin_v1alpha1_coin_proto_goTypes = []interface{}{
	(*Coin)(nil),    // 0: tmos.core.coin.v1alpha1.Coin
	(*DecCoin)(nil), // 1: tmos.core.coin.v1alpha1.DecCoin
}
var file_apis_core_coin_v1alpha1_coin_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_apis_core_coin_v1alpha1_coin_proto_init() }
func file_apis_core_coin_v1alpha1_coin_proto_init() {
	if File_apis_core_coin_v1alpha1_coin_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_apis_core_coin_v1alpha1_coin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Coin); i {
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
		file_apis_core_coin_v1alpha1_coin_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DecCoin); i {
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
			RawDescriptor: file_apis_core_coin_v1alpha1_coin_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_apis_core_coin_v1alpha1_coin_proto_goTypes,
		DependencyIndexes: file_apis_core_coin_v1alpha1_coin_proto_depIdxs,
		MessageInfos:      file_apis_core_coin_v1alpha1_coin_proto_msgTypes,
	}.Build()
	File_apis_core_coin_v1alpha1_coin_proto = out.File
	file_apis_core_coin_v1alpha1_coin_proto_rawDesc = nil
	file_apis_core_coin_v1alpha1_coin_proto_goTypes = nil
	file_apis_core_coin_v1alpha1_coin_proto_depIdxs = nil
}