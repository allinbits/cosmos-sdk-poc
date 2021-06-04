package authn

import (
	rbacv1alpha1 "github.com/fdymylja/tmos/core/rbac/v1alpha1"
	"github.com/fdymylja/tmos/runtime/client"
	"github.com/fdymylja/tmos/runtime/statetransition"
	"github.com/fdymylja/tmos/x/authn/v1alpha1"
)

func NewCreateAccountHandler() statetransition.ExecutionHandlerFunc {
	return func(client client.RuntimeClient, req statetransition.ExecutionRequest) (resp statetransition.ExecutionResponse, err error) {
		authClient := v1alpha1.NewClientSet(client)
		rbacClient := rbacv1alpha1.NewClientSet(client)
		msg := req.Transition.(*v1alpha1.MsgCreateAccount)
		// get last acc number
		lastAccNum, err := authClient.CurrentAccountNumber().Get()
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
		err = authClient.Accounts().Create(account)
		if err != nil {
			return resp, err
		}
		// update last account number
		lastAccNum.Number = lastAccNum.Number + 1
		err = authClient.CurrentAccountNumber().Update(lastAccNum)
		if err != nil {
			return
		}
		// bind account to external_account role
		err = rbacClient.ExecMsgBindRole(&rbacv1alpha1.MsgBindRole{
			RoleId:  rbacv1alpha1.ExternalAccountRoleID,
			Subject: account.Address,
		})
		if err != nil {
			return
		}
		return resp, nil
	}
}
