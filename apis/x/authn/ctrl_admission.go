package authn

import (
	"fmt"

	"github.com/fdymylja/tmos/apis/x/authn/v1alpha1"
	"github.com/fdymylja/tmos/pkg/controller"
)

func NewCreateAccountAdmissionController() CreateAccountAdmissionController {
	return CreateAccountAdmissionController{}
}

type CreateAccountAdmissionController struct{}

func (CreateAccountAdmissionController) Validate(request controller.AdmissionRequest) (resp controller.AdmissionResponse, err error) {
	msg := request.Transition.(*v1alpha1.MsgCreateAccount)
	// validate message
	if msg.Account == nil {
		return resp, fmt.Errorf("account is nil")
	}
	// validate account
	acc := msg.Account
	switch {
	case acc.Address == "":
		err = fmt.Errorf("no address provided")
	}
	return resp, err
}
