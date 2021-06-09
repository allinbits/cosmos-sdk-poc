package v1alpha1

import (
	"testing"

	"github.com/fdymylja/tmos/core/runtime/v1alpha1"
	"github.com/fdymylja/tmos/testdata/testpb"
	"github.com/stretchr/testify/require"
)

func TestRole_Extend(t *testing.T) {
	role := NewEmptyRole("bank")

	resources := role.GetResourcesForVerb(v1alpha1.Verb_Get)
	require.Len(t, resources, 0)

	err := role.Extend(v1alpha1.Verb_Get, &testpb.SimpleMessage{})
	require.NoError(t, err)

	resources = role.GetResourcesForVerb(v1alpha1.Verb_Get)
	require.Len(t, resources, 1)
}
