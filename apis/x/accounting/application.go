package accounting

import (
	"github.com/fdymylja/tmos/apis/x/accounting/v1alpha1"
	"github.com/fdymylja/tmos/pkg/controller/basic"
)

// NewModule instantiates a new basic.Controller of the accounting module
func NewModule() basic.Controller {
	return &Module{}
}

type Module struct {
	client basic.Client
}

func (a *Module) Name() string {
	return "bank"
}

func (a *Module) RegisterStateTransitionInterceptor(client basic.Client, register basic.StateTransitionHandler) {

}

func (a *Module) RegisterStateTransitions(client basic.Client, register basic.RegisterTransitionFn) {
	a.client = client
	register(&v1alpha1.MsgSend{}, newMsgSendHandler(client))
}

//
func (a *Module) RegisterStateObjects(register basic.RegisterStateObjectsFn) {
	register(&v1alpha1.Balance{})
}

func (a *Module) Dependencies(interface{}) {

}
