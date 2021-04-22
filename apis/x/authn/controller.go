package authn

import (
	"github.com/fdymylja/tmos/apis/meta"
	"github.com/fdymylja/tmos/pkg/controller/authentication"
	"github.com/fdymylja/tmos/pkg/controller/basic"
)

// Controller implements the authentication.Controller
type Controller struct {
}

func NewController() authentication.Controller {
	return Controller{}
}

func (c Controller) Name() string {
	panic("implement me")
}

func (c Controller) RegisterStateTransitions(client basic.Client, register basic.RegisterTransitionFn) {
	panic("implement me")
}

func (c Controller) RegisterStateObjects(register basic.RegisterStateObjectsFn) {
	panic("implement me")
}

func (c Controller) DecodeTx(txBytes []byte) (transitions []meta.StateTransition, authenticatedAccounts []string, err error) {
	panic("implement me")
}
