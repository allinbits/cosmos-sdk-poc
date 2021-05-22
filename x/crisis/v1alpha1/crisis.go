package v1alpha1

func (x *InvariantHandler) StateObject() {}

func (x *Params) StateObject() {}

func (x *MsgRegisterInvariant) StateTransition()        {}
func (x *MsgVerifyInvariant) StateTransition()          {}
func (x *MsgVerifyInvariantCosmosSDK) StateTransition() {}
