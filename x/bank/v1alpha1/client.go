package v1alpha1

import (
	"github.com/fdymylja/tmos/core/coin/v1alpha1"
	"github.com/fdymylja/tmos/runtime/meta"
	"github.com/fdymylja/tmos/runtime/module"
)

func NewClient(c module.Client) *Client {
	return &Client{Client: c}
}

type Client struct {
	module.Client
}

func (c *Client) GetBalance(address string) (*Balance, error) {
	balance := new(Balance)
	err := c.Client.Get(meta.NewStringID(address), balance)
	return balance, err
}

func (c *Client) Send(sender, recipient string, amount []*v1alpha1.Coin) error {
	return c.Client.Deliver(&MsgSendCoins{
		FromAddress: sender,
		ToAddress:   recipient,
		Amount:      amount,
	})
}

func (c *Client) SetBalance(target string, amount []*v1alpha1.Coin) error {
	return c.Client.Deliver(&MsgSetBalance{
		Address: target,
		Amount:  amount,
	})
}
