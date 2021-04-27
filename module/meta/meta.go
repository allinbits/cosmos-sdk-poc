package meta

import (
	"github.com/fdymylja/tmos/module/meta/v1alpha1"
	"google.golang.org/protobuf/proto"
)

// Type defines a generalized type that can be fed to the runtime
type Type interface {
	proto.Message
}

// StateTransition is a type which is used to cause state transitions
type StateTransition interface {
	Type
	GetTransitionMeta() *v1alpha1.TransitionMeta
}

// Name returns the unique name for the Type
func Name(t Type) string {
	return (string)(t.ProtoReflect().Descriptor().FullName())
}
