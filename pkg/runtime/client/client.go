package client

import (
	"github.com/fdymylja/tmos/apis/meta"
	"github.com/fdymylja/tmos/pkg/module"
)

// RuntimeServer defines runtime functionalities needed by clients
type RuntimeServer interface {
	Get(object meta.StateObject) error
	List() // TBD
	Create(user string, object meta.StateObject) error
	Update(user string, object meta.StateObject) error
	Delete(user string, object meta.StateObject) error
	// Deliver delivers a meta.StateTransition to the handling controller
	Deliver(identities []string, transition meta.StateTransition) error
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

func (c *Client) Get(object meta.StateObject) error {
	return c.runtime.Get(object)
}

func (c *Client) Create(object meta.StateObject) error {
	return c.runtime.Create(c.user, object)
}

func (c *Client) Update(object meta.StateObject) error {
	return c.runtime.Update(c.user, object)
}

func (c *Client) Delete(object meta.StateObject) error {
	return c.runtime.Delete(c.user, object)
}

func (c Client) Deliver(transition meta.StateTransition) error {
	return c.runtime.Deliver([]string{c.user}, transition)
}

func (c *Client) SetUser(user string) {
	c.user = user
}
