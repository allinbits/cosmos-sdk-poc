package apiserver

import (
	"testing"

	"github.com/fdymylja/tmos/runtime"
	"github.com/stretchr/testify/require"
)

func TestModule(t *testing.T) {
	builder := runtime.NewBuilder()
	builder.AddModule(Module{})
	rt, err := builder.Build()
	if err != nil {
		t.Fatal(err)
	}
	require.NoError(t, rt.InitGenesis())
}
