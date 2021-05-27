package runtime

import (
	"github.com/fdymylja/tmos/runtime/orm"
	"github.com/fdymylja/tmos/runtime/orm/schema"
)

type Querier struct {
	reg   *schema.Registry
	store orm.Store
}

func NewQuerier(store orm.Store) {

}
