package codec

import (
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

// Object defines an object that can be marshalled deterministically
type Object interface {
	proto.Message
}

// Name returns the name of the Object
func Name(object Object) string {
	name, err := anypb.New(object)
	if err != nil {
		panic("unable to marshal name " + err.Error())
	}
	return name.TypeUrl
}

func Unmarshal(b []byte, object Object) error {
	return proto.Unmarshal(b, object)
}

func Marshal(object Object) ([]byte, error) {
	return proto.Marshal(object)
}

func NewTx(msg Object) (*anypb.Any, error) {
	return anypb.New(msg)
}
