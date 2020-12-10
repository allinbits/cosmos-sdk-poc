package client

import (
	"github.com/fdymylja/cosmos-os/internal/runtime"
	"github.com/fdymylja/cosmos-os/pkg/apis/core/v1alpha1"
	"github.com/fdymylja/cosmos-os/pkg/codec"
	"github.com/tendermint/tendermint/abci/types"
	"google.golang.org/protobuf/types/known/anypb"
)

type Client struct {
	rt runtime.Runtime
}

func NewClient(rt runtime.Runtime) Client {
	return Client{
		rt: rt,
	}
}

func (c Client) Tx(msg codec.Object) error {
	anyMsg, err := codec.NewTx(msg)
	if err != nil {
		return err
	}
	tx := &v1alpha1.Transaction{
		Auth:     nil,
		Messages: anyMsg,
	}
	txBytes, err := codec.Marshal(tx)
	resp := c.rt.DeliverTx(types.RequestDeliverTx{Tx: txBytes})
	if resp.Code != 0 {
		panic("duh")
	}
	return nil
}

func (c Client) Query(height int64, query codec.Object, response codec.Object) error {
	anyQuery, err := anypb.New(query)
	if err != nil {
		return err
	}
	queryObject := &v1alpha1.Query{Query: anyQuery}

	queryBytes, err := codec.Marshal(queryObject)
	if err != nil {
		return err
	}

	resp := c.rt.Query(types.RequestQuery{
		Data:   queryBytes,
		Path:   "",
		Height: height,
		Prove:  false,
	})

	if resp.Code != 0 {
		panic("duh")
	}

	return codec.Unmarshal(resp.Value, response)
}
