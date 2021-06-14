package prototest

import (
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

func Equal(t *testing.T, expected proto.Message, got proto.Message) {
	require.True(t, proto.Equal(expected, got)) // TODO do me better than this
}
