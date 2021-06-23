package runtime

import (
	"fmt"

	"github.com/fdymylja/tmos/core/meta"
	"github.com/fdymylja/tmos/runtime/authentication/user"
	"github.com/fdymylja/tmos/runtime/client"
)

// NewModuleClient instantiates a RuntimeClient that can only be used by modules.
func NewModuleClient(runtime *Runtime) *ModuleClient {
	return &ModuleClient{
		runtime: runtime,
	}
}

// ModuleClient is the runtime.Runtime client used internally by modules to interact with each other and the store.
type ModuleClient struct {
	users   user.Users
	runtime *Runtime
}

func (c *ModuleClient) Get(id meta.ID, object meta.StateObject, opts ...client.GetOption) error {
	o := new(client.GetOptionsRaw)
	for _, opt := range opts {
		opt(o)
	}
	if o.Height != 0 {
		return fmt.Errorf("runtime.Client: heighted requests are not allowed")
	}
	return c.runtime.Get(id, object)
}

func (c *ModuleClient) Create(object meta.StateObject, _ ...client.CreateOption) error {
	return c.runtime.Create(c.users, object)
}

func (c *ModuleClient) Update(object meta.StateObject, _ ...client.UpdateOption) error {
	return c.runtime.Update(c.users, object)
}

func (c *ModuleClient) Delete(object meta.StateObject, _ ...client.DeleteOption) error {
	return c.runtime.Delete(c.users, object)
}

func (c *ModuleClient) Deliver(transition meta.StateTransition, opts ...client.DeliverOption) error {
	o := new(client.RawDeliverOptions)
	// apply options
	for _, opt := range opts {
		opt(o)
	}
	// impersonate users if defined
	users := user.NewUsersUnion(c.users, o.Impersonate)

	return c.runtime.Deliver(users, transition)
}

func (c *ModuleClient) List(object meta.StateObject, opts ...client.ListOption) (client.ObjectIterator, error) {
	opt := new(client.ListOptionsRaw)
	for _, o := range opts {
		o(opt)
	}
	// assert height is zero
	if opt.Height != 0 {
		return client.ObjectIterator{}, fmt.Errorf("runtime.Client: heighted list is not allowed for module clients")
	}
	return c.runtime.List(object, opt.ORMOptions)
}

func (c *ModuleClient) SetUser(u user.Users) {
	c.users = u
}
