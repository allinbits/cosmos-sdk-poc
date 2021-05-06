package runtime

import (
	"fmt"

	"github.com/fdymylja/tmos/module/abci"
	"github.com/fdymylja/tmos/module/rbac"
	"github.com/fdymylja/tmos/module/runtime"
	"github.com/fdymylja/tmos/runtime/authentication"
	"github.com/fdymylja/tmos/runtime/errors"
	"github.com/fdymylja/tmos/runtime/meta"
	"github.com/fdymylja/tmos/runtime/module"
	"github.com/fdymylja/tmos/runtime/store/badger"
	"k8s.io/klog/v2"
)

// NewBuilder creates a new Builder for the Runtime
func NewBuilder() *Builder {
	return &Builder{
		installedModules: map[string]struct{}{},
		modules:          nil,
		authn:            nil,
		router:           NewRouter(),
		store:            badger.NewStore(),
		rt:               &Runtime{},
	}
}

// Builder is used to create a new runtime from scratch
type Builder struct {
	installedModules map[string]struct{} // installedModules is used to check if multiple modules with the same name are being installed
	modules          []*module.Descriptor
	authn            authentication.Authenticator
	router           *Router
	store            *badger.Store
	rt               *Runtime
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
	b.AddModule(abci.NewModule())
	b.AddModule(runtime.NewModule())
	b.AddModule(rbac.NewModule())
	// install all modules
	for _, m := range b.modules {
		// check if already installed
		if _, exists := b.installedModules[m.Name]; exists {
			return nil, fmt.Errorf("double registration of module named %s", m.Name)
		}
		if err := b.install(m); err != nil {
			return nil, fmt.Errorf("error while installing module %s: %w", m.Name, err)
		}
		// mark as installed module
		b.installedModules[m.Name] = struct{}{}
	}
	b.rt.store = b.store
	b.rt.router = b.router
	b.rt.modules = b.modules

	switch b.authn {
	case nil:
		klog.Warningf("no authenticator was set up - transactions sent to the ABCI application will be rejected")
	default:
		b.rt.authn = b.authn
	}

	return b.rt, nil
}

func (b *Builder) install(m *module.Descriptor) error {
	// check name
	if isModuleNameEmpty(m.Name) {
		return errors.ErrEmptyModuleName
	}

	// install state transition controllers
	for _, ctrl := range m.StateTransitionControllers {
		err := b.router.AddStateTransitionController(ctrl.StateTransition, ctrl.Controller)
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
	for _, xt := range m.AuthenticationExtension.TransitionControllers {
		b.router.AddTransactionPostAuthenticationController(xt.Handler)
		klog.Infof("registering authentication post admission controller %T for module %s", xt.Handler, m.Name)
	}
	return nil
}

func isModuleNameEmpty(name string) bool {
	return name == ""
}
