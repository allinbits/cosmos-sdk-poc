// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.13.0
// source: module/x/authn/tx/testdata/test.proto

package testdata

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

type HasUnknowns struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	A string `protobuf:"bytes,1,opt,name=a,proto3" json:"a,omitempty"`
}

func (x *HasUnknowns) Reset() {
	*x = HasUnknowns{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apis_x_authn_tx_testdata_test_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HasUnknowns) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HasUnknowns) ProtoMessage() {}

func (x *HasUnknowns) ProtoReflect() protoreflect.Message {
	mi := &file_apis_x_authn_tx_testdata_test_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HasUnknowns.ProtoReflect.Descriptor instead.
func (*HasUnknowns) Descriptor() ([]byte, []int) {
	return file_apis_x_authn_tx_testdata_test_proto_rawDescGZIP(), []int{0}
}

func (x *HasUnknowns) GetA() string {
	if x != nil {
		return x.A
	}
	return ""
}

type WithoutUnknowns struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	A string `protobuf:"bytes,1,opt,name=a,proto3" json:"a,omitempty"`
	B string `protobuf:"bytes,2,opt,name=b,proto3" json:"b,omitempty"`
}

func (x *WithoutUnknowns) Reset() {
	*x = WithoutUnknowns{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apis_x_authn_tx_testdata_test_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WithoutUnknowns) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WithoutUnknowns) ProtoMessage() {}

func (x *WithoutUnknowns) ProtoReflect() protoreflect.Message {
	mi := &file_apis_x_authn_tx_testdata_test_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WithoutUnknowns.ProtoReflect.Descriptor instead.
func (*WithoutUnknowns) Descriptor() ([]byte, []int) {
	return file_apis_x_authn_tx_testdata_test_proto_rawDescGZIP(), []int{1}
}

func (x *WithoutUnknowns) GetA() string {
	if x != nil {
		return x.A
	}
	return ""
}

func (x *WithoutUnknowns) GetB() string {
	if x != nil {
		return x.B
	}
	return ""
}

var File_apis_x_authn_tx_testdata_test_proto protoreflect.FileDescriptor

var file_apis_x_authn_tx_testdata_test_proto_rawDesc = []byte{
	0x0a, 0x23, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x78, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x6e, 0x2f, 0x74,
	0x78, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x75, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x2e, 0x74,
	0x65, 0x73, 0x74, 0x22, 0x1b, 0x0a, 0x0b, 0x48, 0x61, 0x73, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77,
	0x6e, 0x73, 0x12, 0x0c, 0x0a, 0x01, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x61,
	0x22, 0x2d, 0x0a, 0x0f, 0x57, 0x69, 0x74, 0x68, 0x6f, 0x75, 0x74, 0x55, 0x6e, 0x6b, 0x6e, 0x6f,
	0x77, 0x6e, 0x73, 0x12, 0x0c, 0x0a, 0x01, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01,
	0x61, 0x12, 0x0c, 0x0a, 0x01, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x62, 0x42,
	0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x64,
	0x79, 0x6d, 0x79, 0x6c, 0x6a, 0x61, 0x2f, 0x74, 0x6d, 0x6f, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x73,
	0x2f, 0x78, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x6e, 0x2f, 0x74, 0x78, 0x2f, 0x74, 0x65, 0x73, 0x74,
	0x64, 0x61, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_apis_x_authn_tx_testdata_test_proto_rawDescOnce sync.Once
	file_apis_x_authn_tx_testdata_test_proto_rawDescData = file_apis_x_authn_tx_testdata_test_proto_rawDesc
)

func file_apis_x_authn_tx_testdata_test_proto_rawDescGZIP() []byte {
	file_apis_x_authn_tx_testdata_test_proto_rawDescOnce.Do(func() {
		file_apis_x_authn_tx_testdata_test_proto_rawDescData = protoimpl.X.CompressGZIP(file_apis_x_authn_tx_testdata_test_proto_rawDescData)
	})
	return file_apis_x_authn_tx_testdata_test_proto_rawDescData
}

var file_apis_x_authn_tx_testdata_test_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_apis_x_authn_tx_testdata_test_proto_goTypes = []interface{}{
	(*HasUnknowns)(nil),     // 0: unknown.test.HasUnknowns
	(*WithoutUnknowns)(nil), // 1: unknown.test.WithoutUnknowns
}
var file_apis_x_authn_tx_testdata_test_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_apis_x_authn_tx_testdata_test_proto_init() }
func file_apis_x_authn_tx_testdata_test_proto_init() {
	if File_apis_x_authn_tx_testdata_test_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_apis_x_authn_tx_testdata_test_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HasUnknowns); i {
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
		file_apis_x_authn_tx_testdata_test_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WithoutUnknowns); i {
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
			RawDescriptor: file_apis_x_authn_tx_testdata_test_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_apis_x_authn_tx_testdata_test_proto_goTypes,
		DependencyIndexes: file_apis_x_authn_tx_testdata_test_proto_depIdxs,
		MessageInfos:      file_apis_x_authn_tx_testdata_test_proto_msgTypes,
	}.Build()
	File_apis_x_authn_tx_testdata_test_proto = out.File
	file_apis_x_authn_tx_testdata_test_proto_rawDesc = nil
	file_apis_x_authn_tx_testdata_test_proto_goTypes = nil
	file_apis_x_authn_tx_testdata_test_proto_depIdxs = nil
}