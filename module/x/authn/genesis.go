package authn

import (
	"encoding/json"

	"github.com/fdymylja/tmos/module/x/authn/v1alpha1"
	"github.com/fdymylja/tmos/runtime/module"
)

type genesis struct {
	c module.Client
}

func (g genesis) Default() error {
	err := g.c.Create(&v1alpha1.Params{
		MaxMemoCharacters:      1000000,
		TxSigLimit:             1000000,
		TxSizeCostPerByte:      0,
		SigVerifyCostEd25519:   0,
		SigVerifyCostSecp256K1: 0,
	})
	if err != nil {
		return err
	}

	err = g.c.Create(&v1alpha1.CurrentAccountNumber{Number: 0})
	if err != nil {
		return err
	}

	return err
}

func (g genesis) Import(state json.RawMessage) error {
	panic("implement me")
}

func (g genesis) Export() (json.RawMessage, error) {
	panic("implement me")
}
