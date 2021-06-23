// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.13.0
// source: core/rbac/v1alpha1/rbac.proto

package v1alpha1

import (
	_ "github.com/fdymylja/tmos/core/modulegen"
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

type Params struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Params) Reset() {
	*x = Params{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_rbac_v1alpha1_rbac_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Params) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Params) ProtoMessage() {}

func (x *Params) ProtoReflect() protoreflect.Message {
	mi := &file_core_rbac_v1alpha1_rbac_proto_msgTypes[0]
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
	return file_core_rbac_v1alpha1_rbac_proto_rawDescGZIP(), []int{0}
}

// Role defines a role, which defines what resources
// can be accessed and which operations can be performed
// on those resources
type Role struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id is the unique of the role, it does not identify the subject
	// it just identifies the role as multiple subjects can be bound
	// to the same role via RoleBinding
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// gets represents the resources (state objects) which the role can runtime.Get
	Gets []string `protobuf:"bytes,2,rep,name=gets,proto3" json:"gets,omitempty"`
	// lists represents the resources (state objects) which the role can runtime.List
	Lists []string `protobuf:"bytes,3,rep,name=lists,proto3" json:"lists,omitempty"`
	// creates represents the resources (state objects) which the role can runtime.Create
	Creates []string `protobuf:"bytes,4,rep,name=creates,proto3" json:"creates,omitempty"`
	// updates represents the resources (state objects) which the role can runtime.Update
	Updates []string `protobuf:"bytes,5,rep,name=updates,proto3" json:"updates,omitempty"`
	// deletes represents the resources(state objects)  which the role can runtime.Delete
	Deletes []string `protobuf:"bytes,6,rep,name=deletes,proto3" json:"deletes,omitempty"`
	// delivers represents the resources (state transitions) which tre role can runtime.Deliver
	Delivers []string `protobuf:"bytes,7,rep,name=delivers,proto3" json:"delivers,omitempty"`
}

func (x *Role) Reset() {
	*x = Role{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_rbac_v1alpha1_rbac_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Role) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Role) ProtoMessage() {}

func (x *Role) ProtoReflect() protoreflect.Message {
	mi := &file_core_rbac_v1alpha1_rbac_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Role.ProtoReflect.Descriptor instead.
func (*Role) Descriptor() ([]byte, []int) {
	return file_core_rbac_v1alpha1_rbac_proto_rawDescGZIP(), []int{1}
}

func (x *Role) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Role) GetGets() []string {
	if x != nil {
		return x.Gets
	}
	return nil
}

func (x *Role) GetLists() []string {
	if x != nil {
		return x.Lists
	}
	return nil
}

func (x *Role) GetCreates() []string {
	if x != nil {
		return x.Creates
	}
	return nil
}

func (x *Role) GetUpdates() []string {
	if x != nil {
		return x.Updates
	}
	return nil
}

func (x *Role) GetDeletes() []string {
	if x != nil {
		return x.Deletes
	}
	return nil
}

func (x *Role) GetDelivers() []string {
	if x != nil {
		return x.Delivers
	}
	return nil
}

// RoleBinding defines the role for a given subject
type RoleBinding struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// subject defines which account the binding refers to
	Subject string `protobuf:"bytes,1,opt,name=subject,proto3" json:"subject,omitempty"`
	// role_ref points to the Role associated with subject
	RoleRef string `protobuf:"bytes,2,opt,name=role_ref,json=roleRef,proto3" json:"role_ref,omitempty"`
}

func (x *RoleBinding) Reset() {
	*x = RoleBinding{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_rbac_v1alpha1_rbac_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleBinding) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleBinding) ProtoMessage() {}

func (x *RoleBinding) ProtoReflect() protoreflect.Message {
	mi := &file_core_rbac_v1alpha1_rbac_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleBinding.ProtoReflect.Descriptor instead.
func (*RoleBinding) Descriptor() ([]byte, []int) {
	return file_core_rbac_v1alpha1_rbac_proto_rawDescGZIP(), []int{2}
}

func (x *RoleBinding) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *RoleBinding) GetRoleRef() string {
	if x != nil {
		return x.RoleRef
	}
	return ""
}

// MsgCreateRole creates a new role
type MsgCreateRole struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NewRole *Role `protobuf:"bytes,1,opt,name=new_role,json=newRole,proto3" json:"new_role,omitempty"`
}

func (x *MsgCreateRole) Reset() {
	*x = MsgCreateRole{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_rbac_v1alpha1_rbac_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgCreateRole) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgCreateRole) ProtoMessage() {}

func (x *MsgCreateRole) ProtoReflect() protoreflect.Message {
	mi := &file_core_rbac_v1alpha1_rbac_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgCreateRole.ProtoReflect.Descriptor instead.
func (*MsgCreateRole) Descriptor() ([]byte, []int) {
	return file_core_rbac_v1alpha1_rbac_proto_rawDescGZIP(), []int{3}
}

func (x *MsgCreateRole) GetNewRole() *Role {
	if x != nil {
		return x.NewRole
	}
	return nil
}

// MsgBindRole binds subject to role_id
type MsgBindRole struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoleId  string `protobuf:"bytes,1,opt,name=role_id,json=roleId,proto3" json:"role_id,omitempty"`
	Subject string `protobuf:"bytes,2,opt,name=subject,proto3" json:"subject,omitempty"`
}

func (x *MsgBindRole) Reset() {
	*x = MsgBindRole{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_rbac_v1alpha1_rbac_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgBindRole) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgBindRole) ProtoMessage() {}

func (x *MsgBindRole) ProtoReflect() protoreflect.Message {
	mi := &file_core_rbac_v1alpha1_rbac_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgBindRole.ProtoReflect.Descriptor instead.
func (*MsgBindRole) Descriptor() ([]byte, []int) {
	return file_core_rbac_v1alpha1_rbac_proto_rawDescGZIP(), []int{4}
}

func (x *MsgBindRole) GetRoleId() string {
	if x != nil {
		return x.RoleId
	}
	return ""
}

func (x *MsgBindRole) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

var File_core_rbac_v1alpha1_rbac_proto protoreflect.FileDescriptor

var file_core_rbac_v1alpha1_rbac_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x72, 0x62, 0x61, 0x63, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2f, 0x72, 0x62, 0x61, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x17, 0x74, 0x6d, 0x6f, 0x73, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x72, 0x62, 0x61, 0x63, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x1e, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x6d,
	0x6f, 0x64, 0x75, 0x6c, 0x65, 0x67, 0x65, 0x6e, 0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x67,
	0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x72,
	0x62, 0x61, 0x63, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x72, 0x62, 0x61,
	0x63, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x11, 0x0a, 0x06, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x3a, 0x07, 0x92, 0xff, 0x9d,
	0x04, 0x02, 0x08, 0x01, 0x22, 0xb5, 0x01, 0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x67, 0x65, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x67, 0x65, 0x74,
	0x73, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x73, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x05, 0x6c, 0x69, 0x73, 0x74, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x73, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x64,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x64, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72,
	0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72,
	0x73, 0x3a, 0x09, 0x92, 0xff, 0x9d, 0x04, 0x04, 0x12, 0x02, 0x69, 0x64, 0x22, 0x5b, 0x0a, 0x0b,
	0x52, 0x6f, 0x6c, 0x65, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x72, 0x65,
	0x66, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x66,
	0x3a, 0x17, 0x92, 0xff, 0x9d, 0x04, 0x12, 0x12, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x1a, 0x07, 0x72, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x66, 0x22, 0x55, 0x0a, 0x0d, 0x4d, 0x73, 0x67,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x38, 0x0a, 0x08, 0x6e, 0x65,
	0x77, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x74,
	0x6d, 0x6f, 0x73, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x72, 0x62, 0x61, 0x63, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x07, 0x6e, 0x65, 0x77,
	0x52, 0x6f, 0x6c, 0x65, 0x3a, 0x0a, 0x8a, 0xff, 0x9d, 0x04, 0x00, 0x82, 0xb0, 0xe3, 0x2d, 0x00,
	0x22, 0x47, 0x0a, 0x0b, 0x4d, 0x73, 0x67, 0x42, 0x69, 0x6e, 0x64, 0x52, 0x6f, 0x6c, 0x65, 0x12,
	0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x3a, 0x05, 0x8a, 0xff, 0x9d, 0x04, 0x00, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x64, 0x79, 0x6d, 0x79, 0x6c, 0x6a, 0x61,
	0x2f, 0x74, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x72, 0x62, 0x61, 0x63, 0x2f,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_core_rbac_v1alpha1_rbac_proto_rawDescOnce sync.Once
	file_core_rbac_v1alpha1_rbac_proto_rawDescData = file_core_rbac_v1alpha1_rbac_proto_rawDesc
)

func file_core_rbac_v1alpha1_rbac_proto_rawDescGZIP() []byte {
	file_core_rbac_v1alpha1_rbac_proto_rawDescOnce.Do(func() {
		file_core_rbac_v1alpha1_rbac_proto_rawDescData = protoimpl.X.CompressGZIP(file_core_rbac_v1alpha1_rbac_proto_rawDescData)
	})
	return file_core_rbac_v1alpha1_rbac_proto_rawDescData
}

var file_core_rbac_v1alpha1_rbac_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_core_rbac_v1alpha1_rbac_proto_goTypes = []interface{}{
	(*Params)(nil),        // 0: tmos.core.rbac.v1alpha1.Params
	(*Role)(nil),          // 1: tmos.core.rbac.v1alpha1.Role
	(*RoleBinding)(nil),   // 2: tmos.core.rbac.v1alpha1.RoleBinding
	(*MsgCreateRole)(nil), // 3: tmos.core.rbac.v1alpha1.MsgCreateRole
	(*MsgBindRole)(nil),   // 4: tmos.core.rbac.v1alpha1.MsgBindRole
}
var file_core_rbac_v1alpha1_rbac_proto_depIdxs = []int32{
	1, // 0: tmos.core.rbac.v1alpha1.MsgCreateRole.new_role:type_name -> tmos.core.rbac.v1alpha1.Role
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_core_rbac_v1alpha1_rbac_proto_init() }
func file_core_rbac_v1alpha1_rbac_proto_init() {
	if File_core_rbac_v1alpha1_rbac_proto != nil {
		return
	}
	file_core_rbac_v1alpha1_rbac_extension_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_core_rbac_v1alpha1_rbac_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_core_rbac_v1alpha1_rbac_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Role); i {
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
		file_core_rbac_v1alpha1_rbac_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleBinding); i {
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
		file_core_rbac_v1alpha1_rbac_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgCreateRole); i {
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
		file_core_rbac_v1alpha1_rbac_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgBindRole); i {
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
			RawDescriptor: file_core_rbac_v1alpha1_rbac_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_core_rbac_v1alpha1_rbac_proto_goTypes,
		DependencyIndexes: file_core_rbac_v1alpha1_rbac_proto_depIdxs,
		MessageInfos:      file_core_rbac_v1alpha1_rbac_proto_msgTypes,
	}.Build()
	File_core_rbac_v1alpha1_rbac_proto = out.File
	file_core_rbac_v1alpha1_rbac_proto_rawDesc = nil
	file_core_rbac_v1alpha1_rbac_proto_goTypes = nil
	file_core_rbac_v1alpha1_rbac_proto_depIdxs = nil
}
