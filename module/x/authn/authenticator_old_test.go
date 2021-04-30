package authn

import (
	"testing"

	"github.com/fdymylja/tmos/module/x/authn/v1alpha1"
	"github.com/fdymylja/tmos/runtime"
	"github.com/tendermint/tendermint/abci/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	"google.golang.org/protobuf/proto"
)

func TestAuthenticator(t *testing.T) {
	rtb := runtime.NewBuilder()
	auth := NewModule()
	rtb.AddModule(auth)
	rtb.SetAuthenticator(auth.GetAuthenticator())
	rt := rtb.Build()
	err := rt.Initialize()
	if err != nil {
		t.Fatal(err)
	}
	// get abci app
	app := runtime.NewABCIApplication(rt)
	// run a begin block tx set forward
	app.BeginBlock(types.RequestBeginBlock{
		Header: tmproto.Header{
			Height: 10000,
		},
	})
	// test timeout header
	timedOutTx := timedOutTx(t)
	resp := app.DeliverTx(types.RequestDeliverTx{Tx: timedOutTx})
	t.Logf("%#v", resp)
}

func timedOutTx(t *testing.T) []byte {
	body := &v1alpha1.TxBody{
		TimeoutHeight: 50,
	}
	b, err := proto.Marshal(body)
	if err != nil {
		t.Fatal(err)
	}
	txRaw, err := proto.Marshal(&v1alpha1.TxRaw{
		BodyBytes:     b,
		AuthInfoBytes: nil,
		Signatures:    nil,
	})
	if err != nil {
		t.Fatal(err)
	}

	return txRaw
}
