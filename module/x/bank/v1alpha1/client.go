package v1alpha1

import (
	coin "github.com/fdymylja/tmos/module/core/coin/v1alpha1"
	"github.com/fdymylja/tmos/runtime/meta"
	"github.com/fdymylja/tmos/runtime/module"
)

func NewClient(c module.Client) *Client {
	return &Client{c: c}
}

type Client struct {
	c module.Client
}

func (c *Client) GetBalance(address string) (*Balance, error) {
	balance := new(Balance)
	err := c.c.Get(meta.NewStringID(address), balance)
	return balance, err
}

func (c *Client) Send(sender, recipient string, amount []*coin.Coin) error {
	return c.c.Deliver(&MsgSendCoins{
		FromAddress: sender,
		ToAddress:   recipient,
		Amount:      amount,
	})
}
