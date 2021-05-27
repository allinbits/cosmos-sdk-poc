package authn

import (
	rbacv1alpha1 "github.com/fdymylja/tmos/core/rbac/v1alpha1"
	"github.com/fdymylja/tmos/runtime/module"
	"github.com/fdymylja/tmos/runtime/statetransition"
	"github.com/fdymylja/tmos/x/authn/v1alpha1"
)

func NewCreateAccountController(c module.Client) *CreateAccountController {
	return &CreateAccountController{
		auth: v1alpha1.NewClientSet(c),
		rbac: rbacv1alpha1.NewClientSet(c),
	}
}

type CreateAccountController struct {
	auth v1alpha1.ClientSet
	rbac rbacv1alpha1.ClientSet
}

func (m *CreateAccountController) Exec(req statetransition.ExecutionRequest) (resp statetransition.ExecutionResponse, err error) {
	msg := req.Transition.(*v1alpha1.MsgCreateAccount)
	// get last acc number
	lastAccNum, err := m.auth.CurrentAccountNumber().Get()
	if err != nil {
		return statetransition.ExecutionResponse{}, err
	}
	// create account
	account := &v1alpha1.Account{
		Address:       msg.Account.Address,
		PubKey:        msg.Account.PubKey,
		AccountNumber: lastAccNum.Number,
		Sequence:      0,
	}
	err = m.auth.Accounts().Create(account)
	if err != nil {
		return resp, err
	}
	// update last account number
	lastAccNum.Number = lastAccNum.Number + 1
	err = m.auth.CurrentAccountNumber().Update(lastAccNum)
	if err != nil {
		return
	}
	// bind account to external_account role
	err = m.rbac.ExecMsgBindRole(&rbacv1alpha1.MsgBindRole{
		RoleId:  rbacv1alpha1.ExternalAccountRoleID,
		Subject: account.Address,
	})
	if err != nil {
		return
	}
	return resp, nil
}
