package v1alpha1

import (
	"fmt"

	"github.com/fdymylja/tmos/module/x/authn/crypto"
	"github.com/fdymylja/tmos/runtime/meta"
	"github.com/fdymylja/tmos/runtime/module"
	gogoproto "github.com/gogo/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
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

func (c *Client) UpdatePublicKey(address string, newPubKey crypto.PubKey) error {
	acc, err := c.GetAccount(address)
	if err != nil {
		return err
	}
	// we do some hacks here, because we're using old sdk types which still rely
	// on gogoproto... but at the same time we're working with protov2 API here
	// so until we migrate all the types to protov2 this is what we do..
	pk, err := gogoproto.Marshal(newPubKey)
	if err != nil {
		return err
	}
	acc.PubKey = &anypb.Any{
		TypeUrl: fmt.Sprintf("/%s", gogoproto.MessageName(newPubKey)),
		Value:   pk,
	}
	return c.c.Update(acc)
}
