package client

import (
	"context"
	"fmt"

	runtimev1alpha1 "github.com/fdymylja/tmos/core/runtime/v1alpha1"
	"github.com/fdymylja/tmos/runtime"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/rpc/client/http"
	"google.golang.org/protobuf/encoding/protojson"
)

type Client interface {
	Resources(ctx context.Context) (*runtimev1alpha1.Resources, error)
	Get(ctx context.Context, apiGroup string, apiKind string, name string) ([]byte, error)
}

func NewRPC(endpoint string) (Client, error) {
	rpc, err := http.New(endpoint, "")
	if err != nil {
		return nil, err
	}
	return &client{rpc: rpc}, nil
}

type client struct {
	rpc *http.HTTP
}

func (c *client) Resources(ctx context.Context) (*runtimev1alpha1.Resources, error) {
	resp, err := c.rpc.ABCIQuery(ctx, runtime.ResourcesPath, nil)
	if err != nil {
		return nil, err
	}
	if resp.Response.Code != abci.CodeTypeOK {
		return nil, fmt.Errorf("client: %d %s", resp.Response.Code, resp.Response.Log)
	}

	res := &runtimev1alpha1.Resources{}
	err = protojson.Unmarshal(resp.Response.Value, res)
	if err != nil {
		return nil, fmt.Errorf("client: codec error: %w", err)
	}

	return res, nil
}

func (c *client) Get(ctx context.Context, apiGroup string, apiKind string, name string) ([]byte, error) {
	path := fmt.Sprintf("%s/%s/%s/%s", runtime.GetStateObjectsPath, apiGroup, apiKind, name)
	resp, err := c.rpc.ABCIQuery(ctx, path, nil)
	if err != nil {
		return nil, err
	}
	if resp.Response.Code != abci.CodeTypeOK {
		return nil, fmt.Errorf("client: %d %s", resp.Response.Code, resp.Response.Log)
	}

	return resp.Response.Value, nil
}
