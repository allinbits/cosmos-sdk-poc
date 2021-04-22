package orm_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"

	accounting "github.com/fdymylja/tmos/apis/accounting/v1alpha1"
	meta "github.com/fdymylja/tmos/apis/core/meta/v1alpha1"
	"github.com/fdymylja/tmos/pkg/runtime/store"
)

func TestStore(t *testing.T) {
	s := orm.NewStore()
	s.RegisterStateObject(&accounting.Balance{})

	require.Panics(t, func() {
		s.Set(&accounting.Test{})
	})

	require.Panics(t, func() {
		s.Set(&accounting.Balance{})
	})

	require.Panics(t, func() {
		s.Set(&accounting.Balance{
			Meta:     &meta.Meta{Id: ""},
			Currency: "",
			Amount:   0,
		})
	})

	setObj := &accounting.Balance{
		Meta:     &meta.Meta{Id: "hello"},
		Currency: "1",
		Amount:   0,
	}
	s.Set(setObj)
	getObj := &accounting.Balance{Meta: &meta.Meta{Id: "hello"}}
	exists := s.Get(getObj)
	require.True(t, exists)
	require.True(t, proto.Equal(getObj, setObj))
}
