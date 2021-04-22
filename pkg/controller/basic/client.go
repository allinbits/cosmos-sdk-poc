package basic

import "github.com/fdymylja/tmos/apis/meta"

// Client defines a client that applications can use to interact with one another
type Client interface {
	// Get gets and unmarshals bytes to the given meta.StateObject
	// returns false if it does not exist
	Get(object meta.StateObject) (exists bool)
	// Set sets the object in the database, panics if the
	// caller does not have the required authorizations
	Set(object meta.StateObject)
	// Delete deletes the given object from the state
	// panics if the caller does not have the required
	// authorizations to delete the object
	Delete(object meta.StateObject)
	// Deliver delivers the state transition to the target module
	// panics if the state transition has no handler
	// returns an error if the Deliver response is not successful
	Deliver(transition meta.StateTransition) error
}
