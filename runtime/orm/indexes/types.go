package indexes

import (
	"fmt"

	"google.golang.org/protobuf/encoding/protowire"
)

type indexerKey struct {
	objectPrefix []byte
	indexName    []byte
	indexValue   []byte
	primaryKey   []byte
}

// marshal marshals indexerKey such as it provides prefix protection
// it prefixes everything (except primary key) with its length
// in the form of binary.LittleEndian.
func (s indexerKey) marshal() []byte {
	// get lengths
	objectPrefixL := len(s.objectPrefix)
	indexNameL := len(s.indexName)
	indexValueL := len(s.indexValue)
	// create pre-allocated buffer
	buf := make([]byte, 1, 1+8+objectPrefixL+8+indexNameL+8+indexValueL+len(s.primaryKey))
	buf[0] = IndexersPrefix
	buf = appendLengthPrefixed(buf, s.objectPrefix) // append object prefix with length
	buf = appendLengthPrefixed(buf, s.indexName)    // append index name with length
	buf = appendLengthPrefixed(buf, s.indexValue)   // append index value with length
	buf = append(buf, s.primaryKey...)
	return buf
}

func (s *indexerKey) unmarshal(buf []byte) error {
	// we exclude the first byte
	buf = buf[1:]
	// we get the object prefix
	objPrefix, read1, err := consumeLengthPrefixed(buf)
	if err != nil {
		return err
	}
	s.objectPrefix = objPrefix
	// we get the index name
	indexName, read2, err := consumeLengthPrefixed(buf[read1:])
	if err != nil {
		return err
	}
	s.indexName = indexName
	// we get the index value
	indexValue, read3, err := consumeLengthPrefixed(buf[read1+read2:])
	if err != nil {
		return err
	}
	s.indexValue = indexValue
	s.primaryKey = buf[read1+read2+read3:]
	return nil
}

func consumeLengthPrefixed(buf []byte) ([]byte, int, error) {
	if len(buf) < 8 {
		return nil, 0, fmt.Errorf("orm: buffer needs at least 8 bytes for length")
	}
	read := 8
	length := uint64(buf[0])<<0 | uint64(buf[1])<<8 | uint64(buf[2])<<16 | uint64(buf[3])<<24 | uint64(buf[4])<<32 | uint64(buf[5])<<40 | uint64(buf[6])<<48 | uint64(buf[7])<<56
	if len(buf) < 8+int(length) {
		return nil, 0, fmt.Errorf("orm: buffer needs to be at least %d bytes", 8+length)
	}
	value := make([]byte, length)
	for i := 0; i < int(length); i++ {
		value[i] = buf[read]
		read++
	}
	return value, read, nil
}

func appendLengthPrefixed(buf []byte, value []byte) []byte {
	// the size of the buffer must be 8 + len(value)
	v := (uint64)(len(value))
	// append length
	buf = append(buf,
		byte(v>>0),
		byte(v>>8),
		byte(v>>16),
		byte(v>>24),
		byte(v>>32),
		byte(v>>40),
		byte(v>>48),
		byte(v>>56))
	// append value
	buf = append(buf, value...)
	return buf
}

// indexList marshals itself into an index list
type indexList [][]byte

func (i indexList) marshal() []byte {
	if len(i) == 0 {
		return nil
	}
	buf := []byte{PrimaryKeyIndexes}
	buf = protowire.AppendVarint(buf, (uint64)(len(i)))
	for _, index := range i {
		buf = protowire.AppendBytes(buf, index)
	}
	return buf
}

func (i *indexList) unmarshal(buf []byte) error {
	buf = buf[1:]
	varint, code := protowire.ConsumeVarint(buf)
	if code < 0 {
		return protowire.ParseError(code)
	}
	buf = buf[code:] // skip varint
	numIndexes := (int)(varint)
	list := make([][]byte, numIndexes)
	for currentIndex := 0; currentIndex < numIndexes; currentIndex++ {
		if len(buf) == 0 {
			return nil
		}
		index, code := protowire.ConsumeBytes(buf)
		if code < 0 {
			return protowire.ParseError(code)
		}
		list[currentIndex] = index
		buf = buf[code:]
	}
	if len(buf) != 0 {
		return fmt.Errorf("orm: index buffer should have been fully consumed got %d bytes left: %v", len(buf), buf)
	}
	*i = list
	return nil
}

type typePrefixedKey struct {
	primaryKey []byte
	typePrefix []byte
}

func (k typePrefixedKey) bytes() []byte {
	key := make([]byte, 0, len(k.typePrefix)+1+len(k.primaryKey))
	key = append(key, k.typePrefix...)
	key = append(key, '/')
	key = append(key, k.primaryKey...)
	return key
}
