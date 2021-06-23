package client

import (
	"github.com/fdymylja/tmos/core/meta"
	"github.com/fdymylja/tmos/runtime/authentication/user"
	"github.com/fdymylja/tmos/runtime/orm"
)

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

// ----------------------- deliver options -----------------------

// RawDeliverOptions ..
type RawDeliverOptions struct {
	Impersonate user.Users
}

type DeliverOption func(opt *RawDeliverOptions)

// DeliverImpersonating is a client.Exec option which allows the client
// to deliver a meta.StateTransition impersonating another subject(s).
func DeliverImpersonating(subjects ...string) DeliverOption {
	return func(opt *RawDeliverOptions) {
		opt.Impersonate = user.NewUsersFromString(subjects...)
	}
}

// -------------------- update options -----------------------

type UpdateOptionsRaw struct {
	createIfNotExists bool
}

type UpdateOption func(opt *UpdateOptionsRaw)

// UpdateCreateIfNotExists signals during client.Update to create the object in
// the runtime.Runtime store if it does not exist.
func UpdateCreateIfNotExists() UpdateOption {
	return func(opt *UpdateOptionsRaw) {
		opt.createIfNotExists = true
	}
}

// -------------------- create options -------------------

type CreateOption func()

// -------------------- get options ---------------------

// GetOption is used to provide extra options when getting an object
type GetOption func(opt *GetOptionsRaw)

type GetOptionsRaw struct {
	Height uint64
}

func GetAtHeight(height uint64) GetOption {
	return func(opt *GetOptionsRaw) {
		opt.Height = height
	}
}

type DeleteOption func()

// ------------------------- LIST OPTIONS ----------------------------

// ListOptionsRaw is the raw instance of the list options
type ListOptionsRaw struct {
	Height     uint64
	ORMOptions orm.ListOptions
}

type ListOption func(opt *ListOptionsRaw)

// ListAtHeight runs the List operation on the provided height.
// NOTE: this option can not be used by ModuleClient
func ListAtHeight(height uint64) ListOption {
	return func(opt *ListOptionsRaw) {
		return
	}
}

// ListMatchFieldInterface matches the provided field with the given interface value.
// Example:
// in object: Account{Address: "cosmos...", AccountNumber: 56 }
// The object can be matched as ListMatchFieldInterface("accountNumber", 56)
// The store will attempt to convert the given value interface to the concrete type.
func ListMatchFieldInterface(field string, value interface{}) ListOption {
	return func(opt *ListOptionsRaw) {
		opt.ORMOptions.MatchFieldInterface = append(opt.ORMOptions.MatchFieldInterface, orm.ListMatchFieldInterface{
			Field: field,
			Value: value,
		})
	}
}

// ListMatchFieldString matches the provided field with the given string value.
// Example:
// in object: Account{Address: "cosmos...", AccountNumber: 56 }
// The object can be matched as ListMatchFieldString("accountNumber", "56")
// The store will attempt to convert the given value string to the concrete type.
func ListMatchFieldString(field string, value string) ListOption {
	return func(opt *ListOptionsRaw) {
		opt.ORMOptions.MatchFieldString = append(opt.ORMOptions.MatchFieldString, orm.ListMatchFieldString{
			Field: field,
			Value: value,
		})
	}
}

// ListRange instructs the store to provide objects only from start to end range.
func ListRange(start, end uint64) ListOption {
	return func(opt *ListOptionsRaw) {
		opt.ORMOptions.Start = start
		opt.ORMOptions.End = end
	}
}
