package abci

import (
	"testing"

	"github.com/fdymylja/tmos/runtime"
)

func TestModule(t *testing.T) {
	b := runtime.NewBuilder()
	b.AddModule(NewModule())
	rt := b.Build()
	err := rt.Initialize()
	if err != nil {
		t.Fatal(err)
	}
}
