package runtime

import (
	"fmt"

	"github.com/fdymylja/tmos/module/abci"
	"github.com/fdymylja/tmos/runtime/authentication"
	"github.com/fdymylja/tmos/runtime/meta"
	"github.com/fdymylja/tmos/runtime/module"
	"github.com/fdymylja/tmos/runtime/store/badger"
	"k8s.io/klog/v2"
)

// NewBuilder creates a new Builder for the Runtime
func NewBuilder() *Builder {
	return &Builder{
		router: NewRouter(),
		store:  badger.NewStore(),
		rt:     &Runtime{},
	}
}

// Builder is used to create a new runtime from scratch
type Builder struct {
	modules []*module.Descriptor
	authn   authentication.Authenticator
	router  *Router
	store   *badger.Store
	rt      *Runtime
}

// AddModule adds a new module.Module to the list of modules to install
func (b *Builder) AddModule(m module.Module) {
	mb := module.NewModuleBuilder()
	mc := newClient(b.rt)
	m.Initialize(mc, mb)
	mc.SetUser(mb.Descriptor.Name) // set the authentication name for the module TODO: we should do this a lil better
	b.modules = append(b.modules, mb.Descriptor)
}

func (b *Builder) SetAuthenticator(authn authentication.Authenticator) {
	b.authn = authn
}

// Build installs the module.Modules provided and returns a fully functional runtime
func (b *Builder) Build() (*Runtime, error) {
	// add core modules
	abciModule := abci.NewModule()
	b.AddModule(abciModule)

	// install all modules
	for _, m := range b.modules {
		if err := b.install(m); err != nil {
			return nil, fmt.Errorf("error while installing module %s: %w", m.Name, err)
		}
	}
	b.rt.store = b.store
	b.rt.router = b.router
	b.rt.modules = b.modules

	switch b.authn {
	case nil:
		klog.Warningf("no authenticator was set up - resorting to a NO-OP authenticator")
	default:
		b.rt.authn = b.authn
	}

	return b.rt, nil
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
		err := b.router.AddStateTransitionAdmissionController(ctrl.StateTransition, ctrl.Controller)
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
	// register authentication extensions
	if m.AuthenticationExtension == nil {
		return nil
	}

	// add authentication admission controllers
	for _, xt := range m.AuthenticationExtension.AdmissionControllers {
		b.router.AddTransactionAdmissionController(xt.Handler)
		klog.Infof("registering authentication admission controller %T for module %s", xt.Handler, m.Name)
	}

	return nil
}

func validModuleName(name string) bool {
	return name != ""
}
