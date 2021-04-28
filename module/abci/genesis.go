package abci

import (
	"encoding/json"

	"github.com/fdymylja/tmos/module/abci/v1alpha1"
	"github.com/fdymylja/tmos/runtime"
	"github.com/fdymylja/tmos/runtime/controller"
)

func newGenesisHandler(client runtime.ModuleClient) controller.Genesis {
	return genesis{client}
}

type genesis struct {
	client runtime.ModuleClient
}

func (g genesis) SetDefault() error {
	// set initial stage
	err := g.client.Create(&v1alpha1.Stage{Stage: v1alpha1.ABCIStage_InitChain})
	if err != nil {
		return err
	}
	// set empty begin block state
	err = g.client.Create(&v1alpha1.BeginBlockState{})
	if err != nil {
		return err
	}
	// set empty check tx state
	err = g.client.Create(&v1alpha1.CheckTxState{})
	if err != nil {
		return err
	}
	// set empty deliver tx state
	err = g.client.Create(&v1alpha1.DeliverTxState{})
	if err != nil {
		return err
	}
	return nil
}

func (g genesis) Import(state json.RawMessage) error {
	panic("implement me")
}

func (g genesis) Export(state json.RawMessage) error {
	panic("implement me")
}
