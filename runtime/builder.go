package runtime

import (
	"fmt"

	"github.com/fdymylja/tmos/core/abci"
	"github.com/fdymylja/tmos/core/meta"
	"github.com/fdymylja/tmos/core/runtime"
	"github.com/fdymylja/tmos/runtime/authentication"
	"github.com/fdymylja/tmos/runtime/authentication/user"
	"github.com/fdymylja/tmos/runtime/authorization"
	"github.com/fdymylja/tmos/runtime/errors"
	"github.com/fdymylja/tmos/runtime/kv"
	"github.com/fdymylja/tmos/runtime/module"
	"github.com/fdymylja/tmos/runtime/orm"
	"github.com/fdymylja/tmos/runtime/orm/indexes"
	"github.com/fdymylja/tmos/runtime/orm/objects"
	"k8s.io/klog/v2"
)

var (
	errEmptyModuleName = errors.New("runtime: empty module name")
)

// NewBuilder creates a new Builder for the Runtime
func NewBuilder() *Builder {
	b := &Builder{
		moduleDescriptors: nil,
		decoder:           nil,
		store:             orm.Store{},
		rt: &Runtime{
			router:   NewRouter(),
			services: NewServiceOrchestrator(),
		},
	}

	b.AddModule(runtime.NewModule()) // needs to be first as it has state transitions/state object info
	b.AddModule(abci.NewModule())    // abci second so other modules can have access to this information

	return b
}

// Builder is used to create a new runtime from scratch
type Builder struct {
	moduleDescriptors []module.Descriptor
	decoder           authentication.TxDecoder

	router *Router
	store  orm.Store
	rt     *Runtime
}

// AddModule adds a new module.Module to the list of modules to install
func (b *Builder) AddModule(m module.Module) {

	mc := NewModuleClient(b.rt)
	descriptor := m.Initialize(mc)
	mc.SetUser(user.NewUsersFromString(descriptor.Name)) // set the authentication name for the module TODO: we should do this a lil better
	b.moduleDescriptors = append(b.moduleDescriptors, descriptor)
}

func (b *Builder) SetDecoder(txDecoder authentication.TxDecoder) {
	b.decoder = txDecoder
}

// Build installs the module.Modules provided and returns a fully functional runtime
func (b *Builder) Build() (*Runtime, error) {
	err := b.initStore()
	if err != nil {
		return nil, err
	}

	err = b.installModules()
	if err != nil {
		return nil, fmt.Errorf("unable to install modules: %w", err)
	}

	b.rt.store = b.store
	b.rt.modules = b.moduleDescriptors
	b.rt.authorizer, err = b.installAuthorizer()
	if err != nil {
		return nil, err
	}
	b.rt.user = user.NewUsersFromString(user.Runtime)
	switch b.decoder {
	case nil:
		klog.Warningf("no decoder - transactions sent to the ABCI application will be rejected")
	default:
		b.rt.txDecoder = b.decoder
	}
	return b.rt, nil
}

func (b *Builder) registerStateTransitionHandlers(m module.Descriptor) error {
	for _, handler := range m.StateTransitionExecutionHandlers {
		err := b.rt.router.AddStateTransitionExecutionHandler(handler.StateTransition, handler.Handler)
		if err != nil {
			return err
		}

		klog.Infof("registered state transition %s for module %s", meta.Name(handler.StateTransition), m.Name)
	}

	return nil
}

func (b *Builder) registerAdmissionHandlers(m module.Descriptor) error {
	for _, handler := range m.StateTransitionAdmissionHandlers {
		err := b.rt.router.AddStateTransitionAdmissionHandler(handler.StateTransition, handler.AdmissionHandler)
		if err != nil {
			return err
		}

		klog.Infof("registered admission handler %s for module %s", meta.Name(handler.StateTransition), m.Name)
	}

	return nil
}

func (b *Builder) registerStateObjects(md module.Descriptor) error {
	for _, so := range md.StateObjects {
		err := b.store.RegisterObject(so.StateObject, so.SchemaDefinition)
		if err != nil {
			return err
		}

		klog.Infof("registered state object %s for module %s", meta.Name(so.StateObject), md.Name)
	}

	return nil
}

func (b *Builder) installStateObjects() error {
	for _, md := range b.moduleDescriptors {
		err := b.registerStateObjects(md)
		if err != nil {
			return fmt.Errorf("unable to install state objects for module %s: %w", md.Name, err)
		}
	}

	return nil
}

func (b *Builder) installStateTransitions() error {
	for _, m := range b.moduleDescriptors {
		err := b.registerStateTransitionHandlers(m)
		if err != nil {
			return fmt.Errorf("unable to install state transitions for module %s: %w", m.Name, err)
		}
	}

	return nil
}

func (b *Builder) installStateTransitionAdmissionHandlers() error {
	for _, m := range b.moduleDescriptors {
		err := b.registerAdmissionHandlers(m)
		if err != nil {
			return err
		}
	}

	return nil
}

func (b *Builder) installStateTransitionPreExecHandlers() error {
	for _, m := range b.moduleDescriptors {
		for _, h := range m.StateTransitionPreExecHandlers {
			err := b.rt.router.AddStateTransitionPreExecutionHandler(h.StateTransition, h.Handler)
			if err != nil {
				return fmt.Errorf("unable to install state transition pre execution handler for module %s: %w", m.Name, err)
			}
			klog.Infof(
				"registered state transition pre execution handler %T for state transition %s by module %s",
				h.Handler,
				meta.Name(h.StateTransition),
				m.Name)
		}
	}

	return nil
}

func (b *Builder) installStateTransitionPostExecHandlers() error {
	for _, m := range b.moduleDescriptors {
		for _, h := range m.StateTransitionPostExecutionHandlers {
			err := b.rt.router.AddStateTransitionPostExecutionHandler(h.StateTransition, h.Handler)
			if err != nil {
				return fmt.Errorf("unable to install state transition post execution handler for module %s: %w", m.Name, err)
			}
			klog.Infof(
				"registered state transition post execution handler %T for state transition %s by module %s",
				h.Handler,
				meta.Name(h.StateTransition),
				m.Name)
		}
	}
	return nil
}

func (b *Builder) installModules() error {

	for _, b := range b.moduleDescriptors {
		if isModuleNameEmpty(b.Name) {
			return fmt.Errorf("empty module name for")
		}
	}

	if err := b.installStateObjects(); err != nil {
		return fmt.Errorf("unable to install state objects: %w", err)
	}

	if err := b.installStateTransitions(); err != nil {
		return fmt.Errorf("unable to install state transitions: %w", err)
	}

	if err := b.installStateTransitionAdmissionHandlers(); err != nil {
		return fmt.Errorf("unable to install state transition admission handlers: %w", err)
	}

	if err := b.installStateTransitionPreExecHandlers(); err != nil {
		return fmt.Errorf("unable to install state transition pre execution handlers: %w", err)
	}

	if err := b.installStateTransitionPostExecHandlers(); err != nil {
		return fmt.Errorf("unable to install state transition post execution handlers: %w", err)
	}

	if err := b.installAuthenticationAdmissionHandlers(); err != nil {
		return fmt.Errorf("unable to install authentication admission handlers: %w", err)
	}

	if err := b.installPostAuthenticationHandlers(); err != nil {
		return fmt.Errorf("unable to install post authentication handlers: %w", err)
	}

	if err := b.installDependencies(); err != nil {
		return fmt.Errorf("unable to install module dependencies: %w", err)
	}

	if err := b.installExtensions(); err != nil {
		return fmt.Errorf("unable to install module extensions: %w", err)
	}

	return nil
}

func (b *Builder) installAuthenticationAdmissionHandlers() error {
	for _, m := range b.moduleDescriptors {
		if m.AuthAdmissionHandlers == nil {
			continue
		}
		for _, h := range m.AuthAdmissionHandlers {
			b.rt.router.AddAuthAdmissionHandler(h)
			klog.Infof("registered transaction admission handler %T for module %s", h, m.Name)
		}
	}
	return nil
}

func (b *Builder) installPostAuthenticationHandlers() error {
	for _, m := range b.moduleDescriptors {
		if m.PostAuthenticationHandler == nil {
			continue
		}
		for _, h := range m.PostAuthenticationHandler {
			b.rt.router.AddTransactionPostAuthenticationHandler(h)
			klog.Infof("registered transaction post authentication handler %T for module %s", h, m.Name)
		}
	}
	return nil
}

func (b *Builder) installDependencies() error {
	for _, md := range b.moduleDescriptors {
		for _, st := range md.Needs {
			_, err := b.rt.router.GetStateTransitionExecutionHandler(st)
			if err != nil {
				return fmt.Errorf("dependency cannot be accomplished: %w", err)
			}
		}
	}
	return nil
}

func (b *Builder) initStore() error {
	okv := kv.NewBadger()
	obj := objects.NewStore(okv)
	idxKv := kv.NewBadger()
	idx := indexes.NewStore(idxKv)

	b.store = orm.NewStore(obj, idx)

	return nil
}

func (b *Builder) installExtensions() error {
	for _, m := range b.moduleDescriptors {
		b.rt.services.AddServices(m.Name, m.Services...)
	}
	return nil
}

func (b *Builder) installAuthorizer() (authorization.Authorizer, error) {
	var authz authorization.Authorizer
	for _, m := range b.moduleDescriptors {
		if authz != nil && m.Authorizer != nil {
			return nil, fmt.Errorf("authorizer was already set, module %s defines another authorizer", m.Name) // TODO(fdymylja): support multiple authorizers by creating a chain authz
		}
		if m.Authorizer != nil {
			authz = m.Authorizer
		}
	}

	if authz == nil {
		klog.Warningf("no authorizer set, using an always allowing authorizer")
		authz = authorization.NewAlwaysAllowAuthorizer()
	}

	return authz, nil
}

func isModuleNameEmpty(name string) bool {
	return name == ""
}
