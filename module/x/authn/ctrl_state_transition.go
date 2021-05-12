package authn

import (
	rbacv1alpha1 "github.com/fdymylja/tmos/module/rbac/v1alpha1"
	"github.com/fdymylja/tmos/module/x/authn/v1alpha1"
	"github.com/fdymylja/tmos/runtime/module"
	"github.com/fdymylja/tmos/runtime/statetransition"
)

func NewCreateAccountController(c module.Client) *CreateAccountController {
	return &CreateAccountController{c: c}
}

type CreateAccountController struct {
	c module.Client
}

func (m *CreateAccountController) Handle(req statetransition.Request) (resp statetransition.Response, err error) {
	msg := req.Transition.(*v1alpha1.MsgCreateAccount)
	// get last account number
	lastAccNum := new(v1alpha1.CurrentAccountNumber)
	err = m.c.Get(v1alpha1.CurrentAccountNumberID, lastAccNum)
	if err != nil {
		return statetransition.Response{}, err
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
	if err != nil {
		return
	}
	// bind account to external_account role
	err = m.c.Deliver(&rbacv1alpha1.MsgBindRole{
		RoleId:  rbacv1alpha1.ExternalAccountRoleID,
		Subject: account.Address,
	})
	if err != nil {
		return
	}
	return resp, nil
}
