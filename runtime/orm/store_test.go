package orm_test

import (
	"testing"

	"github.com/fdymylja/tmos/runtime/kv"
	"github.com/fdymylja/tmos/runtime/orm"
	"github.com/fdymylja/tmos/runtime/orm/indexes"
	"github.com/fdymylja/tmos/runtime/orm/objects"
	crisis "github.com/fdymylja/tmos/x/crisis/v1alpha1"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

func TestIterator(t *testing.T) {
	okv := kv.NewBadger()
	obj := objects.NewStore(okv)
	idxKv := kv.NewBadger()
	idx := indexes.NewStore(idxKv)

	store := orm.NewStore(obj, idx)
	// register object
	require.NoError(t, store.RegisterObject(&crisis.InvariantHandler{}, crisis.InvariantHandlerSchema))
	// create object which belong to bank module
	obj1 := &crisis.InvariantHandler{
		StateTransition: "bank.v1alpha1.NonNegativeBalance",
		Module:          "bank",
		Route:           "/bank/non-negative-balance",
	}
	// create another object which belongs to bank module
	obj2 := &crisis.InvariantHandler{
		StateTransition: "bank.v1alpha1.SupplyMatchesBalances",
		Module:          "bank",
		Route:           "/bank/supply-matches-balances",
	}
	// create an object which belongs to staking
	obj3 := &crisis.InvariantHandler{
		StateTransition: "staking.v1alpha1.NonNegativeBondPool",
		Module:          "staking",
		Route:           "/staking/non-negative-bond-pool",
	}
	require.NoError(t, store.Create(obj1))
	require.NoError(t, store.Create(obj2))
	require.NoError(t, store.Create(obj3))

	iter, err := store.List(&crisis.InvariantHandler{}, orm.ListMatchField{
		Field: "module",
		Value: "bank",
	})
	require.NoError(t, err)

	o := new(crisis.InvariantHandler)
	require.NoError(t, iter.Get(o))
	require.True(t, proto.Equal(o, obj1))
	require.True(t, iter.Valid())

	iter.Next()
	require.NoError(t, iter.Get(o))
	require.True(t, proto.Equal(o, obj2))
	require.True(t, iter.Valid())

	iter.Next()
	require.False(t, iter.Valid())
	iter.Close()

	iter, err = store.List(&crisis.InvariantHandler{}, orm.ListMatchField{
		Field: "module",
		Value: "staking",
	})
	require.NoError(t, err)
	require.NoError(t, iter.Get(o))
	require.True(t, proto.Equal(o, obj3))

	iter.Next()
	require.False(t, iter.Valid())
}
