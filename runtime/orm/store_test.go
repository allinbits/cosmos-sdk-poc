package orm_test

import (
	"testing"

	"github.com/fdymylja/tmos/runtime/kv"
	orm2 "github.com/fdymylja/tmos/runtime/orm"
	crisis "github.com/fdymylja/tmos/x/crisis/v1alpha1"
)

func TestStore_RegisterObject(t *testing.T) {
	s := orm2.NewStore(kv.NewBadger())
	err := s.RegisterObject(&crisis.InvariantHandler{}, orm2.RegisterObjectOptions{
		PrimaryKey:    "stateTransition",
		SecondaryKeys: []string{"module", "route"},
	})
	if err != nil {
		t.Fatal(err)
	}
	err = s.Create(&crisis.InvariantHandler{
		StateTransition: "bank.NonNegativeBalance",
		Module:          "bank",
		Route:           "/bank/invariant",
	})
	if err != nil {
		t.Fatal(err)
	}
}
