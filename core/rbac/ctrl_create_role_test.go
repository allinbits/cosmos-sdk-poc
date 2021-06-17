package rbac_test

import (
	"testing"

	"github.com/fdymylja/tmos/runtime"
	testmodule "github.com/fdymylja/tmos/testapp/module"
)

func TestCreateRole(t *testing.T) {
	builder := runtime.NewBuilder()
	builder.AddModule(testmodule.NewModule())
}
