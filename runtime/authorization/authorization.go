package authorization

import (
	meta "github.com/fdymylja/tmos/core/meta"
	runtimev1alpha1 "github.com/fdymylja/tmos/core/runtime/v1alpha1"
	"github.com/fdymylja/tmos/runtime/authentication/user"
)

// Decision defines the Authorizer choice regarding the request
type Decision int

const (
	DecisionDeny Decision = iota
	DecisionAllow
)

// Attributes contains
type Attributes struct {
	// Verb returns the runtime verb associated with the request
	Verb runtimev1alpha1.Verb
	// Resource contains the resource name being accessed
	Resource meta.Type
	// Users contains the information of the users who made the request
	Users user.Users
}

// Authorizer makes an authorization decision by inspecting the Attributes
type Authorizer interface {
	Authorize(attributes Attributes) (decision Decision, err error)
}
