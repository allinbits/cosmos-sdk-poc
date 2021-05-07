package runtime_test

import (
	"testing"

	"github.com/fdymylja/tmos/runtime"

	"github.com/stretchr/testify/require"

	"github.com/fdymylja/tmos/module/x/bank"
)

func TestRuntime_Import(t *testing.T) {
	rb := runtime.NewBuilder()
	rb.AddModule(bank.NewModule())
	rt, err := rb.Build()
	require.NoError(t, err)

	genesisData := `{
	"bank" : {}
}`
	err = rt.Import([]byte(genesisData))
	require.NoError(t, err)
}
