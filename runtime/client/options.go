package client

import (
	"github.com/fdymylja/tmos/runtime/authentication/user"
	"github.com/fdymylja/tmos/runtime/orm"
)

// ----------------------- deliver options -----------------------
type deliverOptions struct {
	impersonate user.Users
}

type DeliverOption func(opt *deliverOptions)

// DeliverImpersonating is a client.Exec option which allows the client
// to deliver a meta.StateTransition impersonating another subject(s).
func DeliverImpersonating(subjects ...string) DeliverOption {
	return func(opt *deliverOptions) {
		opt.impersonate = user.NewUsersFromString(subjects...)
	}
}

// -------------------- update options -----------------------
type updateOptions struct {
	createIfNotExists bool
}

type UpdateOption func(opt *updateOptions)

// UpdateCreateIfNotExists signals during client.Update to create the object in
// the runtime.Runtime store if it does not exist.
func UpdateCreateIfNotExists() UpdateOption {
	return func(opt *updateOptions) {
		opt.createIfNotExists = true
	}
}

// -------------------- create options -------------------

type CreateOption func()

// -------------------- get options ---------------------

// GetOption is used to provide extra options when getting an object
type GetOption func(opt *getOptions)

type getOptions struct {
	Height uint64
}

func GetAtHeight(height uint64) GetOption {
	return func(opt *getOptions) {
		opt.Height = height
	}
}

type DeleteOption func()

// ------------------------- LIST OPTIONS ----------------------------

// listOptions is the raw instance of the list options
type listOptions struct {
	Height     uint64
	ORMOptions orm.ListOptions
}

type ListOption func(opt *listOptions)

// ListAtHeight runs the List operation on the provided height.
// NOTE: this option can not be used by ModuleClient
func ListAtHeight(height uint64) ListOption {
	return func(opt *listOptions) {
		return
	}
}

// ListMatchFieldInterface matches the provided field with the given interface value.
// Example:
// in object: Account{Address: "cosmos...", AccountNumber: 56 }
// The object can be matched as ListMatchFieldInterface("accountNumber", 56)
// The store will attempt to convert the given value interface to the concrete type.
func ListMatchFieldInterface(field string, value interface{}) ListOption {
	return func(opt *listOptions) {
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
	return func(opt *listOptions) {
		opt.ORMOptions.MatchFieldString = append(opt.ORMOptions.MatchFieldString, orm.ListMatchFieldString{
			Field: field,
			Value: value,
		})
	}
}
