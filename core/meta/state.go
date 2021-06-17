package meta

var SingletonID = BytesID("unique")

func NewStringID(id string) StringID {
	return (StringID)(id)
}

type StringID string

func (s StringID) Bytes() []byte { return []byte(s) }

func NewBytesID(id []byte) BytesID {
	return id
}

type BytesID []byte

func (b BytesID) Bytes() []byte { return b }

// ID defines the unique identification of an StateObject.
type ID interface {
	Bytes() []byte
}
