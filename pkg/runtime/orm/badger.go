package orm

import (
	"errors"

	"github.com/dgraph-io/badger/v3"
)

type Badger struct {
	db *badger.DB
}

func NewBadger() *Badger {
	db, err := badger.Open(badger.DefaultOptions("").WithInMemory(true))
	if err != nil {
		panic(err)
	}
	return &Badger{db: db}
}

func (b *Badger) set(k, v []byte) {
	txn := b.db.NewTransaction(true)
	err := txn.Set(k, v)
	if err != nil {
		panic(err)
	}
	err = txn.Commit()
	if err != nil {
		panic(err)
	}
	txn.Discard()
}

func (b *Badger) get(k []byte) (v []byte, exists bool) {
	err := b.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get(k)
		if err != nil {
			return err
		}
		err = item.Value(func(val []byte) error {
			v = val
			return nil
		})
		if err != nil {
			return err
		}
		return nil
	})

	if errors.Is(err, badger.ErrKeyNotFound) {
		return nil, false
	}
	return v, true
}

func (b *Badger) delete(k []byte) {
	txn := b.db.NewTransaction(true)
	err := txn.Delete(k)
	if err != nil {
		panic(err)
	}
}
