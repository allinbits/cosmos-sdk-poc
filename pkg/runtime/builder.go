package runtime

import (
	"fmt"

	"github.com/fdymylja/tmos/apis/meta"
	"github.com/fdymylja/tmos/pkg/module"
	"github.com/fdymylja/tmos/pkg/runtime/client"
	"github.com/fdymylja/tmos/pkg/runtime/orm"
	"github.com/fdymylja/tmos/pkg/runtime/router"
	"k8s.io/klog/v2"
)

// NewBuilder creates a new Builder for the Runtime
func NewBuilder() *Builder {
	return &Builder{
		router: router.NewRouter(),
		store:  orm.NewStore(),
		rt:     &Runtime{},
	}
}

// Builder is used to create a new runtime from scratch
type Builder struct {
	modules []*module.Descriptor
	router  *router.Router
	store   *orm.Store
	rt      *Runtime
}

// AddModule adds a new module.Module to the list of modules to install
func (b *Builder) AddModule(m module.Module) {
	mb := module.NewBuilder()
	mc := client.NewClient(b.rt)
	m.Initialize(mc, mb)
	mc.SetUser(mb.Descriptor.Name) // set the authentication name for the module TODO: we should do this a lil better
	b.modules = append(b.modules, mb.Descriptor)
}

// Build installs the module.Modules provided and returns a fully functional runtime
func (b *Builder) Build() *Runtime {
	for _, m := range b.modules {
		if err := b.install(m); err != nil {
			panic(fmt.Errorf("error while installing module %s: %w", m.Name, err))
		}
	}
	b.rt.store = b.store
	b.rt.router = b.router
	return b.rt
}

func (b *Builder) install(m *module.Descriptor) error {
	// check name
	if !validModuleName(m.Name) {
		return fmt.Errorf("invalid module name: %s", m.Name)
	}
	// install state transition controllers
	for _, ctrl := range m.StateTransitionControllers {
		err := b.router.AddStateTransitionHandler(ctrl.StateTransition, ctrl.Controller)
		if err != nil {
			return err
		}
		klog.Infof("registered state transition %s for module %s", meta.Name(ctrl.StateTransition), m.Name)
	}
	// register admission controllers
	for _, ctrl := range m.AdmissionControllers {
		err := b.router.AddAdmissionController(ctrl.StateTransition, ctrl.Controller)
		if err != nil {
			return err
		}
		klog.Infof("registered admission controller %s for module %s", meta.Name(ctrl.StateTransition), m.Name)
	}
	// register state objects
	for _, so := range m.StateObjects {
		err := b.store.RegisterStateObject(so.StateObject)
		if err != nil {
			return err
		}
		klog.Infof("registered state object %s for module %s", meta.Name(so.StateObject), m.Name)
	}
	// TODO register admission + mutating admission + hooks
	// TODO register roles and dependencies
	return nil
}

func validModuleName(name string) bool {
	return name != ""
}
