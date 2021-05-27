package extensions

import (
	v1alpha12 "github.com/fdymylja/tmos/core/coin/v1alpha1"
	"github.com/fdymylja/tmos/runtime/authentication"
	"github.com/fdymylja/tmos/runtime/module"
	"github.com/fdymylja/tmos/x/bank/v1alpha1"
)

func NewAuthentication(c module.Client) module.AuthenticationExtension {
	return Authentication{client: c}
}

type Authentication struct {
	client module.Client
}

func (a Authentication) Initialize(builder *module.AuthenticationExtensionBuilder) {
	builder.
		WithAdmissionController(NewFeeChecker(a.client)).
		WithTransitionController(NewFeeDeduction(a.client))
}

func NewFeeChecker(c module.Client) authentication.AdmissionHandler {
	return FeeChecker{bank: v1alpha1.NewClientSet(c)}
}

// FeeChecker checks if the account has the required amount of fees
type FeeChecker struct {
	bank v1alpha1.ClientSet
}

func (x FeeChecker) Validate(req authentication.Tx) (err error) {
	payer := req.Payer()
	fee := req.Fee()

	// get balance of fee payer
	balance, err := x.bank.Balances().Get(payer)
	if err != nil {
		return err
	}

	// check if it has enough coins
	_, err = v1alpha12.SafeSub(balance.Balance, fee)
	if err != nil {
		return err
	}

	// balance is enough
	return nil
}

func NewFeeDeduction(c module.Client) authentication.PostAuthenticationHandler {
	return FeeDeduction{bank: v1alpha1.NewClientSet(c)}
}

// FeeDeduction deducts fees from the transaction fee payer and sends them to the fee collector
type FeeDeduction struct {
	bank v1alpha1.ClientSet
}

func (x FeeDeduction) Exec(req authentication.PostAuthenticationRequest) (resp authentication.PostAuthenticationResponse, err error) {
	// move coins and send them to fee collector
	return resp, x.bank.ExecMsgSendCoins(&v1alpha1.MsgSendCoins{
		FromAddress: req.Tx.Payer(),
		ToAddress:   "fee_collector",
		Amount:      req.Tx.Fee(),
	})
}
