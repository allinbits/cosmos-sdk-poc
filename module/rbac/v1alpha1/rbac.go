package v1alpha1

import "github.com/fdymylja/tmos/runtime/meta"

func (x *Role) GetID() meta.ID        { return meta.NewStringID(x.Id) }
func (x *RoleBinding) GetID() meta.ID { return meta.NewStringID(x.Subject) }

func (x *MsgBindRole) StateTransition()   {}
func (x *MsgCreateRole) StateTransition() {}
