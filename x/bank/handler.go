package bank

import (
	coin "github.com/fdymylja/tmos/core/coin/v1alpha1"
	"github.com/fdymylja/tmos/core/meta"
	"github.com/fdymylja/tmos/runtime/client"
	"github.com/fdymylja/tmos/runtime/errors"
	"github.com/fdymylja/tmos/runtime/statetransition"
	authnv1alpha1 "github.com/fdymylja/tmos/x/authn/v1alpha1"
	v1alpha12 "github.com/fdymylja/tmos/x/bank/v1alpha1"
)

func NewSendCoinsHandler() SendCoinsHandler {
	return SendCoinsHandler{}
}

type SendCoinsHandler struct{}

func (s SendCoinsHandler) Exec(client client.RuntimeClient, req statetransition.ExecutionRequest) (resp statetransition.ExecutionResponse, err error) {
	bankClient := v1alpha12.NewClientSet(client)
	msg := req.Transition.(*v1alpha12.MsgSendCoins)

	senderBalance, err := bankClient.Balances().Get(msg.FromAddress)
	if err != nil {
		return resp, err
	}

	// subtract the coins
	newSenderBalance, err := coin.SafeSub(senderBalance.Balance, msg.Amount)
	if err != nil {
		return resp, err
	}

	// update balance of sender
	err = bankClient.ExecMsgSetBalance(&v1alpha12.MsgSetBalance{
		Address: msg.FromAddress,
		Amount:  newSenderBalance,
	})

	if err != nil {
		return resp, err
	}

	// get balance of receiver
	recvBalance, err := bankClient.Balances().Get(msg.ToAddress)

	// we do a switch check to assert if the balance exists or not
	switch {
	// if no error simply update the balance
	case err == nil:
		var newRecvBalance []*coin.Coin
		newRecvBalance, err = coin.SafeAdd(recvBalance.Balance, msg.Amount)
		if err != nil {
			return resp, err
		}
		err = bankClient.ExecMsgSetBalance(&v1alpha12.MsgSetBalance{
			Address: msg.ToAddress,
			Amount:  newRecvBalance,
		})
		if err != nil {
			return
		}

		return resp, nil
	// if not found create the balance for the account
	// then attempt to create the account itself if it does not exist
	case errors.Is(err, errors.ErrNotFound):
		err = bankClient.ExecMsgSetBalance(&v1alpha12.MsgSetBalance{
			Address: msg.ToAddress,
			Amount:  msg.Amount,
		})
		if err != nil {
			return
		}

		return resp, s.createAccountIfNotExist(client, msg.ToAddress)
	// another error exit...
	default:
		return resp, err
	}
}

// createAccountIfNotExist creates a new account since it has received balance
// so its public key can be sent
// TODO: is this really required?
// TODO: This should be done via Exec to AUTH.
func (s SendCoinsHandler) createAccountIfNotExist(client client.RuntimeClient, address string) error {
	authClient := authnv1alpha1.NewClientSet(client)
	_, err := authClient.Accounts().Get(address)
	switch {
	// default case it's another error we can't handle
	default:
		return err
	// account exists, exit
	case err == nil:
		return nil
	// break so we create it
	case errors.Is(err, errors.ErrNotFound):
		break
	}
	err = authClient.ExecMsgCreateAccount(&authnv1alpha1.MsgCreateAccount{Account: &authnv1alpha1.Account{
		Address: address,
	}})

	return err
}

func NewSetCoinsHandler() SetCoinsHandler {
	return SetCoinsHandler{}
}

type SetCoinsHandler struct {
}

func (s SetCoinsHandler) Exec(client client.RuntimeClient, req statetransition.ExecutionRequest) (resp statetransition.ExecutionResponse, err error) {
	msg := req.Transition.(*v1alpha12.MsgSetBalance)
	// set balance i guess
	balance := new(v1alpha12.Balance)
	err = client.Get(meta.NewStringID(msg.Address), balance)
	switch {
	case err == nil:
		err = client.Update(&v1alpha12.Balance{
			Address: msg.Address,
			Balance: msg.Amount,
		})
		return
	case errors.Is(err, errors.ErrNotFound):
		err = client.Create(&v1alpha12.Balance{
			Address: msg.Address,
			Balance: msg.Amount,
		})
		return
	default:
		return
	}
}
