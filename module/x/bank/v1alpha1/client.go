package v1alpha1

import (
	"github.com/fdymylja/tmos/runtime/meta"
	"github.com/fdymylja/tmos/runtime/module"
)

type Client struct {
	c module.Client
}

func (c *Client) GetBalance(address string) (*Balance, error) {
	balance := new(Balance)
	err := c.c.Get(meta.NewStringID(address), balance)
	return balance, err
}
