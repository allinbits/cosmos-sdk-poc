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

// StateObject defines an object which is saved in the state
type StateObject interface {
	Type
	GetID() ID
}

// ID defines the unique identification of an StateObject.
type ID interface {
	Bytes()
}

// Name returns the unique name for the Type
func Name(t Type) string {
	return (string)(t.ProtoReflect().Descriptor().FullName())
}
