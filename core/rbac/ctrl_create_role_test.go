package rbac_test

import (
	"testing"

	runtimev1alpha1 "github.com/fdymylja/tmos/core/runtime/v1alpha1"
	"github.com/fdymylja/tmos/runtime"
	"github.com/fdymylja/tmos/runtime/authentication/user"
	rterr "github.com/fdymylja/tmos/runtime/errors"
	"github.com/stretchr/testify/require"
)

func TestCreateRole(t *testing.T) {
	builder := runtime.NewBuilder()
	rt, err := builder.Build()
	require.NoError(t, err)
	require.NoError(t, rt.InitGenesis())
	rt.EnableRBAC()

	t.Run("unauthorized", func(t *testing.T) {
		id := user.NewUsersFromString("no-authorizations")
		err = rt.Deliver(id, &runtimev1alpha1.CreateModuleDescriptors{})
		require.ErrorIs(t, err, rterr.ErrUnauthorized)
	})

}
