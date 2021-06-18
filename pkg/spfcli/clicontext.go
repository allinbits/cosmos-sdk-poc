package spfcli

import (
	"context"
	"io"
	"sync"

	runtimev1alpha1 "github.com/fdymylja/tmos/core/runtime/v1alpha1"
	"github.com/fdymylja/tmos/pkg/client"
)

type CLIContext interface {
	Client() client.Client
	Sign()
	Context() context.Context
}

func NewCLIContext(input io.Reader) CLIContext {
	return &cliContext{
		userInput:     input,
		getConfigOnce: new(sync.Once),
		client:        nil,
	}
}

type cliContext struct {
	getConfigOnce *sync.Once
	client        client.Client
	userInput     io.Reader
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

func (f failedClient) Get(_ context.Context, _ string, _ string, _ string) ([]byte, error) {
	return nil, f.err
}
