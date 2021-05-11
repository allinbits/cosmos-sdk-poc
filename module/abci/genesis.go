package abci

import (
	"encoding/json"

	"github.com/fdymylja/tmos/module/abci/v1alpha1"
	"github.com/fdymylja/tmos/runtime/module"
)

func newGenesisHandler(client module.Client) module.GenesisHandler {
	return genesis{client}
}

type genesis struct {
	client module.Client
}

func (g genesis) Default() error {
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
	err = g.client.Create(&v1alpha1.CurrentBlock{})
	if err != nil {
		return err
	}
	err = g.client.Create(&v1alpha1.InitChainInfo{})
	if err != nil {
		return err
	}
	err = g.client.Create(&v1alpha1.ValidatorUpdates{})
	if err != nil {
		return err
	}
	err = g.client.Create(&v1alpha1.EndBlockState{})
	if err != nil {
		return err
	}
	return nil
}

func (g genesis) Import(state json.RawMessage) error {
	return nil
}

func (g genesis) Export() (state json.RawMessage, err error) {
	panic("implement me")
}
