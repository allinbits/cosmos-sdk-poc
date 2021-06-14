package runtime

import (
	"fmt"

	"github.com/fdymylja/tmos/core/abci"
	meta "github.com/fdymylja/tmos/core/meta"
	"github.com/fdymylja/tmos/core/rbac"
	rbacv1alpha1 "github.com/fdymylja/tmos/core/rbac/v1alpha1"
	"github.com/fdymylja/tmos/core/runtime"
	runtimev1alpha1 "github.com/fdymylja/tmos/core/runtime/v1alpha1"
	"github.com/fdymylja/tmos/runtime/api"
	"github.com/fdymylja/tmos/runtime/authentication"
	"github.com/fdymylja/tmos/runtime/authentication/user"
	"github.com/fdymylja/tmos/runtime/client"
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
		moduleRoles:       map[string]*rbacv1alpha1.Role{},
		externalRole:      &rbacv1alpha1.Role{Id: rbacv1alpha1.ExternalAccountRoleID},
		rbac:              nil,
		decoder:           nil,
		rt: &Runtime{
			router: NewRouter(),
		},
		apiServer:         nil,
	}

	// we already add the core modules in order
	// in theory we could add a dependency system
	// for genesis initialization, but for now lets keep it simple.
	b.AddModule(runtime.NewModule()) // needs to be first as it has state transitions/state object info
	b.AddModule(b.rbac)              // needs to be second as it provides the authorization layer
	b.AddModule(abci.NewModule())    // abci third so other modules can have access to this information

	return b
}

// Builder is used to create a new runtime from scratch
type Builder struct {
	moduleDescriptors []module.Descriptor

	moduleRoles  map[string]*rbacv1alpha1.Role
	externalRole *rbacv1alpha1.Role
	rbac         *rbac.Module
	decoder      authentication.TxDecoder

	router *Router
	store  orm.Store
	rt     *Runtime

	apiServer *api.Builder
}

// AddModule adds a new module.Module to the list of modules to install
func (b *Builder) AddModule(m module.Module) {
	type userSetter interface {
		SetUser(users user.Users)
	}

	mc := client.New(newRuntimeAsServer(b.rt))
	descriptor := m.Initialize(mc)
	mc.(userSetter).SetUser(user.NewUsersFromString(descriptor.Name)) // set the authentication name for the core TODO: we should do this a lil better
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

	for moduleName, moduleRole := range b.moduleRoles {
		b.rbac.AddInitialRole(moduleRole, &rbacv1alpha1.RoleBinding{
			Subject: moduleName,
			RoleRef: moduleRole.Id,
		})
	}

	// add external role to rbac with no binding
	b.rbac.AddInitialRole(b.externalRole, nil)
	b.rt.modules = b.moduleDescriptors
	b.rt.authorizer = b.rbac.AsAuthorizer()
	b.rt.user = user.NewUsersFromString(user.Runtime)
	switch b.decoder {
	case nil:
		klog.Warningf("no decoder - transactions sent to the ABCI application will be rejected")
	default:
		b.rt.txDecoder = b.decoder
	}
	// populate api server
	b.apiServer = api.NewServer(b.store)
	for _, m := range b.moduleDescriptors {
		err = b.apiServer.RegisterModuleAPI(m)
		if err != nil {
			return nil, err
		}
	}
	// start api server
	b.apiServer.Start()
	return b.rt, nil
}

func (b *Builder) registerStateTransitionHandlers(m module.Descriptor) error {
	for _, handler := range m.StateTransitionExecutionHandlers {
		err := b.rt.router.AddStateTransitionExecutionHandler(handler.StateTransition, handler.Handler)
		if err != nil {
			return err
		}

		err = b.moduleRoles[m.Name].Extend(runtimev1alpha1.Verb_Deliver, handler.StateTransition)
		if err != nil {
			return err
		}

		if handler.External {
			err = b.externalRole.Extend(runtimev1alpha1.Verb_Deliver, handler.StateTransition)
			if err != nil {
				return err
			}
		}

		klog.Infof("registered state transition %s for core %s", meta.Name(handler.StateTransition), m.Name)
	}

	return nil
}

func (b *Builder) registerAdmissionHandlers(m module.Descriptor) error {
	for _, handler := range m.StateTransitionAdmissionHandlers {
		err := b.rt.router.AddStateTransitionAdmissionHandler(handler.StateTransition, handler.AdmissionHandler)
		if err != nil {
			return err
		}

		klog.Infof("registered admission handler %s for core %s", meta.Name(handler.StateTransition), m.Name)
	}

	return nil
}

func (b *Builder) registerStateObjects(md module.Descriptor) error {
	for _, so := range md.StateObjects {
		err := b.rt.store.RegisterObject(so.StateObject, so.Options)
		if err != nil {
			return err
		}

		err = extendRoleForStateObject(b.moduleRoles[md.Name], so.StateObject)
		if err != nil {
			return err
		}

		klog.Infof("registered state object %s for core %s", meta.Name(so.StateObject), md.Name)
	}

	return nil
}

func (b *Builder) installStateObjects() error {
	for _, md := range b.moduleDescriptors {
		err := b.registerStateObjects(md)
		if err != nil {
			return fmt.Errorf("unable to install state objects for core %s: %w", md.Name, err)
		}
	}

	return nil
}

// initEmptyRoles creates an empty role by every module descriptor.
func (b *Builder) initEmptyRoles() error {
	for _, m := range b.moduleDescriptors {
		if isModuleNameEmpty(m.Name) {
			return errEmptyModuleName
		}

		if b.roleExists(m.Name) {
			return fmt.Errorf("core already registered %s", m.Name)
		}

		b.moduleRoles[m.Name] = rbacv1alpha1.NewEmptyRole(m.Name)
	}

	return nil
}

func (b *Builder) roleExists(name string) bool {
	_, exists := b.moduleRoles[name]
	return exists
}

func (b *Builder) installStateTransitions() error {
	for _, m := range b.moduleDescriptors {
		err := b.registerStateTransitionHandlers(m)
		if err != nil {
			return fmt.Errorf("unable to install state transitions for core %s: %w", m.Name, err)
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
				return fmt.Errorf("unable to install state transition pre execution handler for core %s: %w", m.Name, err)
			}
			klog.Infof(
				"registered state transition pre execution handler %T for state transition %s by core %s",
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
				return fmt.Errorf("unable to install state transition post execution handler for core %s: %w", m.Name, err)
			}
			klog.Infof(
				"registered state transition post execution handler %T for state transition %s by core %s",
				h.Handler,
				meta.Name(h.StateTransition),
				m.Name)
		}
	}
	return nil
}

func (b *Builder) installModules() error {
	if err := b.initEmptyRoles(); err != nil {
		return fmt.Errorf("unable to initialize core roles: %w", err)
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
		return fmt.Errorf("unable to install core dependencies: %w", err)
	}

	return nil
}

func (b *Builder) installAuthenticationAdmissionHandlers() error {
	for _, m := range b.moduleDescriptors {
		if m.AuthAdmissionHandlers == nil {
			continue
		}
		for _, h := range m.AuthAdmissionHandlers {
			b.rt.router.AddTransactionAdmissionHandler(h.Handler)
			klog.Infof("registered transaction admission handler %T for core %s", h.Handler, m.Name)
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
			b.rt.router.AddTransactionPostAuthenticationHandler(h.Handler)
			klog.Infof("registered transaction post authentication handler %T for core %s", h.Handler, m.Name)
		}
	}
	return nil
}

func (b *Builder) installDependencies() error {
	for _, md := range b.moduleDescriptors {
		role := b.moduleRoles[md.Name]
		for _, st := range md.Needs {
			err := role.Extend(runtimev1alpha1.Verb_Deliver, st)
			if err != nil {
				return fmt.Errorf("error while registering core dependency %s: %w", meta.Name(st), err)
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

	b.rt.store = orm.NewStore(obj, idx)

	return nil
}

func isModuleNameEmpty(name string) bool {
	return name == ""
}

func extendRoleForStateObject(role *rbacv1alpha1.Role, so meta.StateObject) (err error) {
	err = role.Extend(runtimev1alpha1.Verb_Create, so)
	if err != nil {
		return err
	}
	err = role.Extend(runtimev1alpha1.Verb_Delete, so)
	if err != nil {
		return err
	}
	err = role.Extend(runtimev1alpha1.Verb_Update, so)
	if err != nil {
		return err
	}
	err = role.Extend(runtimev1alpha1.Verb_Get, so)
	if err != nil {
		return err
	}
	err = role.Extend(runtimev1alpha1.Verb_List, so)
	if err != nil {
		return err
	}
	return nil
}
