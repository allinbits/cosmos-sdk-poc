package authn

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/fdymylja/tmos/module/x/authn/v1alpha1"
	"github.com/fdymylja/tmos/runtime"
	"github.com/fdymylja/tmos/runtime/meta"
)

func TestCreateAccountController(t *testing.T) {
	rtb := runtime.NewBuilder()
	rtb.AddModule(NewModule())
	rt, err := rtb.Build()
	require.NoError(t, err)

	err = rt.Deliver(
		[]string{"authn"},
		&v1alpha1.MsgCreateAccount{
			Account: &v1alpha1.Account{
				Address: "test",
				PubKey:  nil,
			},
		},
	)
	if err != nil {
		t.Fatal(err)
	}

	// get account
	acc := new(v1alpha1.Account)
	err = rt.Get(meta.NewStringID("test"), acc)
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, "test", acc.Address)

	t.Logf("%s", acc)
	// get last account number
	num := new(v1alpha1.CurrentAccountNumber)
	err = rt.Get(num.GetID(), num)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%s", num)
}
