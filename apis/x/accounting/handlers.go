package accounting

import (
	"fmt"

	meta "github.com/fdymylja/tmos/apis/meta/v1alpha1"
	"github.com/fdymylja/tmos/apis/x/accounting/v1alpha1"
	basicctrl "github.com/fdymylja/tmos/pkg/controller/basic"
)

var _ basicctrl.StateTransitionHandler = msgSendHandler{}

func newMsgSendHandler(client basicctrl.Client) msgSendHandler {
	return msgSendHandler{client: client}
}

type msgSendHandler struct {
	client basicctrl.Client
}

func (m msgSendHandler) Check(request basicctrl.CheckRequest) basicctrl.CheckResponse {
	// do checks on msg validity the old msg.ValidateBasic
	panic("implement me")
}

func (m msgSendHandler) Deliver(request basicctrl.DeliverRequest) basicctrl.DeliverResponse {
	// TODO check if request was authorized by sender
	msg := request.StateTransition.(*v1alpha1.MsgSend)
	// get the balance of the sender
	senderBalance := &v1alpha1.Balance{ObjectMeta: &meta.ObjectMeta{Id: msg.Sender}}
	exists := m.client.Get(senderBalance)
	if !exists {
		return basicctrl.DeliverResponse{
			Error: fmt.Errorf("balance not found for user: %s", msg.Sender),
		}
	}
	// check if it has enough money
	if senderBalance.Amount <= msg.Amount {
		return basicctrl.DeliverResponse{Error: fmt.Errorf("insufficient senderBalance")}
	}
	// get the receiver's balance
	recvBalance := &v1alpha1.Balance{ObjectMeta: &meta.ObjectMeta{Id: msg.Receiver}}
	m.client.Get(recvBalance) // we don't care to check for existence
	// apply state change
	senderBalance.Amount = senderBalance.Amount - msg.Amount
	recvBalance.Amount = recvBalance.Amount + msg.Amount
	m.client.Set(senderBalance)
	m.client.Set(recvBalance)
	// no error
	return basicctrl.DeliverResponse{Error: nil}
}
