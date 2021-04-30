package extensions

import (
	coin "github.com/fdymylja/tmos/module/core/coin/v1alpha1"
	bankv1aplha1 "github.com/fdymylja/tmos/module/x/bank/v1alpha1"
	"github.com/fdymylja/tmos/runtime/authentication"
	"github.com/fdymylja/tmos/runtime/module"
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

func NewFeeChecker(c module.Client) authentication.AdmissionController {
	return FeeChecker{bank: bankv1aplha1.NewClient(c)}
}

// FeeChecker checks if the account has the required amount of fees
type FeeChecker struct {
	bank *bankv1aplha1.Client
}

func (x FeeChecker) Validate(req authentication.ValidateRequest) (resp authentication.ValidateResponse, err error) {
	payer := req.Tx.Payer()
	fee := req.Tx.Fee()
	// get balance of fee payer
	balance, err := x.bank.GetBalance(payer)
	if err != nil {
		return
	}
	// check if it has enough coins
	_, err = coin.SafeSub(balance.Balance, fee)
	if err != nil {
		return
	}
	// balance is enough
	return
}

func NewFeeDeduction(c module.Client) authentication.TransitionController {
	return FeeDeduction{bank: bankv1aplha1.NewClient(c)}
}

// FeeDeduction deducts fees from the transaction fee payer and sends them to the fee collector
type FeeDeduction struct {
	bank *bankv1aplha1.Client
}

func (x FeeDeduction) Deliver(req authentication.DeliverRequest) (resp authentication.DeliverResponse, err error) {
	// move coins and send them to fee collector
	return resp, x.bank.Send(req.Tx.Payer(), "fee_collector", req.Tx.Fee())
}
