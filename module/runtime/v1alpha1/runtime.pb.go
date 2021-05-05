// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.13.0
// source: module/runtime/v1alpha1/runtime.proto

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

// Verb defines an action that can be performed on the runtime by external entities
type Verb int32

const (
	// Unknown defines an unknown verb
	Verb_Unknown Verb = 0
	// Get identifies the GET operation
	Verb_Get Verb = 1
	// List identifies the LIST operation
	Verb_List Verb = 2
	// Create identifies the CREATE operation
	Verb_Create Verb = 3
	// Update identifies the UPDATE operation
	Verb_Update Verb = 4
	// Delete identifies the DELETE operation
	Verb_Delete Verb = 5
	// Deliver identifies the DELIVER operation
	Verb_Deliver Verb = 6
)

// Enum value maps for Verb.
var (
	Verb_name = map[int32]string{
		0: "Unknown",
		1: "Get",
		2: "List",
		3: "Create",
		4: "Update",
		5: "Delete",
		6: "Deliver",
	}
	Verb_value = map[string]int32{
		"Unknown": 0,
		"Get":     1,
		"List":    2,
		"Create":  3,
		"Update":  4,
		"Delete":  5,
		"Deliver": 6,
	}
)

func (x Verb) Enum() *Verb {
	p := new(Verb)
	*p = x
	return p
}

func (x Verb) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Verb) Descriptor() protoreflect.EnumDescriptor {
	return file_module_runtime_v1alpha1_runtime_proto_enumTypes[0].Descriptor()
}

func (Verb) Type() protoreflect.EnumType {
	return &file_module_runtime_v1alpha1_runtime_proto_enumTypes[0]
}

func (x Verb) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Verb.Descriptor instead.
func (Verb) EnumDescriptor() ([]byte, []int) {
	return file_module_runtime_v1alpha1_runtime_proto_rawDescGZIP(), []int{0}
}

// StateObjectsList provides the list of state objects
// represented as protobuf.Fullname(meta.StateObject)
type StateObjectsList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StateObjects []string `protobuf:"bytes,1,rep,name=state_objects,json=stateObjects,proto3" json:"state_objects,omitempty"`
}

func (x *StateObjectsList) Reset() {
	*x = StateObjectsList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_module_runtime_v1alpha1_runtime_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StateObjectsList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StateObjectsList) ProtoMessage() {}

func (x *StateObjectsList) ProtoReflect() protoreflect.Message {
	mi := &file_module_runtime_v1alpha1_runtime_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StateObjectsList.ProtoReflect.Descriptor instead.
func (*StateObjectsList) Descriptor() ([]byte, []int) {
	return file_module_runtime_v1alpha1_runtime_proto_rawDescGZIP(), []int{0}
}

func (x *StateObjectsList) GetStateObjects() []string {
	if x != nil {
		return x.StateObjects
	}
	return nil
}

// StateTransitionsList provides the list of state transitions
// that can be delivered to the runtime represented as protobuf.Fullname(meta.StateTransition)
type StateTransitionsList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StateTransitions []string `protobuf:"bytes,1,rep,name=state_transitions,json=stateTransitions,proto3" json:"state_transitions,omitempty"`
}

func (x *StateTransitionsList) Reset() {
	*x = StateTransitionsList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_module_runtime_v1alpha1_runtime_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StateTransitionsList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StateTransitionsList) ProtoMessage() {}

func (x *StateTransitionsList) ProtoReflect() protoreflect.Message {
	mi := &file_module_runtime_v1alpha1_runtime_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StateTransitionsList.ProtoReflect.Descriptor instead.
func (*StateTransitionsList) Descriptor() ([]byte, []int) {
	return file_module_runtime_v1alpha1_runtime_proto_rawDescGZIP(), []int{1}
}

func (x *StateTransitionsList) GetStateTransitions() []string {
	if x != nil {
		return x.StateTransitions
	}
	return nil
}

type CreateStateObjectsList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StateObjects []string `protobuf:"bytes,1,rep,name=state_objects,json=stateObjects,proto3" json:"state_objects,omitempty"`
}

func (x *CreateStateObjectsList) Reset() {
	*x = CreateStateObjectsList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_module_runtime_v1alpha1_runtime_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateStateObjectsList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateStateObjectsList) ProtoMessage() {}

func (x *CreateStateObjectsList) ProtoReflect() protoreflect.Message {
	mi := &file_module_runtime_v1alpha1_runtime_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateStateObjectsList.ProtoReflect.Descriptor instead.
func (*CreateStateObjectsList) Descriptor() ([]byte, []int) {
	return file_module_runtime_v1alpha1_runtime_proto_rawDescGZIP(), []int{2}
}

func (x *CreateStateObjectsList) GetStateObjects() []string {
	if x != nil {
		return x.StateObjects
	}
	return nil
}

type CreateStateTransitionsList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StateTransitions []string `protobuf:"bytes,1,rep,name=state_transitions,json=stateTransitions,proto3" json:"state_transitions,omitempty"`
}

func (x *CreateStateTransitionsList) Reset() {
	*x = CreateStateTransitionsList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_module_runtime_v1alpha1_runtime_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateStateTransitionsList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateStateTransitionsList) ProtoMessage() {}

func (x *CreateStateTransitionsList) ProtoReflect() protoreflect.Message {
	mi := &file_module_runtime_v1alpha1_runtime_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateStateTransitionsList.ProtoReflect.Descriptor instead.
func (*CreateStateTransitionsList) Descriptor() ([]byte, []int) {
	return file_module_runtime_v1alpha1_runtime_proto_rawDescGZIP(), []int{3}
}

func (x *CreateStateTransitionsList) GetStateTransitions() []string {
	if x != nil {
		return x.StateTransitions
	}
	return nil
}

var File_module_runtime_v1alpha1_runtime_proto protoreflect.FileDescriptor

var file_module_runtime_v1alpha1_runtime_proto_rawDesc = []byte{
	0x0a, 0x25, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2f, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65,
	0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x74, 0x6d, 0x6f, 0x73, 0x2e, 0x72, 0x75,
	0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x22, 0x37,
	0x0a, 0x10, 0x53, 0x74, 0x61, 0x74, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x6f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x22, 0x43, 0x0a, 0x14, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x2b, 0x0a, 0x11, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x10, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x3d, 0x0a, 0x16,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f,
	0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x22, 0x49, 0x0a, 0x1a, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x11, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x10, 0x73, 0x74, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2a, 0x57, 0x0a, 0x04, 0x56, 0x65, 0x72, 0x62, 0x12, 0x0b,
	0x0a, 0x07, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x47,
	0x65, 0x74, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x10, 0x02, 0x12, 0x0a,
	0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x10, 0x03, 0x12, 0x0a, 0x0a, 0x06, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x10, 0x04, 0x12, 0x0a, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x10, 0x05, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x10, 0x06, 0x42,
	0x32, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x64,
	0x79, 0x6d, 0x79, 0x6c, 0x6a, 0x61, 0x2f, 0x74, 0x6d, 0x6f, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x75,
	0x6c, 0x65, 0x2f, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_module_runtime_v1alpha1_runtime_proto_rawDescOnce sync.Once
	file_module_runtime_v1alpha1_runtime_proto_rawDescData = file_module_runtime_v1alpha1_runtime_proto_rawDesc
)

func file_module_runtime_v1alpha1_runtime_proto_rawDescGZIP() []byte {
	file_module_runtime_v1alpha1_runtime_proto_rawDescOnce.Do(func() {
		file_module_runtime_v1alpha1_runtime_proto_rawDescData = protoimpl.X.CompressGZIP(file_module_runtime_v1alpha1_runtime_proto_rawDescData)
	})
	return file_module_runtime_v1alpha1_runtime_proto_rawDescData
}

var file_module_runtime_v1alpha1_runtime_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_module_runtime_v1alpha1_runtime_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_module_runtime_v1alpha1_runtime_proto_goTypes = []interface{}{
	(Verb)(0),                          // 0: tmos.runtime.v1alpha1.Verb
	(*StateObjectsList)(nil),           // 1: tmos.runtime.v1alpha1.StateObjectsList
	(*StateTransitionsList)(nil),       // 2: tmos.runtime.v1alpha1.StateTransitionsList
	(*CreateStateObjectsList)(nil),     // 3: tmos.runtime.v1alpha1.CreateStateObjectsList
	(*CreateStateTransitionsList)(nil), // 4: tmos.runtime.v1alpha1.CreateStateTransitionsList
}
var file_module_runtime_v1alpha1_runtime_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_module_runtime_v1alpha1_runtime_proto_init() }
func file_module_runtime_v1alpha1_runtime_proto_init() {
	if File_module_runtime_v1alpha1_runtime_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_module_runtime_v1alpha1_runtime_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StateObjectsList); i {
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
		file_module_runtime_v1alpha1_runtime_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StateTransitionsList); i {
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
		file_module_runtime_v1alpha1_runtime_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateStateObjectsList); i {
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
		file_module_runtime_v1alpha1_runtime_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateStateTransitionsList); i {
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
			RawDescriptor: file_module_runtime_v1alpha1_runtime_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_module_runtime_v1alpha1_runtime_proto_goTypes,
		DependencyIndexes: file_module_runtime_v1alpha1_runtime_proto_depIdxs,
		EnumInfos:         file_module_runtime_v1alpha1_runtime_proto_enumTypes,
		MessageInfos:      file_module_runtime_v1alpha1_runtime_proto_msgTypes,
	}.Build()
	File_module_runtime_v1alpha1_runtime_proto = out.File
	file_module_runtime_v1alpha1_runtime_proto_rawDesc = nil
	file_module_runtime_v1alpha1_runtime_proto_goTypes = nil
	file_module_runtime_v1alpha1_runtime_proto_depIdxs = nil
}
