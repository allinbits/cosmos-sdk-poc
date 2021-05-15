package orm

import (
	"fmt"
	"math"

	"google.golang.org/protobuf/encoding/protowire"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type FieldEncoderFunc func(value protoreflect.Value) []byte

// TODO maybe we should not support all of these... floats/doubles?
// TODO we can preallocate a lot of those slices
var protowireFieldEncoders = map[protoreflect.Kind]FieldEncoderFunc{
	protoreflect.BoolKind: func(v protoreflect.Value) []byte {
		var b []byte
		b = protowire.AppendVarint(b, protowire.EncodeBool(v.Bool()))
		return b
	},
	protoreflect.EnumKind: func(value protoreflect.Value) []byte {
		var b []byte
		b = protowire.AppendVarint(b, uint64(value.Enum()))
		return b
	},
	protoreflect.Int32Kind: func(value protoreflect.Value) []byte {
		var b []byte
		b = protowire.AppendVarint(b, uint64(int32(value.Int())))
		return b
	},
	protoreflect.Sint32Kind: func(value protoreflect.Value) []byte {
		var b []byte
		b = protowire.AppendVarint(b, protowire.EncodeZigZag(int64(int32(value.Int()))))
		return b
	},
	protoreflect.Uint32Kind: func(value protoreflect.Value) []byte {
		var b []byte
		b = protowire.AppendVarint(b, uint64(uint32(value.Uint())))
		return b
	},
	protoreflect.Int64Kind: func(value protoreflect.Value) []byte {
		var b []byte
		b = protowire.AppendVarint(b, uint64(value.Int()))
		return b
	},
	protoreflect.Sint64Kind: func(value protoreflect.Value) []byte {
		var b []byte
		b = protowire.AppendVarint(b, protowire.EncodeZigZag(value.Int()))
		return b
	},
	protoreflect.Uint64Kind: func(value protoreflect.Value) []byte {
		var b []byte
		b = protowire.AppendVarint(b, value.Uint())
		return b
	},
	protoreflect.Sfixed32Kind: func(value protoreflect.Value) []byte {
		var b []byte
		b = protowire.AppendFixed32(b, uint32(value.Int()))
		return b
	},
	protoreflect.Fixed32Kind: func(value protoreflect.Value) []byte {
		var b []byte
		b = protowire.AppendFixed32(b, uint32(value.Uint()))
		return b
	},
	protoreflect.FloatKind: func(value protoreflect.Value) []byte {
		var b []byte
		b = protowire.AppendFixed32(b, math.Float32bits(float32(value.Float())))
		return b
	},
	protoreflect.Sfixed64Kind: func(value protoreflect.Value) []byte {
		var b []byte
		b = protowire.AppendFixed64(b, uint64(value.Int()))
		return b
	},
	protoreflect.Fixed64Kind: func(value protoreflect.Value) []byte {
		var b []byte
		b = protowire.AppendFixed64(b, value.Uint())
		return b
	},
	protoreflect.DoubleKind: func(value protoreflect.Value) []byte {
		var b []byte
		b = protowire.AppendFixed64(b, math.Float64bits(value.Float()))
		return b
	},
	protoreflect.StringKind: func(value protoreflect.Value) []byte {
		var b []byte
		// NOTE: skipping UTF8 checks, anyways marshalling would fail
		// if the string is invalid.
		// NOTE2: this prepends the string length which we can do without..
		b = protowire.AppendString(b, value.String())
		return b
	},
	// NOTE: do we really need to index bytes? when would it be useful?
	protoreflect.BytesKind: func(value protoreflect.Value) []byte {
		var b []byte
		b = protowire.AppendBytes(b, value.Bytes())
		return b
	},
}

func EncoderForKind(kind protoreflect.Kind) (FieldEncoderFunc, error) {
	encoder, exists := protowireFieldEncoders[kind]
	if !exists {
		return nil, fmt.Errorf("store: unsupported secondary index with kind %s", kind)
	}
	return encoder, nil
}
