package v1alpha1

import "github.com/fdymylja/tmos/runtime/module"

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
