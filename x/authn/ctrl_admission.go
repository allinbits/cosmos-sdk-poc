package authn

import (
	"fmt"

	"github.com/fdymylja/tmos/runtime/statetransition"
	v1alpha12 "github.com/fdymylja/tmos/x/authn/v1alpha1"
)

func NewCreateAccountAdmissionHandler() CreateAccountAdmissionHandler {
	return CreateAccountAdmissionHandler{}
}

type CreateAccountAdmissionHandler struct{}

func (CreateAccountAdmissionHandler) Validate(request statetransition.AdmissionRequest) (err error) {
	msg := request.Transition.(*v1alpha12.MsgCreateAccount)
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
