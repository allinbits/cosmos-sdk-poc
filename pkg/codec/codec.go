package codec

import "google.golang.org/protobuf/proto"

type Object interface {
	proto.Message
}

// DeterministicObject defines an object that can be marshalled deterministically
type DeterministicObject interface {
	proto.Message
}
