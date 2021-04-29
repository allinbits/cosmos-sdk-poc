package module

import (
	"github.com/fdymylja/tmos/runtime/meta"
	"github.com/fdymylja/tmos/runtime/store"
)

type Client interface {
	store.Store
	Deliver(transition meta.StateTransition) error
}

// Module defines a basic module which handles changes
type Module interface {
	Initialize(client Client, builder *Builder)
}
