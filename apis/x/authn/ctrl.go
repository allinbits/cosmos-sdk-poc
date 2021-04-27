package authn

import (
	"github.com/fdymylja/tmos/apis/x/authn/v1alpha1"
	"github.com/fdymylja/tmos/pkg/controller"
	"github.com/fdymylja/tmos/pkg/module"
)

func NewCreateAccountController(c module.Client) *CreateAccountController {
	return &CreateAccountController{c: c}
}

type CreateAccountController struct {
	c module.Client
}

func (m *CreateAccountController) Deliver(req controller.StateTransitionRequest) (resp controller.StateTransitionResponse, err error) {
	msg := req.Transition.(*v1alpha1.MsgCreateAccount)
	// create account
	err = m.c.Create(msg.Account)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
