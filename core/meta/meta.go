package meta

import "google.golang.org/protobuf/proto"

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
