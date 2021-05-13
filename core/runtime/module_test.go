package runtime_test

import (
	"testing"

	"github.com/fdymylja/tmos/runtime"
)

func TestModule(t *testing.T) {
	builder := runtime.NewBuilder() // note we're not adding the runtime core because it comes as default
	rt, err := builder.Build()
	if err != nil {
		t.Fatal(err)
	}
	err = rt.InitGenesis()
	if err != nil {
		t.Fatal(err)
	}
	// we should have in store the registered state objects and state transitions
	// TODO test :)
}
