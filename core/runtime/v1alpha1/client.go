package v1alpha1

import (
	"github.com/fdymylja/tmos/runtime/meta"
	"github.com/fdymylja/tmos/runtime/module"
)

func NewClient(c module.Client) *Client {
	return &Client{c: c}
}

type Client struct {
	c module.Client
}

func (c *Client) GetStateObjectsList() (*StateObjectsList, error) {
	list := new(StateObjectsList)
	err := c.c.Get(meta.SingletonID, list)
	return list, err
}

func (c *Client) GetStateTransitionsList() (*StateTransitionsList, error) {
	list := new(StateTransitionsList)
	err := c.c.Get(meta.SingletonID, list)
	return list, err
}
