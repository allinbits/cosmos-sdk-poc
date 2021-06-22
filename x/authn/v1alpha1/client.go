package v1alpha1

import (
	"fmt"

	"google.golang.org/protobuf/proto"

	"github.com/fdymylja/tmos/core/meta"
	"github.com/fdymylja/tmos/runtime/module"
	crypto2 "github.com/fdymylja/tmos/x/authn/crypto"
	"google.golang.org/protobuf/types/known/anypb"
)

func NewClient(c module.Client) *Client {
	return &Client{c: c}
}

type Client struct {
	c module.Client
}

func (c *Client) IncreaseSequence(address string) error {
	acc := new(Account)
	err := c.c.Get(meta.NewStringID(address), acc)
	if err != nil {
		return err
	}
	acc.Sequence = acc.Sequence + 1
	return c.c.Update(acc)
}

func (c *Client) GetParams() (*Params, error) {
	p := new(Params)
	err := c.c.Get(meta.SingletonID, p)
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

func (c *Client) CreateAccount(acc *Account) error {
	return c.c.Deliver(&MsgCreateAccount{
		Account: acc,
	})
}

func (c *Client) UpdatePublicKey(address string, newPubKey crypto2.PubKey) error {
	acc, err := c.GetAccount(address)
	if err != nil {
		return err
	}
	// we do some hacks here, because we're using old sdk types which still rely
	// on gogoproto... but at the same time we're working with protov2 API here
	// so until we migrate all the types to protov2 this is what we do..
	pk, err := proto.Marshal(newPubKey)
	if err != nil {
		return err
	}
	acc.PubKey = &anypb.Any{
		TypeUrl: fmt.Sprintf("/%s", proto.MessageName(newPubKey)),
		Value:   pk,
	}
	return c.c.Update(acc)
}
