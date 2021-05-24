package schema

import (
	"fmt"

	"github.com/fdymylja/tmos/runtime/meta"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func NewSecondaryKey(o meta.StateObject, jsonFieldName string) (*SecondaryKey, error) {
	fd := o.ProtoReflect().Descriptor().Fields().ByJSONName(jsonFieldName)
	if fd == nil {
		return nil, fmt.Errorf("%w: json field %s is not present in object %s", ErrBadOptions, jsonFieldName, meta.Name(o))
	}

	valueEncoder, err := encoderForKind(fd.Kind())
	if err != nil {
		return nil, fmt.Errorf("%w: field %s of object %s: %s", ErrBadOptions, jsonFieldName, meta.Name(o), err)
	}
	interfaceEncoderFunc, err := interfaceToValueEncoderForKind(fd.Kind())
	if err != nil {
		return nil, fmt.Errorf("%w field %s of object %s: %s", ErrBadOptions, jsonFieldName, meta.Name(o), err)
	}

	return &SecondaryKey{
		prefix:           []byte(jsonFieldName),
		name:             jsonFieldName,
		encodeValue:      valueEncoder,
		interfaceToValue: interfaceEncoderFunc,
		fd:               fd,
	}, nil
}

type SecondaryKey struct {
	prefix           []byte
	name             string
	encodeValue      ValueEncoderFunc
	interfaceToValue InterfaceEncoderFunc
	fd               protoreflect.FieldDescriptor
}

func (s *SecondaryKey) Prefix() []byte {
	return s.prefix
}

func (s *SecondaryKey) Name() string {
	return s.name
}

func (s *SecondaryKey) Encode(o meta.StateObject) []byte {
	v := o.ProtoReflect().Get(s.fd)
	return s.encodeValue(v)
}

func (s *SecondaryKey) EncodeInterface(v interface{}) ([]byte, error) {
	if v == nil {
		return nil, fmt.Errorf("%w: nil interface provided", ErrFieldTypeMismatch)
	}
	value, ok := s.interfaceToValue(v)
	if !ok {
		return nil, fmt.Errorf("%w: %v of type %T for field descriptor %s of kind %s", ErrFieldTypeMismatch, v, v, s.fd.FullName(), s.fd.Kind())
	}
	return s.encodeValue(value), nil
}
