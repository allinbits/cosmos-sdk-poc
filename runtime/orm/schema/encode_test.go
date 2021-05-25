package schema

import (
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func Test_interfaceToValueEncoder(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		toValue, err := interfaceToValueEncoderForKind(protoreflect.BytesKind)
		require.NoError(t, err)

		expected := []byte("feel.it.still")
		v, ok := toValue(expected)
		require.True(t, ok)

		got := v.Bytes()
		require.EqualValues(t, expected, got)
	})

	t.Run("type not supported", func(t *testing.T) {
		_, err := interfaceToValueEncoderForKind(protoreflect.MessageKind)
		require.Error(t, err)
	})
}

func Test_encoderForKind(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		enc, err := encoderForKind(protoreflect.BytesKind)
		require.NoError(t, err)

		expected := []byte("feel.it.still")
		got := enc(protoreflect.ValueOfBytes(expected))
		require.EqualValues(t, expected, got)
	})

	t.Run("type not supported", func(t *testing.T) {
		_, err := encoderForKind(protoreflect.MessageKind)
		require.Error(t, err)
	})
}
