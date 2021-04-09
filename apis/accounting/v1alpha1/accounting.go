package v1alpha1

import (
	"fmt"

	metav1alpha1 "github.com/fdymylja/tmos/apis/core/meta/v1alpha1"
	"github.com/fdymylja/tmos/pkg/application"
)

type Application struct {
}

func (a Application) Identifier() string {
	return "accounting"
}

func (a Application) RegisterStateObjects(registerer application.StateObjectRegisterer) {
	registerer.RegisterStateObject(&Balance{})
}

func (a Application) RegisterHandlers(registerer application.HandlerRegisterer) {
	registerer.RegisterDeliverHandler(&MsgSend{}, deliverMsgSend(), nil, application.HandlerPolicy{})
}

func deliverMsgSend() application.DeliverHandlerFn {
	return func(request application.DeliverRequest) error {
		msg := request.StateTransitionObject.(*MsgSend)

		balance := &Balance{State: &metav1alpha1.State{Id: msg.Sender}}
		exists := request.Client.Get(balance)
		if !exists {
			return fmt.Errorf("not found: %s", msg.Sender)
		}

		if balance.Amount < msg.Amount {
			return fmt.Errorf("not enough money")
		}
		// update balance of sender
		balance.Amount = balance.Amount - msg.Amount
		request.Client.Set(balance)
		// now update recipient balance
		recipientBalance := &Balance{State: &metav1alpha1.State{Id: msg.Recipient}}
		request.Client.Get(recipientBalance)
		recipientBalance.Amount = recipientBalance.Amount + msg.Amount
		request.Client.Set(recipientBalance)
		// success
		return nil
	}
}
