package abci_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/fdymylja/tmos/runtime"
)

func TestModule(t *testing.T) {
	b := runtime.NewBuilder() // we're not adding because as of now its already installed

	rt, err := b.Build()
	require.NoError(t, err)

	err = rt.Initialize()
	if err != nil {
		t.Fatal(err)
	}
}
