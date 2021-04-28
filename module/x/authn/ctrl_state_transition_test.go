package authn

import (
	runtime2 "github.com/fdymylja/tmos/runtime"
	"github.com/fdymylja/tmos/runtime/meta"

	"testing"

	"github.com/fdymylja/tmos/module/x/authn/v1alpha1"
)

func TestCreateAccountController(t *testing.T) {
	rtb := runtime2.NewBuilder()
	rtb.AddModule(NewModule())
	rt := rtb.Build()
	err := rt.Deliver([]string{"authn"}, &v1alpha1.MsgCreateAccount{Account: &v1alpha1.Account{
		Address:       "test",
		PubKey:        nil,
		AccountNumber: 0,
	}}, false)

	if err != nil {
		t.Fatal(err)
	}
	// get account
	acc := new(v1alpha1.Account)
	err = rt.Get(meta.NewStringID("test"), acc)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%s", acc)
	// get last account number
	num := new(v1alpha1.CurrentAccountNumber)
	err = rt.Get(num.GetID(), num)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%s", num)
}
