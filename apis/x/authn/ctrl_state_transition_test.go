package authn

import (
	"testing"

	meta "github.com/fdymylja/tmos/apis/meta/v1alpha1"
	"github.com/fdymylja/tmos/apis/x/authn/v1alpha1"
	"github.com/fdymylja/tmos/pkg/runtime"
)

func TestCreateAccountController(t *testing.T) {
	rtb := runtime.NewBuilder()
	rtb.AddModule(NewModule())
	rt := rtb.Build()
	err := rt.Deliver([]string{"authn"}, &v1alpha1.MsgCreateAccount{Account: &v1alpha1.Account{
		Address:       "test",
		PubKey:        nil,
		AccountNumber: 0,
		ObjectMeta:    &meta.ObjectMeta{Id: "test"},
	}}, false)

	if err != nil {
		t.Fatal(err)
	}
	// get account
	acc := &v1alpha1.Account{ObjectMeta: &meta.ObjectMeta{Id: "test"}}
	err = rt.Get(acc)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%s", acc)
	// get last account number
	num := new(v1alpha1.CurrentAccountNumber)
	err = rt.Get(num)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%s", num)
}
