package runtime

import (
	"fmt"

	"github.com/fdymylja/tmos/core/abci"
	"github.com/fdymylja/tmos/core/rbac"
	rbacv1alpha1 "github.com/fdymylja/tmos/core/rbac/v1alpha1"
	"github.com/fdymylja/tmos/core/runtime"
	runtimev1alpha1 "github.com/fdymylja/tmos/core/runtime/v1alpha1"
	"github.com/fdymylja/tmos/runtime/authentication"
	"github.com/fdymylja/tmos/runtime/authentication/user"
	"github.com/fdymylja/tmos/runtime/client"
	"github.com/fdymylja/tmos/runtime/errors"
	"github.com/fdymylja/tmos/runtime/kv"
	"github.com/fdymylja/tmos/runtime/meta"
	"github.com/fdymylja/tmos/runtime/module"
	"github.com/fdymylja/tmos/runtime/orm"
	"github.com/fdymylja/tmos/runtime/orm/indexes"
	"github.com/fdymylja/tmos/runtime/orm/objects"
	"k8s.io/klog/v2"
)

var (
	errEmptyModuleName = errors.New("runtime: empty core name")
)

// NewBuilder creates a new Builder for the Runtime
func NewBuilder() *Builder {
	b := &Builder{
		moduleDescriptors: nil,
		moduleRoles:       map[string]*rbacv1alpha1.Role{},
		externalRole:      &rbacv1alpha1.Role{Id: rbacv1alpha1.ExternalAccountRoleID},
		rbac:              nil,
		decoder:           nil,
		router:            NewRouter(),
		rt:                &Runtime{},
	}

	// we already add the core modules in order
	// in theory we could add a dependency system
	// for genesis initialization, but for now lets keep it simple.
	b.AddModule(runtime.NewModule()) // needs to be first as it has state transitions/state object info
	b.rbac = rbac.NewModule()        // we set the rbac core inside so that we can prepare initial genesis with rbac
	b.AddModule(b.rbac)              // needs to be second as it provides the authorization layer
	b.AddModule(abci.NewModule())    // abci third so other modules can have access to this information

	// we add the initial external role, with basically no authorization towards no resource.
	b.externalRole = &rbacv1alpha1.Role{
		Id: rbacv1alpha1.ExternalAccountRoleID,
	}
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
	b.rt.store = b.store
	b.rt.router = b.router
	b.rt.modules = b.moduleDescriptors
	b.rt.rbac = b.rbac.AsAuthorizer()
	b.rt.user = user.NewUsersFromString(user.Runtime)
	switch b.decoder {
	case nil:
		klog.Warningf("no decoder - transactions sent to the ABCI application will be rejected")
	default:
		b.rt.txDecoder = b.decoder
	}

	return b.rt, nil
}

func (b *Builder) registerStateTransitionControllers(m module.Descriptor, role *rbacv1alpha1.Role) error {
	for _, ctrl := range m.StateTransitionExecutionHandlers {
		// add state transition controller to the router
		err := b.router.AddStateTransitionExecutionHandler(ctrl.StateTransition, ctrl.Controller)
		if err != nil {
			return err
		}

		// add deliver rights for the state transition
		err = role.Extend(runtimev1alpha1.Verb_Deliver, ctrl.StateTransition)
		if err != nil {
			return err
		}

		// if the state transition is marked as external we extend the external_account role
		if ctrl.External {
			err = b.externalRole.Extend(runtimev1alpha1.Verb_Deliver, ctrl.StateTransition)
			if err != nil {
				return err
			}
		}

		klog.Infof("registered state transition %s for core %s", meta.Name(ctrl.StateTransition), m.Name)
	}

	return nil
}

func (b *Builder) registerAdmissionControllers(m module.Descriptor) error {
	for _, ctrl := range m.StateTransitionAdmissionHandlers {
		err := b.router.AddStateTransitionAdmissionHandler(ctrl.StateTransition, ctrl.Controller)
		if err != nil {
			return err
		}
		klog.Infof("registered admission controller %s for core %s", meta.Name(ctrl.StateTransition), m.Name)
	}

	return nil
}

func (b *Builder) registerStateObjects(m module.Descriptor, role *rbacv1alpha1.Role) error {
	for _, so := range m.StateObjects {
		err := b.store.RegisterObject(so.StateObject, so.Options)
		if err != nil {
			return err
		}
		err = extendRoleForStateObject(role, so.StateObject)
		if err != nil {
			return err
		}
		klog.Infof("registered state object %s for core %s", meta.Name(so.StateObject), m.Name)
	}

	return nil
}

func (b *Builder) installStateObjects() error {
	for _, m := range b.moduleDescriptors {
		moduleRole := b.moduleRoles[m.Name]
		err := b.registerStateObjects(m, moduleRole)
		if err != nil {
			return fmt.Errorf("unable to install state objects for core %s: %w", m.Name, err)
		}
	}
	return nil
}

func (b *Builder) initEmptyRoles() error {
	for _, m := range b.moduleDescriptors {
		if isModuleNameEmpty(m.Name) {
			return errEmptyModuleName
		}
		// check if core role exists already
		if _, exists := b.moduleRoles[m.Name]; exists {
			return fmt.Errorf("core already registered %s", m.Name)
		}
		b.moduleRoles[m.Name] = &rbacv1alpha1.Role{Id: roleNameForModule(m.Name)}
	}
	return nil
}

func (b *Builder) installStateTransitions() error {
	for _, m := range b.moduleDescriptors {
		role := b.moduleRoles[m.Name]
		err := b.registerStateTransitionControllers(m, role)
		if err != nil {
			return fmt.Errorf("unable to install state transitions for core %s: %w", m.Name, err)
		}
	}
	return nil
}

func (b *Builder) installStateTransitionAdmissionHandlers() error {
	for _, m := range b.moduleDescriptors {
		err := b.registerAdmissionControllers(m)
		if err != nil {
			return err
		}
	}
	return nil
}

func (b *Builder) installStateTransitionPreExecHandlers() error {
	for _, m := range b.moduleDescriptors {
		for _, h := range m.StateTransitionPreExecHandlers {
			err := b.router.AddStateTransitionPreExecutionHandler(h.StateTransition, h.Handler)
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
			err := b.router.AddStateTransitionPostExecutionHandler(h.StateTransition, h.Handler)
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
	// initialize empty roles for modules
	if err := b.initEmptyRoles(); err != nil {
		return fmt.Errorf("unable to initialize core roles: %w", err)
	}

	// first we install state objects
	if err := b.installStateObjects(); err != nil {
		return fmt.Errorf("unable to install state objects: %w", err)
	}

	// after we install state transitions
	if err := b.installStateTransitions(); err != nil {
		return fmt.Errorf("unable to install state transitions: %w", err)
	}

	// then state transition admission controllers
	if err := b.installStateTransitionAdmissionHandlers(); err != nil {
		return fmt.Errorf("unable to install state transition admission handlers: %w", err)
	}

	// then state transition pre exec handlers
	if err := b.installStateTransitionPreExecHandlers(); err != nil {
		return fmt.Errorf("unable to install state transition pre execution handlers: %w", err)
	}

	// then state transition post exec handlers
	if err := b.installStateTransitionPostExecHandlers(); err != nil {
		return fmt.Errorf("unable to install state transition post execution handlers: %w", err)
	}

	// then transaction admission handlers
	if err := b.installAuthenticationAdmissionHandlers(); err != nil {
		return fmt.Errorf("unable to install authentication admission handlers: %w", err)
	}

	// then transaction post authentication handlers
	if err := b.installPostAuthenticationHandlers(); err != nil {
		return fmt.Errorf("unable to install post authentication handlers: %w", err)
	}

	// then dependencies
	if err := b.installDependencies(); err != nil {
		return fmt.Errorf("unable to install core dependencies: %w", err)
	}

	return nil
}

func (b *Builder) installAuthenticationAdmissionHandlers() error {
	for _, m := range b.moduleDescriptors {
		if m.AuthenticationExtension == nil {
			continue
		}
		for _, h := range m.AuthenticationExtension.AdmissionControllers {
			b.router.AddTransactionAdmissionController(h.Handler)
			klog.Infof("registered transaction admission controller %T for core %s", h.Handler, m.Name)
		}
	}
	return nil
}

func (b *Builder) installPostAuthenticationHandlers() error {
	for _, m := range b.moduleDescriptors {
		if m.AuthenticationExtension == nil {
			continue
		}
		for _, h := range m.AuthenticationExtension.TransitionControllers {
			b.router.AddTransactionPostAuthenticationController(h.Handler)
			klog.Infof("registered transaction post authentication controller %T for core %s", h.Handler, m.Name)
		}
	}
	return nil
}

func (b *Builder) installDependencies() error {
	for _, m := range b.moduleDescriptors {
		role := b.moduleRoles[m.Name]
		for _, st := range m.Needs {
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
	store := orm.NewStore(obj, idx)
	b.store = store

	return nil
}

func isModuleNameEmpty(name string) bool {
	return name == ""
}

func roleNameForModule(name string) string {
	const moduleRoleSuffix = "role"
	return fmt.Sprintf("%s-%s", name, moduleRoleSuffix)
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
