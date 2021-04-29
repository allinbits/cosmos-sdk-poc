package bank

import (
	"errors"

	coin "github.com/fdymylja/tmos/module/core/coin/v1alpha1"
	authv1alpha1 "github.com/fdymylja/tmos/module/x/authn/v1alpha1"
	"github.com/fdymylja/tmos/module/x/bank/v1alpha1"
	"github.com/fdymylja/tmos/runtime"
	"github.com/fdymylja/tmos/runtime/controller"
	"github.com/fdymylja/tmos/runtime/meta"
	"github.com/fdymylja/tmos/runtime/module"
)

func NewSendCoinsHandler(client module.Client) SendCoinsHandler {
	return SendCoinsHandler{c: client}
}

type SendCoinsHandler struct {
	c module.Client
}

func (s SendCoinsHandler) Deliver(req controller.StateTransitionRequest) (resp controller.StateTransitionResponse, err error) {
	msg := req.Transition.(*v1alpha1.MsgSendCoins)
	// get the balance
	senderBalance := new(v1alpha1.Balance)
	err = s.c.Get(meta.NewStringID(msg.FromAddress), senderBalance)
	if err != nil {
		return resp, err
	}
	// subtract the coins
	newSenderBalance, err := coin.SafeSub(senderBalance.Balance, msg.Amount)
	if err != nil {
		return resp, err
	}
	// update balance of sender
	err = s.c.Update(&v1alpha1.Balance{Address: msg.FromAddress, Balance: newSenderBalance})
	if err != nil {
		return resp, err
	}
	// get balance of receiver
	recvBalance := new(v1alpha1.Balance)
	err = s.c.Get(meta.NewStringID(msg.ToAddress), recvBalance)
	// we do a switch check to assert if the balance exists or not
	switch {
	// if no error simply update the balance
	case err == nil:
		var newRecvBalance []*coin.Coin
		newRecvBalance, err = coin.SafeAdd(recvBalance.Balance, msg.Amount)
		if err != nil {
			return resp, err
		}
		err = s.c.Update(&v1alpha1.Balance{
			Address: msg.ToAddress,
			Balance: newRecvBalance,
		})
		if err != nil {
			return
		}
		return resp, nil
	// if not found create the balance for the account
	// then attempt to create the account itself if it does not exist
	case errors.Is(err, runtime.ErrNotFound):
		err = s.c.Create(&v1alpha1.Balance{
			Address: msg.ToAddress,
			Balance: msg.Amount,
		})
		if err != nil {
			return
		}
		return resp, s.createAccountIfNotExist(msg.ToAddress)
	// another error exit...
	default:
		return resp, err
	}
}

// createAccountIfNotExist creates a new account since it has received balance
// so its public key can be sent
// TODO: is this really required?
func (s SendCoinsHandler) createAccountIfNotExist(address string) error {
	acc := new(authv1alpha1.Account)
	err := s.c.Get(meta.NewStringID(address), acc)
	switch {
	// default case it's another error we can't handle
	default:
		return err
	// account exists, exit
	case err == nil:
		return nil
	// break so we create it
	case errors.Is(err, runtime.ErrNotFound):
		break
	}
	err = s.c.Create(&authv1alpha1.Account{Address: address})
	return err
}
