//+build integration_test

package client

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestClient_Resources(t *testing.T) {
	c, err := NewRPC("tcp://localhost:26657")
	require.NoError(t, err)
	res, err := c.Resources(context.TODO())
	require.NoError(t, err)
	t.Logf("%s", res)
}
