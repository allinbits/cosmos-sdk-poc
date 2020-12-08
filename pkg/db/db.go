package db

import (
	"github.com/cosmos/cosmos-sdk/store/iavl"
	"github.com/fdymylja/cosmos-os/pkg/codec"
)

type Prefixed struct {
	prefix []byte
	db *iavl.Store
}

func (p Prefixed) Get(key []byte, object codec.Object) error {
	panic("implement me")
}

func (p Prefixed) Set(key []byte, object codec.Object) error {
	panic("implement me")
}

