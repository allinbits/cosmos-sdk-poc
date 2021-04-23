package authn

import (
	"github.com/fdymylja/tmos/apis/meta"
	"github.com/fdymylja/tmos/apis/x/authn/tx"
	"github.com/fdymylja/tmos/pkg/controller/authentication"
	"github.com/fdymylja/tmos/pkg/controller/basic"
	"k8s.io/klog/v2"
)

// Controller implements the authentication.Controller
type Controller struct {
}

func NewController() authentication.Controller {
	return Controller{}
}

func (c Controller) Name() string {
	panic("implement me")
}

func (c Controller) RegisterStateTransitions(client basic.Client, register basic.RegisterTransitionFn) {
	panic("implement me")
}

func (c Controller) RegisterStateObjects(register basic.RegisterStateObjectsFn) {
	panic("implement me")
}

func (c Controller) DecodeTx(txBytes []byte) (transitions []meta.StateTransition, authenticatedAccounts []string, err error) {
	panic("implement")
}

func (c Controller) DecodeTransaction(txBytes []byte) error {
	t, err := tx.DecodeTx(txBytes)
	if err != nil {
		return err
	}
	klog.Infof("%s", t)
	return nil
}