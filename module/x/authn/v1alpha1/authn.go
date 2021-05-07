package v1alpha1

import (
	"github.com/fdymylja/tmos/runtime/meta"
)

var CurrentAccountNumberID = meta.NewStringID("acc_num")
var ParamsID = meta.NewStringID("params")

func (x *CurrentAccountNumber) GetID() meta.ID { return CurrentAccountNumberID }
func (x *Account) GetID() meta.ID              { return meta.NewStringID(x.Address) }
func (x *Params) GetID() meta.ID               { return ParamsID }
func (x *MsgCreateAccount) StateTransition()   {}
func (x *MsgUpdateAccount) StateTransition()   {}
func (x *MsgDeleteAccount) StateTransition()   {}
