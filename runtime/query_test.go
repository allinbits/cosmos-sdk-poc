package runtime

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/abci/types"
)

func TestQuerier_apiResources(t *testing.T) {
	rtb := NewBuilder()
	rt, err := rtb.Build()
	require.NoError(t, err)

	querier := NewQuerier(rt.store)
	resp, err := querier.apiResources(types.RequestQuery{})
	require.NoError(t, err)

	t.Logf("%s", resp)
}
