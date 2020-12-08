package runtime

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/store/iavl"
	"github.com/fdymylja/cosmos-os/pkg/application"
	"github.com/fdymylja/cosmos-os/pkg/codec"
)

func newAppStore(cdc codec.Codec, db *iavl.Store, id application.ID) application.DB {
	return appStore{cdc: cdc, db: db, prefix: []byte(id)}
}

type appStore struct {
	db *iavl.Store
	prefix []byte
	cdc codec.Codec
}

func (d appStore) Get(key []byte, object codec.Object) error {
	value := d.db.Get(d.key(key))
	if value == nil {
		return fmt.Errorf("not found") // todo change
	}
	return d.cdc.Unmarshal(value, object)
}

func (d appStore) Set(key []byte, object codec.Object) error {
	b, err := d.cdc.Marshal(object)
	if err != nil {
		return fmt.Errorf("marshal error: %w", err) // todo change
	}
	d.db.Set(d.key(key), b)
	return nil
}

func (d appStore) key(key []byte) []byte {
	return append(d.prefix, key...)
}
