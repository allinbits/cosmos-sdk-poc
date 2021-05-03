package bank_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	coin "github.com/fdymylja/tmos/module/core/coin/v1alpha1"
	"github.com/fdymylja/tmos/module/x/authn"
	"github.com/fdymylja/tmos/module/x/bank"
	"github.com/fdymylja/tmos/module/x/bank/v1alpha1"
	"github.com/fdymylja/tmos/runtime"
)

func TestSendCoins(t *testing.T) {
	rtb := runtime.NewBuilder()
	rtb.AddModule(bank.NewModule())
	rtb.AddModule(authn.NewModule())
	rt, err := rtb.Build()
	require.NoError(t, err)

	err = rt.Create("dio", &v1alpha1.Balance{
		Address: "frojdi",
		Balance: []*coin.Coin{
			{
				Denom:  "atom",
				Amount: "1000",
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	err = rt.Deliver([]string{"dio"}, &v1alpha1.MsgSendCoins{
		FromAddress: "frojdi",
		ToAddress:   "jonathan",
		Amount: []*coin.Coin{
			{
				Denom:  "atom",
				Amount: "999",
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}
}
