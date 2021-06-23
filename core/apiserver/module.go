package apiserver

import (
	"context"
	"net/http"

	"github.com/fdymylja/tmos/core/apiserver/router"
	"github.com/fdymylja/tmos/core/meta"
	runtimev1alpha1 "github.com/fdymylja/tmos/core/runtime/v1alpha1"
	"github.com/fdymylja/tmos/runtime/client"
	"github.com/fdymylja/tmos/runtime/module"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"k8s.io/klog/v2"
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
	runtime runtimev1alpha1.ClientSet
	server  *http.Server
}

func (a *apiServer) SetClient(client client.Client) {
	a.runtime = runtimev1alpha1.NewClientSet(client)
	a.Client = client
}

func (a *apiServer) Name() string {
	return "apiserver"
}

func (a *apiServer) Start() error {
	builder := router.NewBuilder(a.Client)
	// we get the available state objects
	desc, err := a.runtime.ModuleDescriptors().Get()
	if err != nil {
		return err
	}
	for _, m := range desc.Modules {
		for _, so := range m.StateObjects {
			obj, err := protoregistry.GlobalTypes.FindMessageByName((protoreflect.FullName)(so.ProtobufFullname))
			if err != nil {
				return err
			}
			stateObject := obj.New().Interface().(meta.StateObject)
			err = builder.CreateStateObjectHandler(stateObject, so.SchemaDefinition)
			if err != nil {
				return err
			}
		}
	}

	mux, err := builder.Build()
	if err != nil {
		return err
	}

	go func() {
		a.server = &http.Server{Handler: mux, Addr: ":8080"} // TODO configurable
		err := a.server.ListenAndServe()
		if err != nil {
			klog.ErrorS(err, "whilst serving")
		}
	}()

	return nil
}

func (a *apiServer) Stop() error {
	return a.server.Shutdown(context.TODO())
}
