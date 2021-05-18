package schema

import (
	"fmt"
	"math"

	"github.com/fdymylja/tmos/runtime/meta"
	"google.golang.org/protobuf/encoding/protowire"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type FieldEncoderFunc func(value protoreflect.Value) []byte

type Schema struct {
	Type                 protoreflect.MessageType
	Name                 string
	Prefix               []byte // TODO should we force copies of this?
	PrimaryKey           protoreflect.FieldDescriptor
	PrimaryKeyEncode     FieldEncoderFunc
	SecondaryKeys        map[string]protoreflect.FieldDescriptor
	SecondaryKeyEncoders map[string]FieldEncoderFunc
}

// EncodePrimaryKey returns the encoded primary given a meta.StateObject
// NOTE: panics if the field does not belong to the message
func (s *Schema) EncodePrimaryKey(o meta.StateObject) []byte {
	pkValue := o.ProtoReflect().Get(s.PrimaryKey)
	return s.PrimaryKeyEncode(pkValue)
}

// EncodeSecondaryKey returns the encoded secondary key given a meta.StateObject
// and the json name of the field to encode.
// NOTE: panics if the key provided is not part of the schema.
func (s *Schema) EncodeSecondaryKey(key string, object meta.StateObject) []byte {
	fd, exists := s.SecondaryKeys[key]
	if !exists {
		panic(fmt.Errorf("schema: object %s is not indexed by secondary key %s", meta.Name(object), key))
	}
	encode := s.SecondaryKeyEncoders[key]
	return encode(object.ProtoReflect().Get(fd))
}

type Options struct {
	// PrimaryKey indicates the field to use as a primary key
	// it must be the json name of the protobuf object
	PrimaryKey string
	// SecondaryKeys indicates the protobuf json names of fields
	// of the object to use as secondary keys, the ones that can be
	// passed to Store.List
	SecondaryKeys []string
}

func NewSchema(o meta.StateObject, options Options) (*Schema, error) {
	return getObjectSchema(o, options)
}

func getObjectSchema(o meta.StateObject, options Options) (*Schema, error) {
	schema := &Schema{}
	fds := o.ProtoReflect().Descriptor().Fields()
	primaryKey := fds.ByJSONName(options.PrimaryKey)
	primaryKeyEncoder, err := EncoderForKind(primaryKey.Kind())
	if err != nil {
		return nil, fmt.Errorf("store: %s has invalid primary key field: %w", meta.Name(o), err)
	}
	schema.PrimaryKey = primaryKey
	schema.PrimaryKeyEncode = primaryKeyEncoder

	schema.SecondaryKeys = make(map[string]protoreflect.FieldDescriptor, len(options.SecondaryKeys))
	schema.SecondaryKeyEncoders = make(map[string]FieldEncoderFunc, len(options.SecondaryKeys))
	for _, sk := range options.SecondaryKeys {
		secondaryKey := fds.ByJSONName(sk)
		secondaryKeyEncoder, err := EncoderForKind(secondaryKey.Kind())
		if err != nil {
			return nil, fmt.Errorf("store: %s has invalid secondary key field: %w", meta.Name(o), err)
		}
		schema.SecondaryKeys[sk] = secondaryKey
		schema.SecondaryKeyEncoders[sk] = secondaryKeyEncoder
	}

	schema.Prefix = []byte(meta.Name(o))
	return schema, nil
}

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
