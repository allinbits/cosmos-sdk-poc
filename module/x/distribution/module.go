package distribution

import (
	"github.com/fdymylja/tmos/module/x/distribution/extensions"
	"github.com/fdymylja/tmos/runtime/module"
)

func NewModule() Module {
	return Module{}
}

type Module struct {
}

func (m Module) Initialize(client module.Client, builder *module.Builder) {
	builder.
		Named("distribution").
		ExtendsAuthentication(extensions.NewAuthentication(client))
}
