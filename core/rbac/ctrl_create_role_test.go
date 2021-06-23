package rbac_test

import (
	"testing"

	rbacv1alpha1 "github.com/fdymylja/tmos/core/rbac/v1alpha1"
	"github.com/fdymylja/tmos/runtime"
	"github.com/fdymylja/tmos/runtime/authentication/user"
	"github.com/fdymylja/tmos/runtime/client"
	rterr "github.com/fdymylja/tmos/runtime/errors"
	"github.com/stretchr/testify/require"
)

func TestRBAC(t *testing.T) {
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

		err = rt.Deliver(id, &rbacv1alpha1.MsgBindRole{
			RoleId:  "a-role",
			Subject: "someone",
		})
		require.NoError(t, err)

		err = rt.Deliver(id, &rbacv1alpha1.MsgBindRole{
			RoleId:  "a-role",
			Subject: "someone-else",
		})
		require.NoError(t, err)

		c := client.NewModuleClient(runtime.NewRuntimeServer(rt))
		c.(interface {
			SetUser(users user.Users)
		}).SetUser(user.NewUsersFromString("rbac"))

		rbacClient := rbacv1alpha1.NewClientSet(c)
		iter, err := rbacClient.RoleBindings().List(client.ListMatchFieldString("roleRef", "a-role"))
		require.NoError(t, err)

		for ; iter.Valid(); iter.Next() {
			role, err := iter.Get()
			require.NoError(t, err)
			t.Logf("%s", role)
		}
	})
}
