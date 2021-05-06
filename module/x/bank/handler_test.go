package bank_test

import (
	"testing"

	"github.com/fdymylja/tmos/runtime/authentication"
	"github.com/fdymylja/tmos/runtime/meta"

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

	err = rt.Create("TODO", &v1alpha1.Balance{
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

	err = rt.Deliver(authentication.NewSubjects("dio"), &v1alpha1.MsgSendCoins{
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

	// Frojdi's balance == 1
	fBalance := &v1alpha1.Balance{}
	err = rt.Get(meta.NewStringID("frojdi"), fBalance)
	require.NoError(t, err)
	require.Equal(t, "1", fBalance.Balance[0].Amount)

	// Jonathans's balance == 999
	jBalance := &v1alpha1.Balance{}
	err = rt.Get(meta.NewStringID("jonathan"), jBalance)
	require.NoError(t, err)
	require.Equal(t, "999", jBalance.Balance[0].Amount)
}
