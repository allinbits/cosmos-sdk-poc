package indexes

import (
	"fmt"

	"github.com/fdymylja/tmos/runtime/kv"
	"github.com/fdymylja/tmos/runtime/orm"
	"github.com/fdymylja/tmos/runtime/orm/schema"
)

func (s Store) List(sch *schema.Schema, options orm.ListOptionsRaw) (kv.Iterator, error) {

	match := options.MatchFields[0]
	skValue, err := sch.EncodeFieldInterface(match.Field, match.Value)
	if err != nil {
		return nil, err
	}
	prefix := indexerKey{
		objectPrefix: sch.TypePrefix,
		indexName:    []byte(match.Field),
		indexValue:   skValue,
		primaryKey:   nil,
	}
	iterator := s.kv.IteratePrefix(prefix.marshal())
	if !iterator.Valid() {
		return nil, fmt.Errorf("%w: no records found for object %s and query %s", ErrNotFound, sch.Name, options)
	}
	return iterator, nil
}
