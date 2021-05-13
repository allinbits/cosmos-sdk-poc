package runtime_test

import (
	"testing"

	"github.com/fdymylja/tmos/runtime"
	bank2 "github.com/fdymylja/tmos/x/bank"

	"github.com/stretchr/testify/require"
)

func TestRuntime_Import(t *testing.T) {
	rb := runtime.NewBuilder()
	rb.AddModule(bank2.NewModule())
	rt, err := rb.Build()
	require.NoError(t, err)

	genesisData := `{
	"bank" : {}
}`
	err = rt.Import([]byte(genesisData))
	require.NoError(t, err)
}
