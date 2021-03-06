// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.13.0
// source: testdata/testpb/testdata.proto

package testpb

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

type SimpleMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	A string   `protobuf:"bytes,1,opt,name=a,proto3" json:"a,omitempty"`
	B int64    `protobuf:"varint,2,opt,name=b,proto3" json:"b,omitempty"`
	C [][]byte `protobuf:"bytes,3,rep,name=c,proto3" json:"c,omitempty"`
}

func (x *SimpleMessage) Reset() {
	*x = SimpleMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testdata_testpb_testdata_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SimpleMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SimpleMessage) ProtoMessage() {}

func (x *SimpleMessage) ProtoReflect() protoreflect.Message {
	mi := &file_testdata_testpb_testdata_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SimpleMessage.ProtoReflect.Descriptor instead.
func (*SimpleMessage) Descriptor() ([]byte, []int) {
	return file_testdata_testpb_testdata_proto_rawDescGZIP(), []int{0}
}

func (x *SimpleMessage) GetA() string {
	if x != nil {
		return x.A
	}
	return ""
}

func (x *SimpleMessage) GetB() int64 {
	if x != nil {
		return x.B
	}
	return 0
}

func (x *SimpleMessage) GetC() [][]byte {
	if x != nil {
		return x.C
	}
	return nil
}

type SimpleMessageList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*SimpleMessage `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *SimpleMessageList) Reset() {
	*x = SimpleMessageList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testdata_testpb_testdata_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SimpleMessageList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SimpleMessageList) ProtoMessage() {}

func (x *SimpleMessageList) ProtoReflect() protoreflect.Message {
	mi := &file_testdata_testpb_testdata_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SimpleMessageList.ProtoReflect.Descriptor instead.
func (*SimpleMessageList) Descriptor() ([]byte, []int) {
	return file_testdata_testpb_testdata_proto_rawDescGZIP(), []int{1}
}

func (x *SimpleMessageList) GetList() []*SimpleMessage {
	if x != nil {
		return x.List
	}
	return nil
}

var File_testdata_testpb_testdata_proto protoreflect.FileDescriptor

var file_testdata_testpb_testdata_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x70,
	0x62, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x22, 0x39, 0x0a, 0x0d, 0x53, 0x69,
	0x6d, 0x70, 0x6c, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0c, 0x0a, 0x01, 0x61,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x61, 0x12, 0x0c, 0x0a, 0x01, 0x62, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x01, 0x62, 0x12, 0x0c, 0x0a, 0x01, 0x63, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0c, 0x52, 0x01, 0x63, 0x22, 0x40, 0x0a, 0x11, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x04, 0x6c, 0x69,
	0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x64,
	0x61, 0x74, 0x61, 0x2e, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x64, 0x79, 0x6d, 0x79, 0x6c, 0x6a, 0x61, 0x2f, 0x74,
	0x6d, 0x6f, 0x73, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x74, 0x65, 0x73,
	0x74, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_testdata_testpb_testdata_proto_rawDescOnce sync.Once
	file_testdata_testpb_testdata_proto_rawDescData = file_testdata_testpb_testdata_proto_rawDesc
)

func file_testdata_testpb_testdata_proto_rawDescGZIP() []byte {
	file_testdata_testpb_testdata_proto_rawDescOnce.Do(func() {
		file_testdata_testpb_testdata_proto_rawDescData = protoimpl.X.CompressGZIP(file_testdata_testpb_testdata_proto_rawDescData)
	})
	return file_testdata_testpb_testdata_proto_rawDescData
}

var file_testdata_testpb_testdata_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_testdata_testpb_testdata_proto_goTypes = []interface{}{
	(*SimpleMessage)(nil),     // 0: testdata.SimpleMessage
	(*SimpleMessageList)(nil), // 1: testdata.SimpleMessageList
}
var file_testdata_testpb_testdata_proto_depIdxs = []int32{
	0, // 0: testdata.SimpleMessageList.list:type_name -> testdata.SimpleMessage
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_testdata_testpb_testdata_proto_init() }
func file_testdata_testpb_testdata_proto_init() {
	if File_testdata_testpb_testdata_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_testdata_testpb_testdata_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SimpleMessage); i {
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
		file_testdata_testpb_testdata_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SimpleMessageList); i {
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
			RawDescriptor: file_testdata_testpb_testdata_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_testdata_testpb_testdata_proto_goTypes,
		DependencyIndexes: file_testdata_testpb_testdata_proto_depIdxs,
		MessageInfos:      file_testdata_testpb_testdata_proto_msgTypes,
	}.Build()
	File_testdata_testpb_testdata_proto = out.File
	file_testdata_testpb_testdata_proto_rawDesc = nil
	file_testdata_testpb_testdata_proto_goTypes = nil
	file_testdata_testpb_testdata_proto_depIdxs = nil
}
