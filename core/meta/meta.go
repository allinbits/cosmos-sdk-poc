package meta

import (
	"fmt"

	"google.golang.org/protobuf/proto"
)

func (x *APIDefinition) Name() string {
	return fmt.Sprintf("%s.%s", x.Group, x.Kind)
}

// APIObject defines a golang object that can be processed and understood by the runtime
type APIObject interface {
	proto.Message
	// APIDefinition describes the API
	APIDefinition() *APIDefinition
}

// StateObject defines an APIObject that is meant to be saved into state
type StateObject interface {
	APIObject
	NewStateObject() StateObject
}

// StateTransition defines an APIObject that is used to change state
type StateTransition interface {
	APIObject
	NewStateTransition() StateTransition
}

// Name returns the unique name for the Type
func Name(t APIObject) string {
	return t.APIDefinition().Name()
}
