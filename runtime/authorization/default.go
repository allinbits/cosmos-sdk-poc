package authorization

import (
	"fmt"

	meta "github.com/fdymylja/tmos/core/meta"
)

var _ Authorizer = AlwaysAllowAuthorizer{}
var _ Authorizer = AlwaysDenyAuthorizer{}

func NewAlwaysAllowAuthorizer() AlwaysAllowAuthorizer {
	return AlwaysAllowAuthorizer{}
}

type AlwaysAllowAuthorizer struct{}

func (n AlwaysAllowAuthorizer) Authorize(_ Attributes) (Decision, error) {
	return DecisionAllow, nil
}

func NewAlwaysDenyAuthorizer() AlwaysDenyAuthorizer {
	return AlwaysDenyAuthorizer{}
}

type AlwaysDenyAuthorizer struct{}

func (a AlwaysDenyAuthorizer) Authorize(attributes Attributes) (decision Decision, err error) {
	return DecisionDeny, fmt.Errorf("no user in %s has role %s towards resource %s", attributes.Users.String(), attributes.Verb, meta.Name(attributes.Resource))
}
