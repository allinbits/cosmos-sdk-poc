package authn

import (
	"testing"

	"github.com/fdymylja/tmos/core/meta"
	"github.com/fdymylja/tmos/core/rbac"
	"github.com/fdymylja/tmos/runtime/authentication/user"
	v1alpha12 "github.com/fdymylja/tmos/x/authn/v1alpha1"
	"github.com/stretchr/testify/require"

	"github.com/fdymylja/tmos/runtime"
)

func TestCreateAccountController(t *testing.T) {
	rtb := runtime.NewBuilder()
	rtb.AddModule(NewModule())
	rtb.AddModule(rbac.NewModule())
	rt, err := rtb.Build()
	require.NoError(t, err)

	err = rt.Deliver(
		user.NewUsers(),
		&v1alpha12.MsgCreateAccount{
			Account: &v1alpha12.Account{
				Address: "test",
				PubKey:  nil,
			},
		},
	)
	if err != nil {
		t.Fatal(err)
	}

	// get account
	acc := new(v1alpha12.Account)
	err = rt.Get(meta.NewStringID("test"), acc)
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, "test", acc.Address)

	t.Logf("%s", acc)
	// get last account number
	num := new(v1alpha12.CurrentAccountNumber)
	err = rt.Get(meta.SingletonID, num)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%s", num)
}
