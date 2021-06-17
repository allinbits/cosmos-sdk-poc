package runtime

import (
	"encoding/json"
	"fmt"

	meta "github.com/fdymylja/tmos/core/meta"
	runtimev1alpha1 "github.com/fdymylja/tmos/core/runtime/v1alpha1"
	"github.com/fdymylja/tmos/runtime/authentication/user"
	"github.com/fdymylja/tmos/runtime/errors"
	"github.com/fdymylja/tmos/runtime/orm"
	"k8s.io/klog/v2"

	"github.com/fdymylja/tmos/runtime/authentication"
	"github.com/fdymylja/tmos/runtime/authorization"
	"github.com/fdymylja/tmos/runtime/module"
	"github.com/fdymylja/tmos/runtime/statetransition"
)

type deliverOptions struct {
	skipAdmissionHandler bool
}

type DeliverOption func(opt *deliverOptions)

func DeliverSkipAdmissionHandlers() DeliverOption {
	return func(opt *deliverOptions) {
		opt.skipAdmissionHandler = true
	}
}

type Runtime struct {
	initialized bool
	user        user.Users // user uniquely identifies Runtime as a user.User in the system

	modules []module.Descriptor

	txDecoder authentication.TxDecoder

	authorizer  authorization.Authorizer
	rbacEnabled bool

	router *Router
	store  orm.Store

	services *ServiceGroup
}

func (r *Runtime) EnableRBAC() {
	r.rbacEnabled = true
}

func (r *Runtime) DisableRBAC() {
	r.rbacEnabled = false
}

// InitGenesis initializes the runtime with default state from modules which have genesis
func (r *Runtime) InitGenesis() error {
	// check if runtime is initialized
	if r.initialized {
		return fmt.Errorf("runtime: already initialized")
	}
	// initialize the initial runtime components information
	// so that modules such as Authorizer can have access to it.
	klog.Infof("initializing runtime handler default state")
	var modules []*runtimev1alpha1.ModuleDescriptor
	for _, m := range r.modules {
		stateObjects := make([]*runtimev1alpha1.StateObject, len(m.StateObjects))
		for i, so := range m.StateObjects {
			stateObjects[i] = &runtimev1alpha1.StateObject{
				ApiDefinition:    so.StateObject.APIDefinition(),
				ProtobufFullname: (string)(so.StateObject.ProtoReflect().Descriptor().FullName()),
				SchemaDefinition: &runtimev1alpha1.SchemaDefinition{
					Singleton:     so.Options.Singleton,
					PrimaryKey:    so.Options.PrimaryKey,
					SecondaryKeys: so.Options.SecondaryKeys,
				},
			}
		}
		stateTransitions := make([]*runtimev1alpha1.StateTransition, len(m.StateTransitionExecutionHandlers))
		for i, st := range m.StateTransitionExecutionHandlers {
			stateTransitions[i] = &runtimev1alpha1.StateTransition{
				ApiDefinition:    st.StateTransition.APIDefinition(),
				ProtobufFullname: (string)(st.StateTransition.ProtoReflect().Descriptor().FullName()),
			}
		}

		modules = append(modules, &runtimev1alpha1.ModuleDescriptor{
			Name:             m.Name,
			StateObjects:     stateObjects,
			StateTransitions: stateTransitions,
		})
	}

	klog.Infof("initializing default genesis state for modules")
	err := r.deliver(r.user, &runtimev1alpha1.CreateModuleDescriptors{Modules: modules})
	if err != nil {
		return err
	}
	// iterate through modules and call the genesis
	for _, m := range r.modules {
		if m.GenesisHandler == nil {
			continue
		}
		klog.Infof("initializing genesis state for %s", m.Name)
		if err := m.GenesisHandler.Default(); err != nil {
			return fmt.Errorf("runtime: failed genesis initalization for module %s: %w", m.Name, err)
		}
	}
	klog.Infof("default genesis initialization completed")

	r.initialized = true
	klog.Infof("initializing extensions...")

	err = r.services.Start(r.store) // TODO(fdymylja): Runtime should be aware when to start services even on node re-start.
	if err != nil {
		return err
	}
	return nil
}

func (r *Runtime) Import(stateBytes []byte) error {
	genesisData := make(map[string]json.RawMessage)
	err := json.Unmarshal(stateBytes, &genesisData)
	if err != nil {
		return err
	}

	for _, m := range r.modules {
		if m.GenesisHandler == nil {
			continue
		}

		err := m.GenesisHandler.Import(genesisData[m.Name])
		if err != nil {
			return err
		}
	}

	klog.Infof("%v", genesisData)

	return nil
}

func (r *Runtime) Get(id meta.ID, object meta.StateObject) error {
	return convertStoreError(r.store.Get(id, object))
}

func (r *Runtime) List(object meta.StateObject, options orm.ListOptions) (orm.Iterator, error) {
	iter, err := r.store.List(object, options)
	return iter, convertStoreError(err)
}

func (r *Runtime) Create(users user.Users, object meta.StateObject) error {
	if err := r.authorized(
		authorization.NewAttributes(runtimev1alpha1.Verb_Create, object, users),
	); err != nil {
		return err
	}
	return convertStoreError(r.store.Create(object))
}

func (r *Runtime) Update(users user.Users, object meta.StateObject) error {
	if err := r.authorized(
		authorization.NewAttributes(runtimev1alpha1.Verb_Update, object, users),
	); err != nil {
		return err
	}
	return convertStoreError(r.store.Update(object))
}

func (r *Runtime) Delete(users user.Users, object meta.StateObject) error {
	if err := r.authorized(
		authorization.NewAttributes(runtimev1alpha1.Verb_Delete, object, users),
	); err != nil {
		return err
	}
	return convertStoreError(r.store.Delete(object))
}

func (r *Runtime) Deliver(subjects user.Users, transition meta.StateTransition, opts ...DeliverOption) (err error) {
	return r.deliver(subjects, transition, opts...)
}

// deliver delivers a meta.StateTransition to the handler
// returns error in case of routing errors or execution errors.
func (r *Runtime) deliver(users user.Users, stateTransition meta.StateTransition, opts ...DeliverOption) (err error) {
	deliverOpt := new(deliverOptions)
	for _, opt := range opts {
		opt(deliverOpt)
	}
	if !deliverOpt.skipAdmissionHandler {
		err := r.runAdmissionChain(users, stateTransition)
		if err != nil {
			return err
		}
	}
	// identity here should be used for authorization checks
	// ex: identity is module/user then can it call the state transition?
	if err = r.authorized(authorization.NewAttributes(runtimev1alpha1.Verb_Deliver, stateTransition, users)); err != nil {
		return err
	}
	// get the handler
	handler, err := r.router.GetStateTransitionExecutionHandler(stateTransition)
	if err != nil {
		return err
	}
	// deliver the request
	_, err = handler.Exec(statetransition.ExecutionRequest{
		Users:      users,
		Transition: stateTransition,
	})
	if err != nil {
		return err
	}

	return nil
}

// runAdmissionChain runs the AdmissionHandler handlers related to the
// provided state transition.
func (r *Runtime) runAdmissionChain(users user.Users, transition meta.StateTransition) error {
	ctrls, err := r.router.GetStateTransitionAdmissionHandlers(transition)
	if err != nil {
		return fmt.Errorf("unable to execute request %s: %w", meta.Name(transition), err)
	}
	for _, ctrl := range ctrls {
		err = ctrl.Validate(statetransition.AdmissionRequest{
			Transition: transition,
			Users:      users,
		})
		if err != nil {
			return fmt.Errorf("%w: %s", errors.ErrBadRequest, err.Error())
		}
	}
	return nil
}

// runTxAdmissionChain runs the authentication.AdmissionHandler handlers
func (r *Runtime) runTxAdmissionChain(tx authentication.Tx) error {
	ctrls := r.router.GetAuthAdmissionHandlers()
	for _, ctrl := range ctrls {
		err := ctrl.Validate(tx)
		if err != nil {
			return fmt.Errorf("%w: %s", errors.ErrBadRequest, err)
		}
	}
	return nil
}

func (r *Runtime) runTxPostAuthenticationChain(tx authentication.Tx) error {
	ctrls := r.router.GetTransactionPostAuthenticationHandlers()
	for _, ctrl := range ctrls {
		_, err := ctrl.Exec(authentication.PostAuthenticationRequest{Tx: tx})
		if err != nil {
			return fmt.Errorf("%w: %s", errors.ErrBadRequest, err)
		}
	}
	return nil
}

func (r *Runtime) authorized(attributes authorization.Attributes) error {
	if !r.rbacEnabled {
		return nil
	}
	decision, err := r.authorizer.Authorize(attributes)
	if err == nil && decision == authorization.DecisionAllow {
		return nil
	}
	return fmt.Errorf("%w: %s", errors.ErrUnauthorized, err)
}

// convertStoreError converts the store error to a runtime error
func convertStoreError(err error) error {
	if err == nil {
		return nil
	}
	switch {
	case errors.Is(err, orm.ErrNotFound):
		return fmt.Errorf("%w: %s", errors.ErrNotFound, err)
	default:
		panic("unrecognized error type: " + err.Error())
	}
}
