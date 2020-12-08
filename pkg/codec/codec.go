package codec

import "google.golang.org/protobuf/proto"

// Object defines an object that can be marshalled deterministically
type Object interface {
	proto.Message
}

// Name returns the name of the Object
func Name(object Object) string {
	return string(object.ProtoReflect().Descriptor().FullName())
}

// NewCodec is Codec's constructor
func NewCodec() Codec {
	return Codec{}
}

// Codec defines the object serializer and deserializer
type Codec struct {}

func (c Codec) Unmarshal(b []byte, object Object) error {
	return proto.Unmarshal(b, object)
}

func (c Codec) Marshal(object Object) ([]byte, error) {
	return proto.Marshal(object)
}

func (c Codec) Name(o Object) string {
	return Name(o)
}