package testpb

import (
	meta "github.com/fdymylja/tmos/core/meta"
)

func (x *SimpleMessage) NewStateObject() meta.StateObject {
	return new(SimpleMessage)
}

func (x *SimpleMessage) APIDefinition() *meta.APIDefinition {
	return &meta.APIDefinition{Group: "", Kind: ""}
}
