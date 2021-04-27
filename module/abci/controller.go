package abci

import (
	"github.com/fdymylja/tmos/module/abci/v1alpha1"
)

type Controller struct {
}

func (c Controller) Name() string {
	return "abci"
}

func (c Controller) RegisterStateTransitions(client basicctrl.Client, register basicctrl.RegisterTransitionFn) {
	register(&v1alpha1.MsgSetBeginBlockState{}, newBeginBlockHandler(client))
}

func (c Controller) RegisterStateObjects(register basicctrl.RegisterStateObjectsFn) {
	register(&v1alpha1.BeginBlockState{})
}
