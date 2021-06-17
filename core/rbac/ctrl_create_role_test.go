package rbac_test

import (
	"testing"

	rbacv1alpha1 "github.com/fdymylja/tmos/core/rbac/v1alpha1"
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
		err = rt.Deliver(id, &rbacv1alpha1.MsgCreateRole{NewRole: &rbacv1alpha1.Role{Id: "a-role"}})
		require.ErrorIs(t, err, rterr.ErrUnauthorized)
	})

	t.Run("authorized", func(t *testing.T) {
		id := user.NewUsersFromString("rbac")
		err = rt.Deliver(id, &rbacv1alpha1.MsgCreateRole{NewRole: &rbacv1alpha1.Role{Id: "a-role"}})
		require.NoError(t, err)
	})
}
