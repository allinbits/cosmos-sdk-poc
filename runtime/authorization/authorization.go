package authorization

import (
	"github.com/fdymylja/tmos/runtime"
)

// Verb defines an action that can be performed on the runtime
type Verb int8

const (
	GET Verb = iota
	LIST
	CREATE
	UPDATE
	DELETE
	DELIVER
)

// Authorizer defines an authorizer module in the runtime.Runtime
type Authorizer interface {
	// Allowed checks if the provided subject is allowed to do the Verb action on the defined resource
	Allowed(verb Verb, subject string, resource runtime.Type) bool
}
