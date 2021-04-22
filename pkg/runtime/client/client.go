package client

import (
	"github.com/fdymylja/tmos/apis/meta"
)

// Runtime defines runtime functionalities needed by clients
type Runtime interface {
	// Deliver delivers a meta.StateTransition to the handling controller
	Deliver(identities []string, transition meta.StateTransition) error
}

// Store defines the store
type Store interface {
	Get(object meta.StateObject) bool
	Set(user string, object meta.StateObject) error
	Delete(user string, object meta.StateObject) error
}

func NewClient(identity string, store Store, runtime Runtime) Client {
	return Client{
		user:    identity,
		store:   store,
		runtime: runtime,
	}
}

type Client struct {
	user    string
	store   Store
	runtime Runtime
}

func (c Client) Get(object meta.StateObject) (exists bool) {
	return c.store.Get(object)
}

func (c Client) Set(object meta.StateObject) {
	err := c.store.Set(c.user, object)
	if err != nil {
		panic(err)
	}
}

func (c Client) Delete(object meta.StateObject) {
	err := c.store.Delete(c.user, object)
	if err != nil {
		panic(err)
	}
}

func (c Client) Deliver(transition meta.StateTransition) error {
	return c.runtime.Deliver([]string{c.user}, transition)
}
