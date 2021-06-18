package runtime_test

import (
	"testing"

	"github.com/fdymylja/tmos/x/authn/v1alpha1"

	"github.com/fdymylja/tmos/runtime/module"

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

type dependentModule struct {
}

func (d dependentModule) Initialize(client module.Client) module.Descriptor {
	return module.NewDescriptorBuilder().
		NeedsStateTransition(&v1alpha1.MsgCreateAccount{}).
		Named("someDependent").Build()
}

func TestNeedsDependencyFails(t *testing.T) {
	builder := runtime.NewBuilder()

	builder.AddModule(dependentModule{})
	_, err := builder.Build()
	require.Error(t, err)
	require.EqualError(t, err, "unable to install modules: unable to install module dependencies: dependency cannot be accomplished: router: state transition not found: tmos.x.authn.v1alpha1.MsgCreateAccount")
}
