package indexes

import (
	"fmt"

	"github.com/fdymylja/tmos/runtime/kv"
	"github.com/fdymylja/tmos/runtime/orm"
	"github.com/fdymylja/tmos/runtime/orm/schema"
)

func (s Store) List(sch *schema.Schema, options orm.ListOptions) (kv.Iterator, error) {
	prefixes, err := prefixesToScan(sch, options)
	if err != nil {
		return nil, err
	}
	// TODO we need to add the logic of query ranges and joint prefixes
	if len(prefixes) == 0 {
		return nil, fmt.Errorf("TODO remove me and implement me")
	}
	iterator := s.kv.IteratePrefix(prefixes[0].marshal())
	if !iterator.Valid() {
		return nil, fmt.Errorf("%w: no records found for object %s and query %s", errNotFound, sch.Name(), options)
	}
	return iterator, nil
}

// prefixesToScan returns the list of prefixes that need to be scanned to fetch primary keys
// associated with the provided fields
func prefixesToScan(schema *schema.Schema, opt orm.ListOptions) ([]indexObjectWithSecondaryKey, error) {
	prefixes := make([]indexObjectWithSecondaryKey, 0, len(opt.MatchFieldString)+len(opt.MatchFieldString))
	for _, o := range opt.MatchFieldInterface {
		idx, err := schema.Indexer(o.Field)
		if err != nil {
			return nil, fmt.Errorf("object %s does not have field %s", schema.Name(), o.Field)
		}
		b, err := idx.EncodeInterface(o.Value)
		if err != nil {
			return nil, fmt.Errorf("unable to encode field %s of value %#v of type %T for object %s", o.Field, o.Value, o.Value, schema.Name())
		}
		prefixes = append(prefixes, indexObjectWithSecondaryKey{
			objectPrefix:      schema.TypePrefix(),
			indexPrefix:       idx.Prefix(),
			secondaryKeyValue: b,
		})
	}
	// TODO maybe we should return an error in case the user is trying to query two indexes with the same field and value?
	for _, o := range opt.MatchFieldString {
		idx, err := schema.Indexer(o.Field)
		if err != nil {
			return nil, fmt.Errorf("object %s does not have field %s", schema.Name(), o.Field)
		}
		b, err := idx.EncodeString(o.Value)
		if err != nil {
			return nil, fmt.Errorf("unable to encode field %s of value %#v from string for object %s", o.Field, o.Value, schema.Name())
		}
		prefixes = append(prefixes, indexObjectWithSecondaryKey{
			objectPrefix:      schema.TypePrefix(),
			indexPrefix:       idx.Prefix(),
			secondaryKeyValue: b,
		})
	}

	return prefixes, nil
}
