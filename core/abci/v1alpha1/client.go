package v1alpha1

import "github.com/fdymylja/tmos/runtime/module"

func NewClient(c module.Client) *Client {
	return &Client{c: c}
}

type Client struct {
	c module.Client
}

func (c *Client) SetBeginBlock(block *BeginBlockState) error {
	return c.c.Create(block)
}

func (c *Client) GetBeginBlock() (*BeginBlockState, error) {
	state := new(BeginBlockState)
	err := c.c.Get(nil, state)
	if err != nil {
		return nil, err
	}
	return state, nil
}

func (c *Client) GetCurrentBlock() (*CurrentBlock, error) {
	currentBlock := new(CurrentBlock)
	err := c.c.Get(nil, currentBlock)
	return currentBlock, err
}

func (c *Client) GetChainID() (string, error) {
	chainInfo := new(InitChainInfo)
	err := c.c.Get(nil, chainInfo)
	if err != nil {
		return "", err
	}
	return chainInfo.ChainId, nil
}
