package meta

import (
	"google.golang.org/protobuf/proto"
)

func NewStringID(id string) StringID {
	return (StringID)(id)
}

type StringID string

func (s StringID) Bytes() []byte { return []byte(s) }

func NewBytesID(id []byte) BytesID {
	return id
}

type BytesID []byte

func (b BytesID) Bytes() []byte { return b }

// Type defines a generalized type that can be fed to the runtime
type Type interface {
	proto.Message
}

// StateTransition is a type which is used to cause state transitions
type StateTransition interface {
	Type
	StateTransition()
}

// StateObject defines an object which is saved in the state
type StateObject interface {
	Type
	GetID() ID
}

// ID defines the unique identification of an StateObject.
type ID interface {
	Bytes() []byte
}

// Name returns the unique name for the Type
func Name(t Type) string {
	return (string)(t.ProtoReflect().Descriptor().FullName())
}
