package bank

import (
	"errors"

	coin "github.com/fdymylja/tmos/module/core/coin/v1alpha1"
	authv1alpha1 "github.com/fdymylja/tmos/module/x/authn/v1alpha1"
	"github.com/fdymylja/tmos/module/x/bank/v1alpha1"
	errors2 "github.com/fdymylja/tmos/runtime/errors"
	"github.com/fdymylja/tmos/runtime/meta"
	"github.com/fdymylja/tmos/runtime/module"
	"github.com/fdymylja/tmos/runtime/statetransition"
)

func NewSendCoinsHandler(client *v1alpha1.Client) SendCoinsHandler {
	return SendCoinsHandler{c: client}
}

type SendCoinsHandler struct {
	c *v1alpha1.Client
}

func (s SendCoinsHandler) Handle(req statetransition.ExecutionRequest) (resp statetransition.ExecutionResponse, err error) {
	msg := req.Transition.(*v1alpha1.MsgSendCoins)

	senderBalance, err := s.c.GetBalance(msg.FromAddress)
	if err != nil {
		return resp, err
	}

	// subtract the coins
	newSenderBalance, err := coin.SafeSub(senderBalance.Balance, msg.Amount)
	if err != nil {
		return resp, err
	}

	// update balance of sender
	err = s.c.SetBalance(msg.FromAddress, newSenderBalance)
	if err != nil {
		return resp, err
	}

	// get balance of receiver
	recvBalance, err := s.c.GetBalance(msg.ToAddress)

	// we do a switch check to assert if the balance exists or not
	switch {
	// if no error simply update the balance
	case err == nil:
		var newRecvBalance []*coin.Coin
		newRecvBalance, err = coin.SafeAdd(recvBalance.Balance, msg.Amount)
		if err != nil {
			return resp, err
		}
		err = s.c.SetBalance(msg.ToAddress, newRecvBalance)
		if err != nil {
			return
		}

		return resp, nil
	// if not found create the balance for the account
	// then attempt to create the account itself if it does not exist
	case errors.Is(err, errors2.ErrNotFound):
		err = s.c.SetBalance(
			msg.ToAddress,
			msg.Amount,
		)
		if err != nil {
			return
		}

		return resp, s.createAccountIfNotExist(msg.ToAddress)
	// another error exit...
	default:
		return resp, err
	}
}

func (s SendCoinsHandler) getBalanceFrom(address string) (*v1alpha1.Balance, error) {
	senderBalance, err := s.c.GetBalance(address)
	if err != nil {
		return nil, err
	}

	return senderBalance, nil
}

// createAccountIfNotExist creates a new account since it has received balance
// so its public key can be sent
// TODO: is this really required?
// TODO: This should be done via Deliver to AUTH.
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
	case errors.Is(err, errors2.ErrNotFound):
		break
	}

	err = s.c.Deliver(&authv1alpha1.MsgCreateAccount{Account: &authv1alpha1.Account{
		Address: address,
	}})
	return err
}

func NewSetCoinsHandler(c module.Client) SetCoinsHandler {
	return SetCoinsHandler{c: c}
}

type SetCoinsHandler struct {
	c module.Client
}

func (s SetCoinsHandler) Handle(req statetransition.ExecutionRequest) (resp statetransition.ExecutionResponse, err error) {
	msg := req.Transition.(*v1alpha1.MsgSetBalance)
	// set balance i guess
	balance := new(v1alpha1.Balance)
	err = s.c.Get(meta.NewStringID(msg.Address), balance)
	switch {
	case err == nil:
		err = s.c.Update(&v1alpha1.Balance{
			Address: msg.Address,
			Balance: msg.Amount,
		})
		return
	case errors.Is(err, errors2.ErrNotFound):
		err = s.c.Create(&v1alpha1.Balance{
			Address: msg.Address,
			Balance: msg.Amount,
		})
		return
	default:
		return
	}
}
