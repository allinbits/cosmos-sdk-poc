package runtime

import (
	"encoding/json"
	"fmt"

	runtimev1alpha1 "github.com/fdymylja/tmos/module/runtime/v1alpha1"
	"github.com/fdymylja/tmos/runtime/errors"
	"k8s.io/klog/v2"

	"github.com/fdymylja/tmos/runtime/authentication"
	"github.com/fdymylja/tmos/runtime/authorization"
	"github.com/fdymylja/tmos/runtime/controller"
	"github.com/fdymylja/tmos/runtime/meta"
	"github.com/fdymylja/tmos/runtime/module"
	"github.com/fdymylja/tmos/runtime/store/badger"
)

type deliverOptions struct {
	skipAdmissionController bool
}

type DeliverOption func(opt *deliverOptions)

func DeliverSkipAdmissionControllers(opt *deliverOptions) {
	opt.skipAdmissionController = true
}

type Runtime struct {
	initialized bool

	modules []*module.Descriptor

	authn authentication.Authenticator

	rbac        authorization.RBAC
	rbacEnabled bool

	router *Router
	store  *badger.Store
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
	// so that modules such as RBAC can have access to it.
	klog.Infof("initializing runtime controller default state")
	err := r.deliver(nil, &runtimev1alpha1.CreateStateObjectsList{StateObjects: r.store.ListRegisteredStateObjects()})
	if err != nil {
		return err
	}
	err = r.deliver(nil, &runtimev1alpha1.CreateStateTransitionsList{StateTransitions: r.router.ListStateTransitions()})
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
		if err := m.GenesisHandler.SetDefault(); err != nil {
			return fmt.Errorf("runtime: failed genesis initalization for module %s: %w", m.Name, err)
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

func (r *Runtime) Create(subject string, object meta.StateObject) error {
	if err := r.authorized(runtimev1alpha1.Verb_Create, object, subject); err != nil {
		return err
	}
	return convertStoreError(r.store.Create(object))
}

func (r *Runtime) Update(subject string, object meta.StateObject) error {
	if err := r.authorized(runtimev1alpha1.Verb_Update, object, subject); err != nil {
		return err
	}
	return convertStoreError(r.store.Update(object))
}

func (r *Runtime) Delete(subject string, object meta.StateObject) error {
	if err := r.authorized(runtimev1alpha1.Verb_Delete, object, subject); err != nil {
		return err
	}
	return convertStoreError(r.store.Delete(object))
}

func (r *Runtime) Deliver(subjects *authentication.Subjects, transition meta.StateTransition, opts ...DeliverOption) (err error) {
	return r.deliver(subjects, transition, opts...)
}

// deliver delivers a meta.StateTransition to the handling controller
// returns error in case of routing errors or execution errors.
func (r *Runtime) deliver(subjects *authentication.Subjects, stateTransition meta.StateTransition, opts ...DeliverOption) (err error) {
	deliverOpt := new(deliverOptions)
	for _, opt := range opts {
		opt(deliverOpt)
	}
	if !deliverOpt.skipAdmissionController {
		err := r.runAdmissionChain(stateTransition)
		if err != nil {
			return err
		}
	}
	// identity here should be used for authorization checks
	// ex: identity is module/user then can it call the state transition?
	if err = r.authorized(runtimev1alpha1.Verb_Deliver, stateTransition, subjects.List()...); err != nil {
		return err
	}
	// execute pre state transitions hooks
	err = r.RunPreStateTransitionHooks(stateTransition)
	// get the handler
	handler, err := r.router.GetStateTransitionController(stateTransition)
	if err != nil {
		return err
	}
	// deliver the request
	_, err = handler.Deliver(controller.StateTransitionRequest{
		Subjects:   subjects,
		Transition: stateTransition,
	})
	if err != nil {
		return err
	}
	// execute post state transition hooks
	err = r.RunPostStateTransitionHooks(stateTransition)
	if err != nil {
		return err
	}

	return nil
}

// runAdmissionChain runs the controller.Admission handlers related to the
// provided state transition.
func (r *Runtime) runAdmissionChain(transition meta.StateTransition) error {
	ctrls, err := r.router.GetAdmissionControllers(transition)
	if err != nil {
		return fmt.Errorf("unable to execute request: %s", meta.Name(transition))
	}
	for _, ctrl := range ctrls {
		_, err = ctrl.Validate(controller.AdmissionRequest{Transition: transition})
		if err != nil {
			return fmt.Errorf("%w: %s", errors.ErrBadRequest, err.Error())
		}
	}
	return nil
}

// runTxAdmissionChain runs the authentication.AdmissionController handlers
func (r *Runtime) runTxAdmissionChain(tx authentication.Tx) error {
	ctrls := r.router.GetTransactionAdmissionControllers()
	for _, ctrl := range ctrls {
		err := ctrl.Validate(tx)
		if err != nil {
			return fmt.Errorf("%w: %s", errors.ErrBadRequest, err)
		}
	}
	return nil
}

func (r *Runtime) runTxPostAuthenticationChain(tx authentication.Tx) error {
	ctrls := r.router.GetTransactionPostAuthenticationControllers()
	for _, ctrl := range ctrls {
		_, err := ctrl.Deliver(authentication.DeliverRequest{Tx: tx})
		if err != nil {
			return fmt.Errorf("%w: %s", errors.ErrBadRequest, err)
		}
	}
	return nil
}

func (r *Runtime) authorized(verb runtimev1alpha1.Verb, resource meta.Type, subjects ...string) error {
	if !r.rbacEnabled {
		return nil
	}
	err := r.rbac.Allowed(verb, resource, subjects...)
	if err == nil {
		return nil
	}
	return fmt.Errorf("%w: %s", errors.ErrUnauthorized, err)
}

func (r *Runtime) RunPreStateTransitionHooks(transition meta.StateTransition) error {
	return nil
}

func (r *Runtime) RunPostStateTransitionHooks(transition meta.StateTransition) error {
	return nil
}

// convertStoreError converts the store error to a runtime error
func convertStoreError(err error) error {
	if err == nil {
		return nil
	}
	switch {
	case errors.Is(err, badger.ErrNotFound):
		return fmt.Errorf("%w: %s", errors.ErrNotFound, err)
	default:
		panic("unrecognized error type:" + err.Error())
	}
}
