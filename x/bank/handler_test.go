package bank_test

import (
	"testing"

	"github.com/fdymylja/tmos/core/coin/v1alpha1"
	"github.com/fdymylja/tmos/core/meta"
	"github.com/fdymylja/tmos/runtime/authentication/user"
	authn2 "github.com/fdymylja/tmos/x/authn"
	bank2 "github.com/fdymylja/tmos/x/bank"
	v1alpha12 "github.com/fdymylja/tmos/x/bank/v1alpha1"

	"github.com/stretchr/testify/require"

	"github.com/fdymylja/tmos/runtime"
)

func TestSendCoins(t *testing.T) {
	rtb := runtime.NewBuilder()
	rtb.AddModule(bank2.NewModule())
	rtb.AddModule(authn2.NewModule())
	rt, err := rtb.Build()
	require.NoError(t, err)
	err = rt.InitGenesis()
	require.NoError(t, err)
	rt.EnableRBAC()

	err = rt.Create(user.NewUsersFromString("bank"), &v1alpha12.Balance{
		Address: "frojdi",
		Balance: []*v1alpha1.Coin{
			{
				Denom:  "atom",
				Amount: "1000",
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	err = rt.Deliver(user.NewUsersFromString("bank"), &v1alpha12.MsgSendCoins{
		FromAddress: "frojdi",
		ToAddress:   "jonathan",
		Amount: []*v1alpha1.Coin{
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
	fBalance := &v1alpha12.Balance{}
	err = rt.Get(meta.NewStringID("frojdi"), fBalance)
	require.NoError(t, err)
	require.Equal(t, "1", fBalance.Balance[0].Amount)

	// Jonathans's balance == 999
	jBalance := &v1alpha12.Balance{}
	err = rt.Get(meta.NewStringID("jonathan"), jBalance)
	require.NoError(t, err)
	require.Equal(t, "999", jBalance.Balance[0].Amount)
}
