package runtime

import (
	"google.golang.org/protobuf/proto"
)

// Type defines a generalized type that can be fed to the runtime
type Type interface {
	proto.Message
}

// StateTransition is a type which is used to cause state transitions
type StateTransition interface {
	Type
}

// Name returns the unique name for the Type
func Name(t Type) string {
	return (string)(t.ProtoReflect().Descriptor().FullName())
}
