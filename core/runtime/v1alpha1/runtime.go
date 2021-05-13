package v1alpha1

import "github.com/fdymylja/tmos/runtime/meta"

var (
	StateObjectsListID     = meta.NewStringID("state_objects")
	StateTransitionsListID = meta.NewStringID("state_transitions")
)

func (x *StateObjectsList) GetID() meta.ID     { return StateObjectsListID }
func (x *StateTransitionsList) GetID() meta.ID { return StateTransitionsListID }

func (x *CreateStateObjectsList) StateTransition()     {}
func (x *CreateStateTransitionsList) StateTransition() {}
