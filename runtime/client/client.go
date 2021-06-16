package client

import (
	"github.com/fdymylja/tmos/core/meta"
	"github.com/fdymylja/tmos/runtime/authentication/user"
)

// RuntimeServer represents runtime.Runtime as a server
// which other module.Client can interact with.
type RuntimeServer interface {
	Get(id meta.ID, object meta.StateObject) error
	Create(users user.Users, object meta.StateObject) error
	Update(users user.Users, object meta.StateObject) error
	Delete(users user.Users, object meta.StateObject) error
	Deliver(users user.Users, transition meta.StateTransition) error
}

// RuntimeClient is the client that modules use to interact
// with the RuntimeServer and talk with the store and other modules.
type RuntimeClient interface {
	Get(id meta.ID, object meta.StateObject, opts ...GetOption) error
	Create(object meta.StateObject, opts ...CreateOption) error
	Update(object meta.StateObject, opts ...UpdateOption) error
	Delete(object meta.StateObject, opts ...DeleteOption) error
	Deliver(transition meta.StateTransition, opts ...DeliverOption) error
}
