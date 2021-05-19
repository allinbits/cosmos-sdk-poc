package kv

import (
	"bytes"
	"errors"
	"fmt"

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
	err = txn.Commit()
	if err != nil {
		panic(err)
	}
	return true
}

func (b *Badger) Iterate(start, end []byte) Iterator {
	txn := b.db.NewTransaction(true)
	iter := txn.NewKeyIterator(start, badger.IteratorOptions{
		PrefetchSize:   1,
		PrefetchValues: false,
		Reverse:        false,
		AllVersions:    false,
		InternalAccess: false,
	})
	iter.Rewind()
	return BadgerIterator{
		start: start,
		end:   end,
		iter:  iter,
	}
}

func (b *Badger) IteratePrefix(prefix []byte) Iterator {
	txn := b.db.NewTransaction(true)
	iter := txn.NewIterator(badger.DefaultIteratorOptions)
	iter.Seek(prefix)
	return BaderPrefixIterator{
		iter:      iter,
		prefix:    prefix,
		pfxLength: len(prefix),
	}
}

type BaderPrefixIterator struct {
	iter      *badger.Iterator
	prefix    []byte
	pfxLength int
}

func (b BaderPrefixIterator) Next() {
	if !b.Valid() {
		panic(fmt.Errorf("kv: iterator is not valid"))
	}
	b.iter.Next()
}

func (b BaderPrefixIterator) Key() []byte {
	return b.iter.Item().Key()[b.pfxLength:]
}

func (b BaderPrefixIterator) Value() []byte {
	v, err := b.iter.Item().ValueCopy(nil)
	if err != nil {
		panic(err)
	}
	return v
}

func (b BaderPrefixIterator) Valid() bool {
	return b.iter.ValidForPrefix(b.prefix)
}

func (b BaderPrefixIterator) Close() {
	b.iter.Close()
}

type BadgerIterator struct {
	start []byte
	end   []byte
	iter  *badger.Iterator
}

func (b BadgerIterator) Next() {
	if !b.Valid() {
		panic(fmt.Errorf("kv: iterator is not valid"))
	}
	b.iter.Next()
}

func (b BadgerIterator) Key() []byte {
	return b.iter.Item().Key()
}

func (b BadgerIterator) Value() []byte {
	v, err := b.iter.Item().ValueCopy(nil) // TODO maybe don't copy
	if err != nil {
		panic(err)
	}
	return v
}

func (b BadgerIterator) Valid() bool {
	if !b.iter.Valid() {
		return false
	}
	if len(b.end) != 0 {
		if bytes.Compare(b.iter.Item().Key(), b.end) == 1 {
			return false
		}
	}
	return true
}

func (b BadgerIterator) Close() {
	b.iter.Close()
}
