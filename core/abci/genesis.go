package abci

import (
	"encoding/json"

	"github.com/fdymylja/tmos/core/abci/v1alpha1"
	"github.com/fdymylja/tmos/runtime/module"
)

func newGenesisHandler() module.GenesisHandler {
	return genesis{}
}

type genesis struct {
}

func (g genesis) Default(client module.Client) error {
	// set initial stage
	err := client.Create(&v1alpha1.Stage{Stage: v1alpha1.ABCIStage_InitChain})
	if err != nil {
		return err
	}
	// set empty begin block state
	err = client.Create(&v1alpha1.BeginBlockState{})
	if err != nil {
		return err
	}
	// set empty check tx state
	err = client.Create(&v1alpha1.CheckTxState{})
	if err != nil {
		return err
	}
	// set empty deliver tx state
	err = client.Create(&v1alpha1.DeliverTxState{})
	if err != nil {
		return err
	}
	err = client.Create(&v1alpha1.CurrentBlock{})
	if err != nil {
		return err
	}
	err = client.Create(&v1alpha1.InitChainInfo{})
	if err != nil {
		return err
	}
	err = client.Create(&v1alpha1.ValidatorUpdates{})
	if err != nil {
		return err
	}
	err = client.Create(&v1alpha1.EndBlockState{})
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
