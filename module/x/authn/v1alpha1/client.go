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

func (c *Client) GetParams() (*Params, error) {
	p := new(Params)
	err := c.c.Get(ParamsID, p)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (c *Client) GetAccount(address string) (*Account, error) {
	a := new(Account)
	err := c.c.Get(meta.NewStringID(address), a)
	return a, err
}
