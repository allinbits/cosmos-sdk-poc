package client

import (
	"github.com/fdymylja/tmos/core/meta"
	"github.com/fdymylja/tmos/runtime/authentication/user"
	"github.com/fdymylja/tmos/runtime/orm"
)

// RuntimeServer represents runtime.Runtime as a server
// which other module.Client can interact with.
type RuntimeServer interface {
	Get(id meta.ID, object meta.StateObject) error
	List(object meta.StateObject, options orm.ListOptions) (orm.Iterator, error)
	Create(users user.Users, object meta.StateObject) error
	Update(users user.Users, object meta.StateObject) error
	Delete(users user.Users, object meta.StateObject) error
	Deliver(users user.Users, transition meta.StateTransition) error
}

// ObjectIterator aliases orm.Iterator
type ObjectIterator = orm.Iterator

// RuntimeClient is the client that modules use to interact
// with the RuntimeServer and talk with the store and other modules.
type RuntimeClient interface {
	Get(id meta.ID, object meta.StateObject, opts ...GetOption) error
	List(object meta.StateObject, opts ...ListOption) (ObjectIterator, error)
	Create(object meta.StateObject, opts ...CreateOption) error
	Update(object meta.StateObject, opts ...UpdateOption) error
	Delete(object meta.StateObject, opts ...DeleteOption) error
	Deliver(transition meta.StateTransition, opts ...DeliverOption) error
}
