package orm

import (
	"encoding/binary"
	"fmt"

	"github.com/fdymylja/tmos/runtime/meta"
	"google.golang.org/protobuf/proto"
)

var marshaler = proto.MarshalOptions{
	Deterministic: true,
}

func (s *Store) Create(o meta.StateObject) error {
	if !s.knownObject(o) {
		return fmt.Errorf("%w: %s", ErrNotRegistered, objectName(o))
	}

	schema := s.schemas[objectName(o)]
	pk, err := getPrimaryKey(schema, o)
	if err != nil {
		return err
	}
	saveKey := getObjectSaveKey(schema, pk)
	// check if object exists
	if s.kv.Has(saveKey) {
		return fmt.Errorf("%w: %s", ErrAlreadyExists, pk)
	}
	// create object
	b, err := marshaler.Marshal(o)
	if err != nil {
		return fmt.Errorf("%w: %s", ErrMarshal, o)
	}
	s.kv.Set(saveKey, b)
	// create indexes
	err = s.indexObject(pk, schema, o)
	if err != nil {
		panic(err)
	}
	return nil
}

// indexObject indexes the given object by saving its primary key in the indexes declared
// in the ObjectSchema, if no index is provided this skips.
func (s *Store) indexObject(pointerPrimaryKey []byte, schema *ObjectSchema, o meta.StateObject) error {
	// skip if there's no indexes
	if len(schema.secondaryKeys) == 0 {
		return nil
	}
	savedIndexes := make([][]byte, len(schema.secondaryKeys))
	for indexName, fieldDescriptor := range schema.secondaryKeys {
		indexValue := o.ProtoReflect().Get(fieldDescriptor)
		fieldEncoder := schema.secondaryKeysEncoder[indexName]

		indexValueBytes := fieldEncoder(indexValue)

		sk := getSecondaryKey(indexValueBytes, indexName)
		saveKey := getIndexObjectSaveKey(pointerPrimaryKey, sk)
		if s.kv.Has(saveKey) {
			panic(fmt.Sprintf("store: unexpected index secondary key found: %x", saveKey))
		}
		s.kv.Set(saveKey, []byte{})
		savedIndexes = append(savedIndexes)
	}
	return nil
}

// getObjectSaveKey returns the key that must be used when saving the object
// given the object schema and the object primary key
func getObjectSaveKey(schema *ObjectSchema, primaryKey []byte) []byte {
	storePrimaryKey := make([]byte, 1+len(schema.prefix)+len(primaryKey))
	storePrimaryKey = append(storePrimaryKey, ObjectsPrefix)    // add objects prefix
	storePrimaryKey = append(storePrimaryKey, schema.prefix...) // add the object namespace prefix
	storePrimaryKey = append(storePrimaryKey, primaryKey...)    // add the object primary key
	return primaryKey
}

// getPrimaryKey returns the primary key bytes of the object given its ObjectSchema
func getPrimaryKey(schema *ObjectSchema, o meta.StateObject) ([]byte, error) {
	value := o.ProtoReflect().Get(schema.primaryKey)
	primaryKey := schema.primaryKeyEncoder(value)
	if len(primaryKey) == 0 {
		return nil, fmt.Errorf("%w: for object %s", ErrEmptyPrimaryKey, o)
	}
	return primaryKey, nil
}

// getSecondaryKey creates the secondary key bytes given the primary key, the secondary key value
// (so the field value) and the index name (which is given by the json name of the protobuf field)
// the key is then computed in the following way
// <bigEndianLen><[]byte(indexName)><secondaryKeyValue><primaryKey>
// where bigEndianLen = bigEndian(len(indexName) + len(secondaryKeyValue))
//
// key length prefixing is required to avoid iterations over domains not owned by the secondary key.
func getSecondaryKey(secondaryKeyValue []byte, indexName string) []byte {
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

// getIndexObjectSaveKey returns the key used to index an object
// given the object primary key and the secondary key produced by getSecondaryKey
func getIndexObjectSaveKey(primaryKey, secondaryKey []byte) []byte {
	saveKey := make([]byte, 1+len(secondaryKey)+len(primaryKey))
	saveKey = append(saveKey, IndexesPrefix)
	saveKey = append(saveKey, secondaryKey...)
	saveKey = append(saveKey, primaryKey...)
	return saveKey
}
