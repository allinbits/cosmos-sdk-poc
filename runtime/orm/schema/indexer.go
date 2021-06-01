package schema

import (
	"fmt"

	"github.com/fdymylja/tmos/core/meta"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func NewIndexer(o meta.StateObject, jsonFieldName string) (*Indexer, error) {
	fd := o.ProtoReflect().Descriptor().Fields().ByJSONName(jsonFieldName)
	if fd == nil {
		return nil, fmt.Errorf("%w: json field %s is not present in object %s", ErrBadDefinition, jsonFieldName, meta.Name(o))
	}

	valueEncoder, err := encoderForKind(fd.Kind())
	if err != nil {
		return nil, fmt.Errorf("%w: field %s of object %s: %s", ErrBadDefinition, jsonFieldName, meta.Name(o), err)
	}
	interfaceEncoderFunc, err := interfaceToValueEncoderForKind(fd.Kind())
	if err != nil {
		return nil, fmt.Errorf("%w field %s of object %s: %s", ErrBadDefinition, jsonFieldName, meta.Name(o), err)
	}

	return &Indexer{
		prefix:           []byte(jsonFieldName),
		name:             jsonFieldName,
		encodeValue:      valueEncoder,
		interfaceToValue: interfaceEncoderFunc,
		fd:               fd,
	}, nil
}

type Indexer struct {
	prefix           []byte
	name             string
	encodeValue      ValueEncoderFunc
	interfaceToValue InterfaceEncoderFunc
	fd               protoreflect.FieldDescriptor
}

func (s *Indexer) Prefix() []byte {
	return s.prefix
}

func (s *Indexer) Name() string {
	return s.name
}

func (s *Indexer) Encode(o meta.StateObject) []byte {
	v := o.ProtoReflect().Get(s.fd)
	return s.encodeValue(v)
}

func (s *Indexer) EncodeInterface(v interface{}) ([]byte, error) {
	if v == nil {
		return nil, fmt.Errorf("%w: nil interface provided", ErrFieldTypeMismatch)
	}
	value, ok := s.interfaceToValue(v)
	if !ok {
		return nil, fmt.Errorf("%w: %v of type %T for field descriptor %s of kind %s", ErrFieldTypeMismatch, v, v, s.fd.FullName(), s.fd.Kind())
	}
	return s.encodeValue(value), nil
}
