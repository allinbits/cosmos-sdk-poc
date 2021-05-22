package v1alpha1

func (x *ValidatorUpdates) StateObject() {}
func (x *InitChainInfo) StateObject()    {}
func (x *EndBlockState) StateObject()    {}
func (x *Stage) StateObject()            {}
func (x *BeginBlockState) StateObject()  {}
func (x *DeliverTxState) StateObject()   {}
func (x *CheckTxState) StateObject()     {}
func (x *CurrentBlock) StateObject()     {}

func (x *MsgSetCheckTxState) StateTransition()     {}
func (x *MsgSetBeginBlockState) StateTransition()  {}
func (x *MsgSetDeliverTxState) StateTransition()   {}
func (x *MsgSetInitChain) StateTransition()        {}
func (x *MsgSetEndBlockState) StateTransition()    {}
func (x *MsgSetValidatorUpdates) StateTransition() {}
