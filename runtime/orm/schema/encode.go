package schema

import (
	"fmt"
	"math"

	"google.golang.org/protobuf/encoding/protowire"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func interfaceToValueEncoderForKind(kind protoreflect.Kind) (InterfaceEncoderFunc, error) {
	encoder, exists := safeInterfaceToValue[kind]
	if !exists {
		return nil, fmt.Errorf("protobuf kind %s can not be encoded to bytes", kind)
	}
	return encoder, nil
}

var safeInterfaceToValue = map[protoreflect.Kind]func(i interface{}) (value protoreflect.Value, valid bool){
	protoreflect.BoolKind: func(i interface{}) (value protoreflect.Value, valid bool) {
		v, ok := i.(bool)
		if !ok {
			return
		}
		return protoreflect.ValueOfBool(v), true
	},
	protoreflect.EnumKind: func(i interface{}) (value protoreflect.Value, valid bool) {
		v, ok := i.(int32)
		if !ok {
			return
		}
		return protoreflect.ValueOfEnum(protoreflect.EnumNumber(v)), true
	},
	protoreflect.Int32Kind: func(i interface{}) (value protoreflect.Value, valid bool) {
		v, ok := i.(int32)
		if !ok {
			return
		}
		return protoreflect.ValueOfInt32(v), true
	},
	protoreflect.Sint32Kind: func(i interface{}) (value protoreflect.Value, valid bool) {
		v, ok := i.(int32)
		if !ok {
			return
		}
		return protoreflect.ValueOfInt32(v), true
	},
	protoreflect.Uint32Kind: func(i interface{}) (value protoreflect.Value, valid bool) {
		v, ok := i.(uint32)
		if !ok {
			return
		}
		return protoreflect.ValueOfUint32(v), true
	},
	protoreflect.Int64Kind: func(i interface{}) (value protoreflect.Value, valid bool) {
		v, ok := i.(int64)
		if !ok {
			return
		}
		return protoreflect.ValueOfInt64(v), true
	},
	protoreflect.Sint64Kind: func(i interface{}) (value protoreflect.Value, valid bool) {
		v, ok := i.(int64)
		if !ok {
			return
		}
		return protoreflect.ValueOfInt64(v), true
	},
	protoreflect.Uint64Kind: func(i interface{}) (value protoreflect.Value, valid bool) {
		v, ok := i.(uint64)
		if !ok {
			return
		}
		return protoreflect.ValueOfUint64(v), true
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
	protoreflect.StringKind: func(i interface{}) (value protoreflect.Value, valid bool) {
		v, ok := i.(string)
		if !ok {
			return
		}
		return protoreflect.ValueOfString(v), true
	},
	protoreflect.BytesKind: func(i interface{}) (value protoreflect.Value, valid bool) {
		v, ok := i.([]byte)
		if !ok {
			return
		}
		return protoreflect.ValueOfBytes(v), true
	},
}

func safeFieldEncodeInterface(fd protoreflect.FieldDescriptor, i interface{}) ([]byte, error) {
	if i == nil {
		return nil, fmt.Errorf("%w: for field descriptor %s", ErrEmptyFieldValue, fd.FullName())
	}
	// get the interface to protoreflect.Value function
	toValue, supported := safeInterfaceToValue[fd.Kind()]
	if !supported {
		panic(fmt.Errorf("%w: '%s' in object %s", ErrUnsupportedFieldKind, fd.Kind(), fd.Parent().FullName())) // this should be blocked at GetSchema level
	}

	// convert to value
	value, valid := toValue(i)
	if !valid {
		return nil, fmt.Errorf("%w: protobuf field kind is %s which does not support values of golang type %T", ErrFieldTypeMismatch, fd.Kind(), i)
	}

	// encode to bytes
	return protowireFieldEncoders[fd.Kind()](value), nil
}

// TODO maybe we should not support all of these... floats/doubles?
// TODO we can preallocate a lot of those slices
var protowireFieldEncoders = map[protoreflect.Kind]ValueEncoderFunc{
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
		// NOTE3: this removes the length prefix which is not needed
		return b[1:]
	},
	// NOTE: do we really need to index bytes? when would it be useful?
	protoreflect.BytesKind: func(value protoreflect.Value) []byte {
		var b []byte
		b = protowire.AppendBytes(b, value.Bytes())
		// NOTE: removes the length prefix which is not needed
		return b[1:]
	},
}

func encoderForKind(kind protoreflect.Kind) (ValueEncoderFunc, error) {
	encoder, exists := protowireFieldEncoders[kind]
	if !exists {
		return nil, fmt.Errorf("protobuf kind %s can not be encoded to bytes", kind)
	}
	return encoder, nil
}
