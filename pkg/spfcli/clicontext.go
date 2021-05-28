package spfcli

import (
	"context"
	"sync"

	runtimev1alpha1 "github.com/fdymylja/tmos/core/runtime/v1alpha1"
	"github.com/fdymylja/tmos/pkg/client"
	"github.com/fdymylja/tmos/runtime/meta"
)

type CLIContext interface {
	Client() client.Client
	Sign()
	Context() context.Context
}

func NewCLIContext() CLIContext {
	return &cliContext{}
}

type cliContext struct {
	getConfigOnce *sync.Once
	client        client.Client
}

func (c *cliContext) init() {
	conf, err := GetConfig()
	if err != nil {
		c.client = newFailedClient(err)
		return
	}

	rpc, err := client.NewRPC(conf.TendermintRPC)
	if err != nil {
		c.client = newFailedClient(err)
		return
	}

	c.client = rpc
}

func (c *cliContext) Client() client.Client {
	c.getConfigOnce.Do(c.init)
	return c.client
}

func (c *cliContext) Sign() {
	panic("implement me")
}

func (c *cliContext) Context() context.Context {
	return context.TODO()
}

var _ client.Client = failedClient{}

func newFailedClient(err error) client.Client {
	return failedClient{err: err}
}

// failedClient is used when there were some precondition issues
// to report the error when the client is being used
type failedClient struct {
	err error
}

func (f failedClient) Resources(_ context.Context) (*runtimev1alpha1.Resources, error) {
	return nil, f.err
}

func (f failedClient) Get(_ context.Context, _ meta.APIGroup, _ meta.APIKind, _ string) ([]byte, error) {
	return nil, f.err
}
