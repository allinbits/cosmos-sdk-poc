package authn

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	"github.com/fdymylja/tmos/module/x/authn/v1alpha1"
	"github.com/fdymylja/tmos/runtime"
	gogoproto "github.com/gogo/protobuf/proto"
	"github.com/tendermint/tendermint/abci/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
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
		Messages: []*anypb.Any{
			{
				TypeUrl: "/tmos.x.authn.v1alpha1.MsgCreateAccount",
				Value:   nil,
			},
		},
		TimeoutHeight: 50,
	}
	b, err := proto.Marshal(body)
	if err != nil {
		t.Fatal(err)
	}
	pk := secp256k1.GenPrivKey().PubKey()
	pkB, err := gogoproto.Marshal(pk)
	if err != nil {
		t.Fatal(err)
	}
	auth := &v1alpha1.AuthInfo{
		SignerInfos: []*v1alpha1.SignerInfo{
			{
				PublicKey: &anypb.Any{
					TypeUrl: "/" + gogoproto.MessageName(pk),
					Value:   pkB,
				},
				ModeInfo: nil,
				Sequence: 5,
			},
		},
		Fee: nil,
	}

	authB, err := proto.Marshal(auth)
	if err != nil {
		t.Fatal(err)
	}
	txRaw, err := proto.Marshal(&v1alpha1.TxRaw{
		BodyBytes:     b,
		AuthInfoBytes: authB,
		Signatures:    [][]byte{[]byte("yahallo")},
	})
	if err != nil {
		t.Fatal(err)
	}

	return txRaw
}
