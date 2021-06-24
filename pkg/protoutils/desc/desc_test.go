package desc_test

import (
	"testing"

	abciv1alpha1 "github.com/fdymylja/tmos/core/abci/v1alpha1"
	"github.com/fdymylja/tmos/pkg/protoutils/desc"
)

func TestDependencies(t *testing.T) {
	fd := (&abciv1alpha1.MsgSetInitChain{}).ProtoReflect().Descriptor().ParentFile()

	deps := desc.Dependencies(fd)

	for _, dep := range deps {
		t.Logf("%s", dep.Path())
	}
}
