package indexes

import (
	"encoding/binary"
	"fmt"

	"github.com/fdymylja/tmos/runtime/kv"
	"github.com/fdymylja/tmos/runtime/meta"
	"github.com/fdymylja/tmos/runtime/orm/schema"
)

const IndexersPrefix = 0x0    // where real indexes are
const PrimaryKeyIndexes = 0x1 // where we save the mapping between primary key -> index keys pointing to the primary key

type Store struct {
	kv      kv.KV
	schemas map[string]*schema.Schema
}

func (s *Store) IndexObject(o meta.StateObject) error {
	sch, err := s.getSchema(o)
	if err != nil {
		return err
	}
	// skip
	if len(sch.SecondaryKeys) == 0 {
		return nil
	}
	// generate indexes
	indexerKeys := make(indexList, len(sch.SecondaryKeys))
	for indexName := range sch.SecondaryKeys {
		secondaryKey := getSecondaryKey(sch, o, indexName)
		indexingKey, err := getIndexingKey(sch, secondaryKey, o)
		if err != nil {
			return err
		}
		// save index key
		s.kv.Set(indexingKey, []byte{})
		indexerKeys = append(indexerKeys, indexingKey)
	}
	s.kv.Set(sch.EncodePrimaryKey(o), indexerKeys.marshal()) // TODO we can avoid to re-encode primarykey
	return nil
}

func (s *Store) UnindexObject(o meta.StateObject) error {
	sch, err := s.getSchema(o)
	if err != nil {
		return err
	}
	pk := sch.EncodePrimaryKey(o)
	s.kv.Get(pk)
	panic("shee")
}

func (s *Store) getSchema(o meta.StateObject) (*schema.Schema, error) {
	sch, exists := s.schemas[meta.Name(o)]
	if !exists {
		return nil, fmt.Errorf("orm: schema not found for %s", meta.Name(o))
	}
	return sch, nil
}

// getSecondaryKey creates the secondary key bytes given the primary key, the secondary key value
// (so the field value) and the index name (which is given by the json name of the protobuf field)
// the key is then computed in the following way
// <bigEndianLen><[]byte(indexName)><secondaryKeyValue>
// where bigEndianLen = bigEndian(len(indexName) + len(secondaryKeyValue))
//
// key length prefixing is required to avoid iterations over domains not owned by the secondary key.
func getSecondaryKey(s *schema.Schema, o meta.StateObject, indexName string) []byte {
	secondaryKeyValue := s.EncodeSecondaryKey(indexName, o)
	keyLength := len(secondaryKeyValue) + len(indexName)
	// 8 (bigEndian(uint64)) + keyLength
	// initial length is 8 due to how binary.BigEndian saves
	// the big endian bytes of the length
	sk := make([]byte, 8, 8+keyLength)                // 8 is used for key length prefixing
	binary.BigEndian.PutUint64(sk, uint64(keyLength)) // append key length
	sk = append(sk, []byte(indexName)...)             // append index name
	sk = append(sk, secondaryKeyValue...)             // append objects secondary key value

	return sk
}

// getIndexingKey returns the key which maps the primary key of meta.StateObject
// with the provided index.
func getIndexingKey(sch *schema.Schema, secondaryKey []byte, o meta.StateObject) ([]byte, error) {
	primaryKey := sch.EncodePrimaryKey(o)
	if len(primaryKey) == 0 {
		return nil, fmt.Errorf("orm: empty primary key during indexing for object %s", o)
	}
	saveKey := make([]byte, 1+len(sch.Prefix)+len(secondaryKey)+len(primaryKey))
	saveKey = append(saveKey, IndexersPrefix)
	saveKey = append(saveKey, sch.Prefix...)
	saveKey = append(saveKey, secondaryKey...)
	saveKey = append(saveKey, primaryKey...)
	return saveKey, nil
}
