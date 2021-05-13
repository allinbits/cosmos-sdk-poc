package module

import (
	"encoding/hex"
	"encoding/json"

	coin "github.com/fdymylja/tmos/module/core/coin/v1alpha1"
	"github.com/fdymylja/tmos/runtime/module"
	v1alpha12 "github.com/fdymylja/tmos/x/authn/v1alpha1"
	"github.com/fdymylja/tmos/x/bank/v1alpha1"
	"google.golang.org/protobuf/types/known/anypb"
)

const privKey = "f44351066b09af7e8b1c98de10214a3eeb8f60b01867b75867f27f162613e3a6"
const pubKeyAsAny = "0a2103695a0767494e7b8a161b6a561522b32e1129c10d766912f4a7441766d0d55e06"
const pubKeyType = "/cosmos.crypto.secp256k1.PubKey"
const accountAddress = "test17hrfajk9ukj6tkkcy2ftgsmr3fp9hk9rkzcc7w"

func NewModule() Module {
	return Module{}
}

// Module implements a simple test module which during init genesis
// sets a default account with some money
type Module struct {
}

func (m Module) Initialize(client module.Client) module.Descriptor {
	return module.NewDescriptorBuilder().
		Named("testing").
		WithGenesis(newGenesisController(client)).Build()
}

func newGenesisController(client module.Client) genesisController {
	return genesisController{
		authn: v1alpha12.NewClient(client),
		bank:  v1alpha1.NewClient(client),
	}
}

type genesisController struct {
	authn *v1alpha12.Client
	bank  *v1alpha1.Client
}

func (g genesisController) Default() error {
	pkB, err := hex.DecodeString(pubKeyAsAny)
	if err != nil {
		return err
	}
	// create account
	acc := &v1alpha12.Account{
		Address: accountAddress,
		PubKey: &anypb.Any{
			TypeUrl: pubKeyType,
			Value:   pkB,
		},
	}
	err = g.authn.CreateAccount(acc)
	if err != nil {
		return err
	}
	// set an initial balance for the given account
	err = g.bank.SetBalance(acc.Address, []*coin.Coin{
		{
			Denom:  "test",
			Amount: "5000000000",
		},
	})
	if err != nil {
		return err
	}
	return nil
}

func (g genesisController) Import(state json.RawMessage) error {
	panic("implement me")
}

func (g genesisController) Export() (json.RawMessage, error) {
	panic("implement me")
}
