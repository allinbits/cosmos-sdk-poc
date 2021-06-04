package bank

import (
	"encoding/json"

	"github.com/fdymylja/tmos/runtime/module"
)

type genesisHandler struct {
}

func newGenesisHandler() module.GenesisHandler {
	return &genesisHandler{}
}

func (g genesisHandler) Default(module.Client) error {
	return nil
}

func (g genesisHandler) Import(state json.RawMessage) error {
	return nil
}

func (g genesisHandler) Export() (state json.RawMessage, err error) {
	panic("implement me")
}
