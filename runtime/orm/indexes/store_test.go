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
	kv := kv.NewBadger()
	store := indexes.NewStore(kv, reg)
	obj := &crisis.InvariantHandler{
		StateTransition: "/someTransition",
		Module:          "bank",
		Route:           "/invariance/bank",
	}
	err := store.IndexObject(obj)
	require.NoError(t, err)
	err = store.UnindexObject(obj)
	require.NoError(t, err)
	x, err := store.List(obj, indexes.MatchField("module", "bank"))
	require.NoError(t, err)
	t.Logf("%s, %v", x.Key(), x.Key())
}
