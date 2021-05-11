package extensions_test

import (
	"testing"
	"time"

	"github.com/cosmos/cosmos-sdk/types"
	app "github.com/fdymylja/tmos/testapp/app"
	"github.com/fdymylja/tmos/testapp/module"
	"github.com/stretchr/testify/require"
	types2 "github.com/tendermint/tendermint/abci/types"
)

const privKey = "f44351066b09af7e8b1c98de10214a3eeb8f60b01867b75867f27f162613e3a6"
const pubKeyAsAny = "0a2103695a0767494e7b8a161b6a561522b32e1129c10d766912f4a7441766d0d55e06"
const pubKeyType = "/cosmos.crypto.secp256k1.PubKey"
const accountAddress = "test17hrfajk9ukj6tkkcy2ftgsmr3fp9hk9rkzcc7w"

func TestSigVerification(t *testing.T) {
	tm := app.NewApp()
	tm.InitChain(types2.RequestInitChain{
		Time:            time.Time{},
		ChainId:         "test",
		ConsensusParams: nil,
		Validators:      nil,
		AppStateBytes:   nil,
		InitialHeight:   0,
	})
	xd, err := module.Sign(privKey, "test", accountAddress, "some_rankdf_a'ccount", types.NewCoins(types.NewCoin("test", types.NewInt(500))))
	if err != nil {
		t.Fatal(err)
	}
	resp := tm.DeliverTx(types2.RequestDeliverTx{Tx: xd})
	require.Equal(t, resp.Code, uint32(0))

	newTx, err := module.Sign(privKey, "test", accountAddress, "some_rankdf_a'ccount", types.NewCoins(types.NewCoin("test", types.NewInt(500))))
	if err != nil {
		t.Fatal(err)
	}

	resp = tm.DeliverTx(types2.RequestDeliverTx{Tx: newTx})
	require.NotEqual(t, resp.Code, uint32(0))
	t.Logf("%s", resp.Log)
}
