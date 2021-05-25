package schema

import (
	"testing"

	"github.com/fdymylja/tmos/testdata/testpb"
	"github.com/stretchr/testify/require"
)

func TestRegistry(t *testing.T) {
	reg := NewRegistry()
	// test unknown
	_, err := reg.Get(&testpb.SimpleMessage{})
	require.ErrorIs(t, err, ErrNotFound)
	// test simple add
	require.NoError(t, reg.AddObject(&testpb.SimpleMessage{}, Options{PrimaryKey: "a"}))
	// test double registration
	require.ErrorIs(t, reg.AddObject(&testpb.SimpleMessage{}, Options{PrimaryKey: "a"}), ErrAlreadyExists)
	// test get
	_, err = reg.Get(&testpb.SimpleMessage{})
	require.NoError(t, err)
}
