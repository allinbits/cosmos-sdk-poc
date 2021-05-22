package v1alpha1

func (x *CurrentAccountNumber) StateObject() {}
func (x *Account) StateObject()              {}
func (x *Params) StateObject()               {}
func (x *MsgCreateAccount) StateTransition() {}
func (x *MsgUpdateAccount) StateTransition() {}
func (x *MsgDeleteAccount) StateTransition() {}
