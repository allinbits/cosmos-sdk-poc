package extensions

import (
	"fmt"

	abciv1alpha1 "github.com/fdymylja/tmos/module/abci/v1alpha1"
	"github.com/fdymylja/tmos/module/x/authn/v1alpha1"
	"github.com/fdymylja/tmos/runtime/authentication"
	"github.com/fdymylja/tmos/runtime/module"
)

func New(c module.Client) module.AuthenticationExtension {
	return authExtension{c: c}
}

type authExtension struct {
	c module.Client
}

func (a authExtension) Initialize(builder *module.AuthenticationExtensionBuilder) {
	builder.
		WithAdmissionController(newTimeoutBlockExtension(a.c)).
		WithAdmissionController(newValidateMemoExtension(a.c))
}

func newTimeoutBlockExtension(c module.Client) timeoutBlockExtension {
	return timeoutBlockExtension{abci: abciv1alpha1.NewClient(c)}
}

type timeoutBlockExtension struct {
	abci *abciv1alpha1.Client
}

func (t timeoutBlockExtension) Validate(request authentication.ValidateRequest) (authentication.ValidateResponse, error) {
	tx := request.Tx.Raw().(*v1alpha1.Tx)
	currentBlock, err := t.abci.GetCurrentBlock()
	if err != nil {
		return authentication.ValidateResponse{}, err
	}
	if currentBlock.BlockNumber > tx.Body.TimeoutHeight {
		return authentication.ValidateResponse{}, fmt.Errorf("invalid block height")
	}
	return authentication.ValidateResponse{}, nil
}

func newValidateMemoExtension(c module.Client) validateMemoExtension {
	return validateMemoExtension{c: v1alpha1.NewClient(c)}
}

type validateMemoExtension struct {
	c *v1alpha1.Client
}

func (e validateMemoExtension) Validate(request authentication.ValidateRequest) (authentication.ValidateResponse, error) {
	tx := request.Tx.Raw().(*v1alpha1.Tx)
	memo := tx.Body.Memo
	memoLength := len(memo)
	params, err := e.c.GetParams()
	if err != nil {
		return authentication.ValidateResponse{}, nil
	}
	if uint64(memoLength) > params.MaxMemoCharacters {
		return authentication.ValidateResponse{}, fmt.Errorf("invalid memo length")
	}
	return authentication.ValidateResponse{}, nil
}
