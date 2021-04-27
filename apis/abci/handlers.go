package abci

import (
	"github.com/fdymylja/tmos/apis/abci/v1alpha1"
	basicctrl "github.com/fdymylja/tmos/pkg/controller/basic"
)

func newBeginBlockHandler(c basicctrl.Client) beginBlockHandler {
	return beginBlockHandler{c: c}
}

type beginBlockHandler struct {
	c basicctrl.Client
}

func (b beginBlockHandler) Check(_ basicctrl.CheckRequest) basicctrl.CheckResponse {
	return basicctrl.CheckResponse{}
}

func (b beginBlockHandler) Deliver(request basicctrl.DeliverRequest) basicctrl.DeliverResponse {
	msg := request.StateTransition.(*v1alpha1.MsgSetBeginBlockState)
	// set begin block state
	b.c.Set(&v1alpha1.BeginBlockState{
		BeginBlock: msg.BeginBlock,
	})
	//
	return basicctrl.DeliverResponse{}
}
