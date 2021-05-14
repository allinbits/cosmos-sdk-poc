package crisis

import "github.com/fdymylja/tmos/runtime/module"

type Module struct {
}

func (m Module) Initialize(client module.Client) module.Descriptor {
	return module.NewDescriptorBuilder().
		Named("crisis").
		Build()
}
