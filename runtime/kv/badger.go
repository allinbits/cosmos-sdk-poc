package kv

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

func (b *Badger) Set(k, v []byte) {
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

func (b *Badger) Get(k []byte) (v []byte, exists bool) {
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
	switch {
	case err == nil:
		return v, true
	case errors.Is(err, badger.ErrKeyNotFound):
		return nil, false
	default:
		panic(err)
	}
}

func (b *Badger) Has(k []byte) bool {
	_, exists := b.Get(k)
	return exists
}

func (b *Badger) Delete(k []byte) (exists bool) {
	txn := b.db.NewTransaction(true)
	err := txn.Delete(k)
	if err != nil {
		panic(err)
	}
	return true
}

func (b *Badger) Iterate(start, end []byte) Iterator {
	panic("")
}
