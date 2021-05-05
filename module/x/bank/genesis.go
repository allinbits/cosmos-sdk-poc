package bank

import (
	"encoding/json"

	"github.com/fdymylja/tmos/runtime/controller"
)

type genesisHandler struct {
}

func newGenesisHandler() controller.Genesis {
	return &genesisHandler{}
}

func (g genesisHandler) SetDefault() error {
	return nil
}

func (g genesisHandler) Import(state json.RawMessage) error {
	// fmt.Println("%v", state)
	return nil
}

func (g genesisHandler) Export(state json.RawMessage) error {
	panic("implement me")
}
