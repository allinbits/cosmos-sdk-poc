package authn

import (
	meta "github.com/fdymylja/tmos/apis/meta/v1alpha1"
	"github.com/fdymylja/tmos/apis/x/authn/v1alpha1"
	"github.com/fdymylja/tmos/pkg/controller"
	"github.com/fdymylja/tmos/pkg/module"
	"github.com/fdymylja/tmos/pkg/runtime/errors"
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
	err = m.c.Get(lastAccNum)
	// if it does not exist we create it
	if errors.IsNotFound(err) {
		err = m.c.Create(&v1alpha1.CurrentAccountNumber{Number: 0})
		if err != nil {
			return resp, err
		}
	}
	// create account
	account := &v1alpha1.Account{
		Address:       msg.Account.Address,
		PubKey:        msg.Account.PubKey,
		AccountNumber: lastAccNum.Number,
		Sequence:      0,
		ObjectMeta:    &meta.ObjectMeta{Id: msg.Account.Address},
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
