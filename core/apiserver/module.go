package apiserver

import (
	runtimev1alpha1 "github.com/fdymylja/tmos/core/runtime/v1alpha1"
	"github.com/fdymylja/tmos/runtime/client"
	"github.com/fdymylja/tmos/runtime/module"
	"github.com/gorilla/mux"
)

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
	mux     *mux.Router
	runtime runtimev1alpha1.ClientSet
}

func (a *apiServer) SetClient(client client.RuntimeClient) {
	a.runtime = runtimev1alpha1.NewClientSet(client)
}

func (a *apiServer) Name() string {
	return "apiserver"
}

func (a *apiServer) Start() error {
	// we get the available state objects
	desc, err := a.runtime.ModuleDescriptors().Get()
	if err != nil {
		return err
	}
	panic(desc)
}
