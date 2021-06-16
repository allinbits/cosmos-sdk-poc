package client

import (
	"fmt"

	"github.com/fdymylja/tmos/core/meta"
	"github.com/fdymylja/tmos/runtime/authentication/user"
)

var _ RuntimeClient = (*ModuleClient)(nil)

// NewModuleClient instantiates a RuntimeClient that can only be used by modules.
func NewModuleClient(runtime RuntimeServer) RuntimeClient {
	return &ModuleClient{
		runtime: runtime,
	}
}

// ModuleClient is the runtime.Runtime client used internally by modules to interact with each other and the store.
type ModuleClient struct {
	users   user.Users
	runtime RuntimeServer
}

func (c *ModuleClient) Get(id meta.ID, object meta.StateObject, opts ...GetOption) error {
	return c.runtime.Get(id, object)
}

func (c *ModuleClient) Create(object meta.StateObject, opts ...CreateOption) error {
	return c.runtime.Create(c.users, object)
}

func (c *ModuleClient) Update(object meta.StateObject, opts ...UpdateOption) error {
	return c.runtime.Update(c.users, object)
}

func (c *ModuleClient) Delete(object meta.StateObject, opts ...DeleteOption) error {
	return c.runtime.Delete(c.users, object)
}

func (c *ModuleClient) Deliver(transition meta.StateTransition, opts ...DeliverOption) error {
	o := new(deliverOptions)
	// apply options
	for _, opt := range opts {
		opt(o)
	}
	// impersonate users if defined
	users := user.NewUsersUnion(c.users, o.impersonate)

	return c.runtime.Deliver(users, transition)
}

func (c *ModuleClient) List(object meta.StateObject, opts ...ListOption) (ObjectIterator, error) {
	opt := new(listOptions)
	for _, o := range opts {
		o(opt)
	}
	// assert height is zero
	if opt.Height != 0 {
		return ObjectIterator{}, fmt.Errorf("client: heighted list is not allowed for module clients")
	}
	return c.runtime.List(object, opt.ORMOptions)
}

func (c *ModuleClient) SetUser(u user.Users) {
	c.users = u
}
