package runtime_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/fdymylja/tmos/runtime"
	"github.com/fdymylja/tmos/runtime/module"
)

type TestModule struct {
}

func (t TestModule) Initialize(client module.Client, builder *module.Builder) {
}

func TestNewBuilder_ModuleWithoutName(t *testing.T) {
	builder := runtime.NewBuilder()
	builder.AddModule(&TestModule{})

	_, err := builder.Build()
	require.NotNil(t, err)
	require.Contains(t, err.Error(), "empty module name")
}
