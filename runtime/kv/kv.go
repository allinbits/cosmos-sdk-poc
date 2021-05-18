package kv

type KV interface {
	Get(key []byte) (value []byte, exists bool)
	Set(key, value []byte)
	Has(key []byte) (exists bool)
	Delete(key []byte) (exists bool)
	Iterate(start, end []byte) Iterator
	IteratePrefix(prefix []byte) Iterator
}

type Iterator interface {
	Next()
	Key() []byte
	Value() []byte
	Valid() bool
	Close()
}
