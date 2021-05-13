package authn_test

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	"github.com/fdymylja/tmos/module/coin/v1alpha1"
	"github.com/fdymylja/tmos/runtime/authentication/user"
	authn2 "github.com/fdymylja/tmos/x/authn"
	v1alpha12 "github.com/fdymylja/tmos/x/authn/v1alpha1"
	bank2 "github.com/fdymylja/tmos/x/bank"
	v1alpha13 "github.com/fdymylja/tmos/x/bank/v1alpha1"
	gogoproto "github.com/gogo/protobuf/proto"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/abci/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"

	"github.com/fdymylja/tmos/runtime"
)

func TestAuthenticator(t *testing.T) {
	rtb := runtime.NewBuilder()
	auth := authn2.NewModule()
	rtb.AddModule(bank2.NewModule())
	rtb.AddModule(auth)
	rtb.SetDecoder(auth.GetTxDecoder())
	rt, err := rtb.Build()
	require.NoError(t, err)

	err = rt.InitGenesis()
	if err != nil {
		t.Fatal(err)
	}

	// initialize with money...
	err = rt.Deliver(user.NewUsersFromString("bank"), &v1alpha13.MsgSetBalance{
		Address: "frojdi",
		Amount: []*v1alpha1.Coin{{
			Denom:  "test",
			Amount: "5000",
		}},
	})
	if err != nil {
		t.Fatal(err)
	}

	// get abci app
	app := runtime.NewABCIApplication(rt)

	// run a begin block tx set forward
	app.BeginBlock(types.RequestBeginBlock{
		Header: tmproto.Header{
			Height: 1,
		},
	})
	// test timeout header
	timedOutTx := timedOutTx(t)
	resp := app.DeliverTx(types.RequestDeliverTx{Tx: timedOutTx})
	t.Logf("%#v", resp)
}

func timedOutTx(t *testing.T) []byte {
	msg := &v1alpha12.MsgCreateAccount{Account: &v1alpha12.Account{
		Address: "cookies",
		PubKey: &anypb.Any{
			TypeUrl: "pk",
			Value:   nil,
		},
	}}
	anyMsg, err := anypb.New(msg)
	if err != nil {
		t.Fatal(err)
	}
	body := &v1alpha12.TxBody{
		Messages:      []*anypb.Any{anyMsg},
		TimeoutHeight: 50,
	}
	b, err := proto.Marshal(body)
	if err != nil {
		t.Fatal(err)
	}
	privKey := secp256k1.GenPrivKey()
	t.Logf("%x", privKey.Key)
	pk := privKey.PubKey()
	pkB, err := gogoproto.Marshal(pk)
	if err != nil {
		t.Fatal(err)
	}
	auth := &v1alpha12.AuthInfo{
		SignerInfos: []*v1alpha12.SignerInfo{
			{
				PublicKey: &anypb.Any{
					TypeUrl: "/" + gogoproto.MessageName(pk),
					Value:   pkB,
				},
				ModeInfo: &v1alpha12.ModeInfo{Sum: &v1alpha12.ModeInfo_Single_{Single: &v1alpha12.ModeInfo_Single{
					Mode: v1alpha12.SignMode_SIGN_MODE_DIRECT,
				}}},
				Sequence: 5,
			},
		},
		Fee: &v1alpha12.Fee{
			Amount: []*v1alpha1.Coin{
				{
					Denom:  "test",
					Amount: "1000",
				},
			},
			GasLimit: 0,
			Payer:    "frojdi",
			Granter:  "",
		},
	}
	t.Logf("%s %x", auth.SignerInfos[0].PublicKey.TypeUrl, auth.SignerInfos[0].PublicKey.Value)
	authB, err := proto.Marshal(auth)
	if err != nil {
		t.Fatal(err)
	}
	txRaw, err := proto.Marshal(&v1alpha12.TxRaw{
		BodyBytes:     b,
		AuthInfoBytes: authB,
		Signatures:    [][]byte{[]byte("yahallo")},
	})
	if err != nil {
		t.Fatal(err)
	}

	return txRaw
}
