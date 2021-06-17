package apiserver

import "github.com/fdymylja/tmos/runtime/module"

type Module struct{}

func (m Module) Initialize(_ module.Client) module.Descriptor {
	return module.NewDescriptorBuilder().
		Named("apiserver").
		WithExtensionService(NewService()).
		Build()
}

func NewService() module.ExtensionService {
	return &apiServer{}
}

type apiServer struct {
	module.BaseService
}

func (a apiServer) Name() string {
	return "apiserver"
}
