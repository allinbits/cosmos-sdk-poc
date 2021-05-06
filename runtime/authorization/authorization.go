package authorization

import (
	runtimev1alpha1 "github.com/fdymylja/tmos/module/runtime/v1alpha1"
	"github.com/fdymylja/tmos/runtime/meta"
)

// Authorizer defines an authorizer module in the runtime.Runtime
type Authorizer interface {
	// Allowed checks if the provided subject is allowed to do the Verb action on the defined resource
	Allowed(verb runtimev1alpha1.Verb, subject string, resource meta.Type) error
}
