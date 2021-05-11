package authn

import (
	"fmt"

	"github.com/fdymylja/tmos/module/x/authn/v1alpha1"
	"github.com/fdymylja/tmos/runtime/admission"
)

func NewCreateAccountAdmissionController() CreateAccountAdmissionController {
	return CreateAccountAdmissionController{}
}

type CreateAccountAdmissionController struct{}

func (CreateAccountAdmissionController) Validate(request admission.Request) (err error) {
	msg := request.Transition.(*v1alpha1.MsgCreateAccount)
	// validate message
	if msg.Account == nil {
		return fmt.Errorf("account is nil")
	}
	// validate account
	acc := msg.Account
	switch {
	case acc.Address == "":
		err = fmt.Errorf("no address provided")
	}
	return err
}
