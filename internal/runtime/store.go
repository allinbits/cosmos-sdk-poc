package runtime

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/store/iavl"
	"github.com/fdymylja/cosmos-os/pkg/application"
	"github.com/fdymylja/cosmos-os/pkg/codec"
	"k8s.io/klog/v2"
)

func newAppStore(db *iavl.Store, id application.ID) application.DB {
	return appStore{db: db, prefix: []byte(id)}
}

type appStore struct {
	db     *iavl.Store
	prefix []byte
}

func (d appStore) Get(key []byte, object codec.Object) error {
	if len(key) == 0 {
		panic("empty key")
	}
	key = d.key(key)
	klog.InfoS("querying store", "key", fmt.Sprintf("%s", key), "object", fmt.Sprintf("%#v", object))

	value := d.db.Get(key)
	if value == nil {
		return fmt.Errorf("not found") // todo change
	}
	return codec.Unmarshal(value, object)
}

func (d appStore) Set(key []byte, object codec.Object) error {
	if len(key) == 0 {
		panic("empty key")
	}
	key = d.key(key)
	klog.InfoS("saving to store", "key", fmt.Sprintf("%s", key), "object", fmt.Sprintf("%#v", object))
	b, err := codec.Marshal(object)
	if err != nil {
		return fmt.Errorf("marshal error: %w", err) // todo change
	}
	d.db.Set(key, b)

	return nil
}

func (d appStore) key(key []byte) []byte {
	return append(d.prefix, key...)
}
