package oas3schema

import (
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/known/anypb"
)

func TestNewOpenAPIv3Generator(t *testing.T) {
	fd, err := protoregistry.GlobalFiles.FindFileByPath("google/protobuf/any.proto")
	require.NoError(t, err)
	md := fd.Messages().ByName("Any")

	require.NoError(t, err)

	oas3 := NewOpenAPIv3Generator()
	require.NoError(t, oas3.AddRequiredMessage(md))
	require.NoError(t, oas3.AddRawOperation("GET", "uniqueID", "some comment", "/hello/cookies", "", &anypb.Any{}))
	_, err = oas3.Build()
	require.NoError(t, err)
}
