package kindencoder

import (
	"math"

	"google.golang.org/protobuf/encoding/protowire"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// NOTE(fdymylja): this are the types that we should decide if to support in kind encoder (and hence to index)

var safeInterfaceToValue = map[protoreflect.Kind]func(i interface{}) (value protoreflect.Value, valid bool){
	protoreflect.EnumKind: func(i interface{}) (value protoreflect.Value, valid bool) {
		v, ok := i.(int32)
		if !ok {
			return
		}
		return protoreflect.ValueOfEnum(protoreflect.EnumNumber(v)), true
	},
	protoreflect.Sint32Kind: func(i interface{}) (value protoreflect.Value, valid bool) {
		v, ok := i.(int32)
		if !ok {
			return
		}
		return protoreflect.ValueOfInt32(v), true
	},
	protoreflect.Sint64Kind: func(i interface{}) (value protoreflect.Value, valid bool) {
		v, ok := i.(int64)
		if !ok {
			return
		}
		return protoreflect.ValueOfInt64(v), true
	},
	protoreflect.Sfixed32Kind: func(i interface{}) (value protoreflect.Value, valid bool) {
		v, ok := i.(int32)
		if !ok {
			return
		}
		return protoreflect.ValueOfInt32(v), true
	},
	protoreflect.Fixed32Kind: func(i interface{}) (value protoreflect.Value, valid bool) {
		v, ok := i.(uint32)
		if !ok {
			return
		}
		return protoreflect.ValueOfUint32(v), true
	},
	protoreflect.FloatKind: func(i interface{}) (value protoreflect.Value, valid bool) {
		v, ok := i.(float32)
		if !ok {
			return
		}
		return protoreflect.ValueOfFloat32(v), true
	},
	protoreflect.Sfixed64Kind: func(i interface{}) (value protoreflect.Value, valid bool) {
		v, ok := i.(int64)
		if !ok {
			return
		}
		return protoreflect.ValueOfInt64(v), true
	},
	protoreflect.Fixed64Kind: func(i interface{}) (value protoreflect.Value, valid bool) {
		v, ok := i.(uint64)
		if !ok {
			return
		}
		return protoreflect.ValueOfUint64(v), true
	},
	protoreflect.DoubleKind: func(i interface{}) (value protoreflect.Value, valid bool) {
		v, ok := i.(float64)
		if !ok {
			return
		}
		return protoreflect.ValueOfFloat64(v), true
	},

	protoreflect.BytesKind: func(i interface{}) (value protoreflect.Value, valid bool) {
		v, ok := i.([]byte)
		if !ok {
			return
		}
		return protoreflect.ValueOfBytes(v), true
	},
}

// TODO maybe we should not support all of these... floats/doubles?
// TODO we can preallocate a lot of those slices
var protowireFieldEncoders = map[protoreflect.Kind]func(value protoreflect.Value) []byte{
	protoreflect.EnumKind: func(value protoreflect.Value) []byte {
		var b []byte
		b = protowire.AppendVarint(b, uint64(value.Enum()))
		return b
	},
	protoreflect.Sint32Kind: func(value protoreflect.Value) []byte {
		var b []byte
		b = protowire.AppendVarint(b, protowire.EncodeZigZag(int64(int32(value.Int()))))
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
}
