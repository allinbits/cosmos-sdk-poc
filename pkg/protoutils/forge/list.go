package forge

import (
	"fmt"

	"github.com/fdymylja/tmos/pkg/ptrprim"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/dynamicpb"
)

const PackagePrefix = "dynamic"

// List forges a dynamic protobuf object which represents a list of the given message
// no changes to the type registry or file registry are made. This type should be only
// used for marshalling and unmarshalling to and from proto or json bytes.
// It shouldn't be used for reflection or introspection as results are not guaranteed to be consistent.
func List(message proto.Message, resolver protodesc.Resolver) (protoreflect.MessageType, error) {
	fd := message.ProtoReflect().Descriptor().ParentFile()
	md := message.ProtoReflect().Descriptor()
	forgedFdp := &descriptorpb.FileDescriptorProto{
		Name:       ptrprim.String(fmt.Sprintf("%s/%s", PackagePrefix, fd.Path())),
		Package:    ptrprim.String(fmt.Sprintf("%s.%s", PackagePrefix, fd.Package())),
		Dependency: []string{fd.Path()},
		MessageType: []*descriptorpb.DescriptorProto{
			{
				Name: ptrprim.String(fmt.Sprintf("%sList", md.Name())),
				Field: []*descriptorpb.FieldDescriptorProto{
					{
						Name:   ptrprim.String(fmt.Sprintf("list")),
						Number: ptrprim.Int32(1),
						Label: func() *descriptorpb.FieldDescriptorProto_Label {
							x := descriptorpb.FieldDescriptorProto_LABEL_REPEATED
							return &x
						}(),
						Type: func() *descriptorpb.FieldDescriptorProto_Type {
							x := descriptorpb.FieldDescriptorProto_TYPE_MESSAGE
							return &x
						}(),
						TypeName: ptrprim.String(fmt.Sprintf("%s", md.FullName())),
					},
				},
			},
		},
		Options: &descriptorpb.FileOptions{
			GoPackage: ptrprim.String(fmt.Sprintf("%s/dynamic", *fd.Options().(*descriptorpb.FileOptions).GoPackage)), // TODO(fdymylja): remove me when the oas3schema rework is done
		},
	}
	fd, err := protodesc.NewFile(forgedFdp, resolver)
	if err != nil {
		return nil, err
	}

	newMessageMD := fd.Messages().Get(0)
	return dynamicpb.NewMessageType(newMessageMD), nil
}
