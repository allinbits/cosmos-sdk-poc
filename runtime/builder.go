package runtime

import (
	"fmt"

	"github.com/fdymylja/tmos/module/abci"
	"github.com/fdymylja/tmos/module/rbac"
	rbacv1alpha1 "github.com/fdymylja/tmos/module/rbac/v1alpha1"
	"github.com/fdymylja/tmos/module/runtime"
	runtimev1alpha1 "github.com/fdymylja/tmos/module/runtime/v1alpha1"
	"github.com/fdymylja/tmos/runtime/authentication"
	"github.com/fdymylja/tmos/runtime/client"
	"github.com/fdymylja/tmos/runtime/errors"
	"github.com/fdymylja/tmos/runtime/meta"
	"github.com/fdymylja/tmos/runtime/module"
	"github.com/fdymylja/tmos/runtime/store/badger"
	"k8s.io/klog/v2"
)

var (
	errEmptyModuleName = errors.New("runtime: empty module name")
)

// NewBuilder creates a new Builder for the Runtime
func NewBuilder() *Builder {
	b := &Builder{
		installedModules: map[string]struct{}{},
		router:           NewRouter(),
		store:            badger.NewStore(),
		rt:               &Runtime{},
	}
	// we already add the core modules in order
	// in theory we could add a dependency system
	// for genesis initialization, but for now lets keep it simple.
	runtimeModule := runtime.NewModule() // needs to be first as it has state transitions/state object info
	rbacModule := rbac.NewModule()       // needs to be second as it provides the authorization layer
	b.rbac = rbacModule                  // we set the rbac module inside so that we can prepare initial genesis with rbac
	abciModule := abci.NewModule()       // abci third so other modules can have access to this information
	b.AddModule(runtimeModule)
	b.AddModule(rbacModule)
	b.AddModule(abciModule)
	// we add the initial external role, with basically no authorization towards no resource.
	b.externalRole = &rbacv1alpha1.Role{
		Id: rbacv1alpha1.ExternalAccountRoleID,
	}
	return b
}

// Builder is used to create a new runtime from scratch
type Builder struct {
	installedModules map[string]struct{} // installedModules is used to check if multiple modules with the same name are being installed
	modules          []module.Descriptor

	externalRole *rbacv1alpha1.Role
	rbac         *rbac.Module
	decoder      authentication.TxDecoder

	router *Router
	store  *badger.Store
	rt     *Runtime
}

// AddModule adds a new module.Module to the list of modules to install
func (b *Builder) AddModule(m module.Module) {
	type subjectSetter interface {
		SetSubject(subject string)
	}

	mc := client.New(newRuntimeAsServer(b.rt))
	descriptor := m.Initialize(mc)
	mc.(subjectSetter).SetSubject(descriptor.Name) // set the authentication name for the module TODO: we should do this a lil better
	b.modules = append(b.modules, descriptor)
}

func (b *Builder) SetDecoder(txDecoder authentication.TxDecoder) {
	b.decoder = txDecoder
}

// Build installs the module.Modules provided and returns a fully functional runtime
func (b *Builder) Build() (*Runtime, error) {
	// install all modules
	for _, m := range b.modules {
		// check if already installed
		if _, exists := b.installedModules[m.Name]; exists {
			return nil, fmt.Errorf("double registration of module named %s", m.Name)
		}
		role, binding, err := b.install(m)
		if err != nil {
			return nil, fmt.Errorf("error while installing module %s: %w", m.Name, err)
		}
		// add initial role to rbac
		b.rbac.AddInitialRole(role, binding)
		// mark as installed module
		b.installedModules[m.Name] = struct{}{}
	}
	// add external role to rbac with no binding
	b.rbac.AddInitialRole(b.externalRole, nil)
	b.rt.store = b.store
	b.rt.router = b.router
	b.rt.modules = b.modules
	b.rt.rbac = b.rbac.AsAuthorizer()
	switch b.decoder {
	case nil:
		klog.Warningf("no decoder - transactions sent to the ABCI application will be rejected")
	default:
		b.rt.authn = b.decoder
	}

	return b.rt, nil
}

func (b *Builder) install(m module.Descriptor) (role *rbacv1alpha1.Role, binding *rbacv1alpha1.RoleBinding, err error) {
	// check name
	if isModuleNameEmpty(m.Name) {
		return nil, nil, errEmptyModuleName
	}
	role = &rbacv1alpha1.Role{
		Id: roleNameForModule(m.Name),
	}
	binding = &rbacv1alpha1.RoleBinding{
		Subject: m.Name,
		RoleRef: roleNameForModule(m.Name),
	}
	// install state transition controllers
	for _, ctrl := range m.StateTransitionControllers {
		// add state transition controller to the router
		err = b.router.AddStateTransitionController(ctrl.StateTransition, ctrl.Controller)
		if err != nil {
			return
		}
		// add deliver rights for the state transition
		err = role.Extend(runtimev1alpha1.Verb_Deliver, ctrl.StateTransition)
		if err != nil {
			return
		}
		// if the state transition is marked as external we extend the external_account role
		if ctrl.External {
			err = b.externalRole.Extend(runtimev1alpha1.Verb_Deliver, ctrl.StateTransition)
			if err != nil {
				return nil, nil, err
			}
		}
		klog.Infof("registered state transition %s for module %s", meta.Name(ctrl.StateTransition), m.Name)
	}

	// register admission controllers
	for _, ctrl := range m.AdmissionControllers {
		err = b.router.AddStateTransitionAdmissionController(ctrl.StateTransition, ctrl.Controller)
		if err != nil {
			return
		}
		klog.Infof("registered admission controller %s for module %s", meta.Name(ctrl.StateTransition), m.Name)
	}

	// register state objects
	for _, so := range m.StateObjects {
		err = b.store.RegisterStateObject(so.StateObject)
		if err != nil {
			return
		}
		err = extendRoleForStateObject(role, so.StateObject)
		if err != nil {
			return
		}
		klog.Infof("registered state object %s for module %s", meta.Name(so.StateObject), m.Name)
	}
	// register dependencies onto other modules
	for _, st := range m.Needs {
		err = role.Extend(runtimev1alpha1.Verb_Deliver, st)
		if err != nil {
			return nil, nil, fmt.Errorf("error while registering module dependency %s: %w", meta.Name(st), err)
		}
	}

	// TODO register admission + mutating admission + hooks
	// register authentication extensions
	if m.AuthenticationExtension == nil {
		return
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
	return
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
