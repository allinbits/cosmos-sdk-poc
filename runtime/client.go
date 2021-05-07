package runtime

import (
	"github.com/fdymylja/tmos/runtime/authentication"
	"github.com/fdymylja/tmos/runtime/meta"
	"github.com/fdymylja/tmos/runtime/module"
)

// server defines runtime functionalities needed by clients
type server interface {
	Get(id meta.ID, object meta.StateObject) error
	List() // TBD
	Create(user string, object meta.StateObject) error
	Update(user string, object meta.StateObject) error
	Delete(user string, object meta.StateObject) error
	// Deliver delivers a meta.StateTransition to the handling controller
	Deliver(subjects *authentication.Subjects, transition meta.StateTransition, opts ...DeliverOption) error
}

var _ module.Client = (*client)(nil)

func newClient(runtime server) *client {
	return &client{
		runtime: runtime,
	}
}

type client struct {
	user    string
	runtime server
}

func (c *client) Get(id meta.ID, object meta.StateObject) error {
	return c.runtime.Get(id, object)
}

func (c *client) Create(object meta.StateObject) error {
	return c.runtime.Create(c.user, object)
}

func (c *client) Update(object meta.StateObject) error {
	return c.runtime.Update(c.user, object)
}

func (c *client) Delete(object meta.StateObject) error {
	return c.runtime.Delete(c.user, object)
}

func (c client) Deliver(transition meta.StateTransition) error {
	subjects := authentication.NewEmptySubjects()
	subjects.Add(c.user)
	return c.runtime.Deliver(subjects, transition)
}

func (c client) List(obj meta.StateObject) error { panic("not impl") }

func (c *client) SetUser(user string) {
	c.user = user
}
