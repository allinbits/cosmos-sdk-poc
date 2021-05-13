package v1alpha1

import "github.com/fdymylja/tmos/runtime/meta"

func (x *Balance) GetID() meta.ID {
	return meta.NewStringID(x.Address)
}

func (x *MsgSendCoins) StateTransition()  {}
func (x *MsgSetBalance) StateTransition() {}
