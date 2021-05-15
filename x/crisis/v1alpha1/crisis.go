package v1alpha1

import "github.com/fdymylja/tmos/runtime/meta"

var (
	ParamsID = meta.NewStringID("params")
)

func (x *InvariantHandler) GetID() meta.ID {
	return nil
}

func (x *Params) GetID() meta.ID {
	return ParamsID
}

func (x *MsgRegisterInvariant) StateTransition()        {}
func (x *MsgVerifyInvariant) StateTransition()          {}
func (x *MsgVerifyInvariantCosmosSDK) StateTransition() {}
