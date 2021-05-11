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

func (g genesisHandler) Default() error {
	return nil
}

func (g genesisHandler) Import(state json.RawMessage) error {
	// fmt.Println("%v", state)
	return nil
}

func (g genesisHandler) Export() (state json.RawMessage, err error) {
	panic("implement me")
}
