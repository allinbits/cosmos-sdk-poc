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

	rbac        authorization.Authorizer
	rbacEnabled bool

	router *Router
	store  orm.Store
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
	err := r.deliver(r.user, &runtimev1alpha1.CreateStateObjectsList{StateObjects: r.store.ListRegisteredStateObjects()})
	if err != nil {
		return err
	}
	err = r.deliver(r.user, &runtimev1alpha1.CreateStateTransitionsList{StateTransitions: r.router.ListStateTransitions()})
	if err != nil {
		return err
	}
	klog.Infof("initializing default genesis state for modules")

	// iterate through modules and call the genesis
	for _, m := range r.modules {
		if m.GenesisHandler == nil {
			continue
		}
		klog.Infof("initializing genesis state for %s", m.Name)
		if err := m.GenesisHandler.Default(); err != nil {
			return fmt.Errorf("runtime: failed genesis initalization for core %s: %w", m.Name, err)
		}
	}
	klog.Infof("default genesis initialization completed")

	r.initialized = true
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

func (r *Runtime) List() {
	panic("implement me")
}

func (r *Runtime) Create(users user.Users, object meta.StateObject) error {
	if err := r.authorized(runtimev1alpha1.Verb_Create, object, users); err != nil {
		return err
	}
	return convertStoreError(r.store.Create(object))
}

func (r *Runtime) Update(users user.Users, object meta.StateObject) error {
	if err := r.authorized(runtimev1alpha1.Verb_Update, object, users); err != nil {
		return err
	}
	return convertStoreError(r.store.Update(object))
}

func (r *Runtime) Delete(users user.Users, object meta.StateObject) error {
	if err := r.authorized(runtimev1alpha1.Verb_Delete, object, users); err != nil {
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
	if err = r.authorized(runtimev1alpha1.Verb_Deliver, stateTransition, users); err != nil {
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
	ctrls := r.router.GetTransactionAdmissionHandlers()
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

func (r *Runtime) authorized(verb runtimev1alpha1.Verb, resource meta.APIObject, users user.Users) error {
	if !r.rbacEnabled {
		return nil
	}
	decision, err := r.rbac.Authorize(authorization.Attributes{
		Verb:     verb,
		Resource: resource,
		Users:    users,
	})
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
