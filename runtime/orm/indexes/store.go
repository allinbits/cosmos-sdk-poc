package indexes

import (
	"fmt"
	"log"

	"github.com/fdymylja/tmos/runtime/kv"
	"github.com/fdymylja/tmos/runtime/meta"
	"github.com/fdymylja/tmos/runtime/orm/schema"
)

const IndexersPrefix = 0x0    // where real indexes are
const PrimaryKeyIndexes = 0x1 // where we save the mapping between primary key -> index keys pointing to the primary key

func NewStore(kv kv.KV, reg *schema.Registry) Store {
	return Store{
		kv:      kv,
		schemas: reg,
	}
}

type Store struct {
	kv      kv.KV
	schemas *schema.Registry
}

func (s *Store) IndexObject(o meta.StateObject) error {
	sch, err := s.schemas.Get(o)
	if err != nil {
		return err
	}
	// skip
	if len(sch.SecondaryKeys) == 0 {
		return nil
	}
	primaryKey := sch.EncodePrimaryKey(o)
	// generate indexes
	indexerKeys := make(indexList, 0, len(sch.SecondaryKeys))
	for indexName := range sch.SecondaryKeys {
		key := &indexerKey{
			objectPrefix: sch.TypePrefix,
			indexName:    []byte(indexName),
			indexValue:   sch.MustEncodeObjectField(indexName, o),
			primaryKey:   primaryKey,
		}
		keyBytes := key.marshal()
		if s.kv.Has(keyBytes) {
			return fmt.Errorf("orm: secondary key already exists")
		}
		s.kv.Set(keyBytes, []byte{})
		indexerKeys = append(indexerKeys, keyBytes)
	}
	// set the type prefixed primary key
	pkToIndexKey := typePrefixedKey{
		primaryKey: primaryKey,
		typePrefix: sch.TypePrefix,
	}
	log.Printf("%s %s %s", pkToIndexKey.primaryKey, pkToIndexKey.typePrefix, pkToIndexKey.bytes())
	s.kv.Set(pkToIndexKey.bytes(), indexerKeys.marshal())
	return nil
}

func (s *Store) UnindexObject(o meta.StateObject) error {
	sch, err := s.schemas.Get(o)
	if err != nil {
		return err
	}
	pk := typePrefixedKey{
		primaryKey: sch.EncodePrimaryKey(o),
		typePrefix: sch.TypePrefix,
	}
	indexListBytes, exists := s.kv.Get(pk.bytes())
	if !exists {
		return fmt.Errorf("orm: primary key not found for object %s", pk.bytes())
	}
	indexes := new(indexList)
	err = indexes.unmarshal(indexListBytes)
	if err != nil {
		return err
	}
	// remove secondary keys
	for _, key := range *indexes {
		exists := s.kv.Delete(key)
		if !exists {
			panic(fmt.Errorf("data corruption, indexes for key %s reported indexer key %s which was not found in indexes list", pk.bytes(), key))
		}
	}
	return nil
}
