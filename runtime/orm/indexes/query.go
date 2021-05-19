package indexes

import (
	"fmt"

	"github.com/fdymylja/tmos/runtime/kv"
	"github.com/fdymylja/tmos/runtime/orm/schema"
)

type ListOptionFunc func(opt *listOptions)

func MatchField(fieldName string, value interface{}) ListOptionFunc {
	return func(opt *listOptions) {
		opt.matchingFields = append(opt.matchingFields, matchField{
			fieldName: fieldName,
			value:     value,
		})
	}
}

type listOptions struct {
	matchingFields []matchField
}

type matchField struct {
	fieldName string
	value     interface{}
}

func (s Store) List(sch *schema.Schema, options ...ListOptionFunc) (kv.Iterator, error) {
	o := new(listOptions)
	for _, opt := range options {
		opt(o)
	}

	match := o.matchingFields[0]
	skValue, err := sch.EncodeFieldInterface(match.fieldName, match.value)
	if err != nil {
		return nil, err
	}
	prefix := indexerKey{
		objectPrefix: sch.TypePrefix,
		indexName:    []byte(match.fieldName),
		indexValue:   skValue,
		primaryKey:   nil,
	}
	iterator := s.kv.IteratePrefix(prefix.marshal())
	if !iterator.Valid() {
		return nil, fmt.Errorf("%w: no records found for object %s and query %s", ErrNotFound, sch.Name, o)
	}
	return iterator, nil
}
