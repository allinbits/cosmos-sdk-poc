package client

import "github.com/fdymylja/tmos/runtime/authentication/user"

type deliverOptions struct {
	impersonate user.Users
}

type DeliverOption func(opt *deliverOptions)

// DeliverImpersonating is a client.Deliver option which allows the client
// to deliver a meta.StateTransition impersonating another subject(s).
func DeliverImpersonating(subjects ...string) DeliverOption {
	return func(opt *deliverOptions) {
		opt.impersonate = user.NewUsersFromString(subjects...)
	}
}

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

type CreateOption func()
type GetOption func()
type DeleteOption func()
