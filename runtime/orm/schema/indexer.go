package schema

import (
	"fmt"

	"github.com/fdymylja/tmos/core/meta"
	"github.com/fdymylja/tmos/pkg/protoutils/kindencoder"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func NewIndexer(o meta.StateObject, jsonFieldName string) (*Indexer, error) {
	fd := o.ProtoReflect().Descriptor().Fields().ByJSONName(jsonFieldName)
	if fd == nil {
		return nil, fmt.Errorf("%w: json field %s is not present in object %s", ErrBadDefinition, jsonFieldName, meta.Name(o))
	}

	kindEncoder, err := kindencoder.NewKindEncoder(fd.Kind())
	if err != nil {
		return nil, fmt.Errorf("%w: field %s of object %s: %s", ErrBadDefinition, jsonFieldName, meta.Name(o), err)
	}

	return &Indexer{
		prefix:      []byte(jsonFieldName),
		name:        jsonFieldName,
		kindEncoder: kindEncoder,
		fd:          fd,
	}, nil
}

type Indexer struct {
	prefix      []byte
	name        string
	kindEncoder kindencoder.KindEncoder
	fd          protoreflect.FieldDescriptor
}

func (s *Indexer) Prefix() []byte {
	return s.prefix
}

func (s *Indexer) Name() string {
	return s.name
}

func (s *Indexer) Encode(o meta.StateObject) []byte {
	v := o.ProtoReflect().Get(s.fd)
	return s.kindEncoder.EncodeBytes(v)
}

func (s *Indexer) EncodeInterface(v interface{}) ([]byte, error) {
	value, err := s.kindEncoder.EncodeInterface(v)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrFieldTypeMismatch, err)
	}
	return s.kindEncoder.EncodeBytes(value), nil
}
