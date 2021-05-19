package indexes_test

import (
	"testing"

	kv "github.com/fdymylja/tmos/runtime/kv"
	"github.com/fdymylja/tmos/runtime/orm/indexes"
	"github.com/fdymylja/tmos/runtime/orm/schema"
	crisis "github.com/fdymylja/tmos/x/crisis/v1alpha1"
	"github.com/stretchr/testify/require"
)

func TestStore(t *testing.T) {
	reg := schema.NewRegistry()
	require.NoError(t, reg.AddObject(
		&crisis.InvariantHandler{},
		schema.Options{
			PrimaryKey:    "stateTransition",
			SecondaryKeys: []string{"module", "route"},
		},
	))
	sch, err := reg.Get(&crisis.InvariantHandler{})
	require.NoError(t, err)

	kv := kv.NewBadger()
	store := indexes.NewStore(kv)
	// test indexing
	obj := &crisis.InvariantHandler{
		StateTransition: "/someTransition",
		Module:          "bank",
		Route:           "/invariance/bank",
	}
	err = store.Index(sch, obj)
	require.NoError(t, err)
	// test list by matching fields
	x, err := store.List(sch, indexes.MatchField("module", "bank"))
	require.NoError(t, err)
	require.True(t, x.Valid())
	// test unindexing
	err = store.ClearIndexes(sch, obj)
	require.NoError(t, err)
	// test list is invalid
	x, err = store.List(sch, indexes.MatchField("module", "bank"))
	require.Nil(t, x)
	require.ErrorIs(t, err, indexes.ErrNotFound)
}
