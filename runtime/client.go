package runtime

import (
	"github.com/fdymylja/tmos/runtime/module"
)

// RuntimeServer defines runtime functionalities needed by clients
type RuntimeServer interface {
	Get(object StateObject) error
	List() // TBD
	Create(user string, object StateObject) error
	Update(user string, object StateObject) error
	Delete(user string, object StateObject) error
	// Deliver delivers a meta.StateTransition to the handling controller
	Deliver(identities []string, transition StateTransition, skipAdmissionControllers bool) error
}

var _ module.Client = (*Client)(nil)

func NewClient(runtime RuntimeServer) *Client {
	return &Client{
		runtime: runtime,
	}
}

type Client struct {
	user    string
	runtime RuntimeServer
}

func (c *Client) Get(object StateObject) error {
	return c.runtime.Get(object)
}

func (c *Client) Create(object StateObject) error {
	return c.runtime.Create(c.user, object)
}

func (c *Client) Update(object StateObject) error {
	return c.runtime.Update(c.user, object)
}

func (c *Client) Delete(object StateObject) error {
	return c.runtime.Delete(c.user, object)
}

func (c Client) Deliver(transition StateTransition) error {
	return c.runtime.Deliver([]string{c.user}, transition, false)
}

func (c *Client) SetUser(user string) {
	c.user = user
}
