package authn

import (
	"errors"

	"github.com/fdymylja/tmos/module/x/authn/v1alpha1"
	"github.com/fdymylja/tmos/runtime"
	"github.com/fdymylja/tmos/runtime/controller"
	"github.com/fdymylja/tmos/runtime/module"
)

func NewCreateAccountController(c module.Client) *CreateAccountController {
	return &CreateAccountController{c: c}
}

type CreateAccountController struct {
	c module.Client
}

func (m *CreateAccountController) Deliver(req controller.StateTransitionRequest) (resp controller.StateTransitionResponse, err error) {
	msg := req.Transition.(*v1alpha1.MsgCreateAccount)
	// get last account number
	lastAccNum := new(v1alpha1.CurrentAccountNumber)
	err = m.c.Get(v1alpha1.CurrentAccountNumberID, lastAccNum)
	// if it does not exist we create it
	switch {
	case err == nil:
		break
	case errors.Is(err, runtime.ErrNotFound):
		err = m.c.Create(&v1alpha1.CurrentAccountNumber{Number: 0})
		if err != nil {
			return resp, err
		}
	default:
		return resp, err
	}
	// create account
	account := &v1alpha1.Account{
		Address:       msg.Account.Address,
		PubKey:        msg.Account.PubKey,
		AccountNumber: lastAccNum.Number,
		Sequence:      0,
	}
	err = m.c.Create(account)
	if err != nil {
		return resp, err
	}
	// update last account number
	lastAccNum.Number = lastAccNum.Number + 1
	err = m.c.Update(lastAccNum)
	return resp, err
}
