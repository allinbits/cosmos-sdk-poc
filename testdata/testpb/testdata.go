package testpb

import "github.com/fdymylja/tmos/runtime/meta"

func (x *SimpleMessage) StateObject() {}
func (x *SimpleMessage) New() meta.StateObject {
	return new(SimpleMessage)
}
