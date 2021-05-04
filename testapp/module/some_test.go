package module

import (
	"context"
	"testing"

	"github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/rpc/client/http"
)

func TestBroadcastTx(t *testing.T) {

	tm, err := http.New("tcp://localhost:26657", "")
	if err != nil {
		t.Fatal(err)
	}
	status, err := tm.Status(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v", status)
	xd, err := Sign(privKey, status.NodeInfo.Network, accountAddress, "some_random_account", types.NewCoins(types.NewCoin("test", types.NewInt(500))))
	if err != nil {
		t.Fatal(err)
	}
	resp, err := tm.BroadcastTxCommit(context.Background(), xd)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v", resp)
}

func TestQuery(t *testing.T) {
	tm, err := http.New("tcp://localhost:26657", "")
	if err != nil {
		t.Fatal(err)
	}
	resp, err := tm.ABCIQuery(context.Background(), "get/tmos.x.bank.v1alpha1.Balance/test17hrfajk9ukj6tkkcy2ftgsmr3fp9hk9rkzcc7w", nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%s", resp.Response.Value)
}
