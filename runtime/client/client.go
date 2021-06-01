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

var _ RuntimeClient = (*client)(nil)

func New(runtime RuntimeServer) RuntimeClient {
	return &client{
		runtime: runtime,
	}
}

// client is the runtime.Runtime client used internally by modules to interact with each other and the store.
type client struct {
	users   user.Users
	runtime RuntimeServer
}

func (c *client) Get(id meta.ID, object meta.StateObject, opts ...GetOption) error {
	return c.runtime.Get(id, object)
}

func (c *client) Create(object meta.StateObject, opts ...CreateOption) error {
	return c.runtime.Create(c.users, object)
}

func (c *client) Update(object meta.StateObject, opts ...UpdateOption) error {
	return c.runtime.Update(c.users, object)
}

func (c *client) Delete(object meta.StateObject, opts ...DeleteOption) error {
	return c.runtime.Delete(c.users, object)
}

func (c *client) Deliver(transition meta.StateTransition, opts ...DeliverOption) error {
	o := new(deliverOptions)
	// apply options
	for _, opt := range opts {
		opt(o)
	}
	// impersonate users if defined
	users := user.NewUsersUnion(c.users, o.impersonate)

	return c.runtime.Deliver(users, transition)
}

func (c *client) SetUser(u user.Users) {
	c.users = u
}
