package orm_test

import (
	"testing"

	"github.com/fdymylja/tmos/runtime/store/orm"
	crisis "github.com/fdymylja/tmos/x/crisis/v1alpha1"
)

func TestStore_RegisterObject(t *testing.T) {
	s := orm.NewStore(nil)
	err := s.RegisterObject(&crisis.InvariantHandler{}, orm.RegisterObjectOptions{
		PrimaryKey:    "stateTransition",
		SecondaryKeys: []string{"module", "route"},
	})
	if err != nil {
		t.Fatal(err)
	}
}
