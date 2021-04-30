package authn

import (
	"fmt"

	abciv1alpha1 "github.com/fdymylja/tmos/module/abci/v1alpha1"
	"github.com/fdymylja/tmos/module/x/authn/v1alpha1"
	"github.com/fdymylja/tmos/runtime/authentication"
	"github.com/fdymylja/tmos/runtime/module"
)

func NewExtension(c module.Client) module.AuthenticationExtension {
	return authExtension{c: c}
}

type authExtension struct {
	c module.Client
}

func (a authExtension) Initialize(builder *module.AuthenticationExtensionBuilder) {
	builder.
		WithAdmissionController(timeoutBlockExtension{abci: abciv1alpha1.NewClient(a.c)})
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
