package authn

import (
	rbacv1alpha1 "github.com/fdymylja/tmos/core/rbac/v1alpha1"
	"github.com/fdymylja/tmos/runtime/meta"
	"github.com/fdymylja/tmos/runtime/module"
	"github.com/fdymylja/tmos/runtime/statetransition"
	v1alpha12 "github.com/fdymylja/tmos/x/authn/v1alpha1"
)

func NewCreateAccountController(c module.Client) *CreateAccountController {
	return &CreateAccountController{c: c}
}

type CreateAccountController struct {
	c module.Client
}

func (m *CreateAccountController) Exec(req statetransition.ExecutionRequest) (resp statetransition.ExecutionResponse, err error) {
	msg := req.Transition.(*v1alpha12.MsgCreateAccount)
	// get last account number
	lastAccNum := new(v1alpha12.CurrentAccountNumber)
	err = m.c.Get(meta.SingletonID, lastAccNum)
	if err != nil {
		return statetransition.ExecutionResponse{}, err
	}
	// create account
	account := &v1alpha12.Account{
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
