package runtime

import "github.com/fdymylja/tmos/runtime/meta"

type ModuleClient interface {
	Store
	Deliver(transition meta.StateTransition) error
}

// Module defines a basic module which handles changes
type Module interface {
	Initialize(client ModuleClient, builder *ModuleBuilder)
}
