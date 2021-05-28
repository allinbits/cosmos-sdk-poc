package module

import (
	"context"
	"testing"

	"github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
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
	xd, err := Sign(privKey, status.NodeInfo.Network, accountAddress, "some_rankdf_a'ccount", types.NewCoins(types.NewCoin("test", types.NewInt(500))))
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
	resp, err := tm.ABCIQuery(context.Background(), "api_resources", nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%s %s", resp.Response.Log, resp.Response.Value)

	resp, err = tm.ABCIQuery(context.Background(), "get/tmos.x.bank.v1alpha1/Balance/fee_collector", nil)
	require.NoError(t, err)
	t.Logf("%s %s", resp.Response.Log, resp.Response.Value)
}
