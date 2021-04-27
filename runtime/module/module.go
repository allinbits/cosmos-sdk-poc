package module

import (
	"github.com/fdymylja/tmos/module/meta"
)

type Client interface {
	Get(object meta.StateObject) error
	Create(object meta.StateObject) error
	Update(object meta.StateObject) error
}

// Module defines a basic module which handles changes
type Module interface {
	Initialize(client Client, builder *Builder)
}
