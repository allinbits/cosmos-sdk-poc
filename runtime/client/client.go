package client

import (
	"github.com/fdymylja/tmos/runtime/authentication"
	"github.com/fdymylja/tmos/runtime/meta"
)

// RuntimeServer represents runtime.Runtime as a server
// which other module.Client can interact with.
type RuntimeServer interface {
	Get(id meta.ID, object meta.StateObject) error
	Create(subject string, object meta.StateObject) error
	Update(subject string, object meta.StateObject) error
	Delete(subject string, object meta.StateObject) error
	Deliver(subjects *authentication.Subjects, transition meta.StateTransition) error
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
	subject string // subject carries identity info regarding the module
	runtime RuntimeServer
}

func (c *client) Get(id meta.ID, object meta.StateObject, opts ...GetOption) error {
	return c.runtime.Get(id, object)
}

func (c *client) Create(object meta.StateObject, opts ...CreateOption) error {
	return c.runtime.Create(c.subject, object)
}

func (c *client) Update(object meta.StateObject, opts ...UpdateOption) error {
	return c.runtime.Update(c.subject, object)
}

func (c *client) Delete(object meta.StateObject, opts ...DeleteOption) error {
	return c.runtime.Delete(c.subject, object)
}

func (c *client) Deliver(transition meta.StateTransition, opts ...DeliverOption) error {
	o := new(deliverOptions)
	for _, opt := range opts {
		opt(o)
	}
	subjects := authentication.NewEmptySubjects()
	subjects.Add(c.subject)
	// add impersonation subjects
	for _, impersonating := range o.impersonate {
		subjects.Add(impersonating)
	}
	return c.runtime.Deliver(subjects, transition)
}

func (c *client) SetSubject(subject string) {
	c.subject = subject
}
