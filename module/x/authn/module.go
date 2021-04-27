package authn

import (
	"github.com/fdymylja/tmos/module/x/authn/tx"
	"github.com/fdymylja/tmos/module/x/authn/v1alpha1"
	"github.com/fdymylja/tmos/runtime/module"
	"k8s.io/klog/v2"
)

// Module implements the authentication.Module
type Module struct {
}

func NewModule() module.Module {
	return Module{}
}

func (m Module) Initialize(c module.Client, builder *module.Builder) {
	builder.
		Named("authn").
		HandlesStateTransition(&v1alpha1.MsgCreateAccount{}, NewCreateAccountController(c)).
		HandlesAdmission(&v1alpha1.MsgCreateAccount{}, NewCreateAccountAdmissionController()).
		OwnsStateObject(&v1alpha1.Account{}).
		OwnsStateObject(&v1alpha1.Params{}).
		OwnsStateObject(&v1alpha1.CurrentAccountNumber{})
}

func (m Module) DecodeTransaction(txBytes []byte) error {
	t, err := tx.DecodeTx(txBytes)
	if err != nil {
		return err
	}
	klog.Infof("%s", t)
	return nil
}
