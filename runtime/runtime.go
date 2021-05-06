package runtime

import (
	"encoding/json"
	"errors"
	"fmt"
	"sync/atomic"

	runtimev1alpha1 "github.com/fdymylja/tmos/module/runtime/v1alpha1"
	errors2 "github.com/fdymylja/tmos/runtime/errors"
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
	modules     []*module.Descriptor
	initialized uint32

	authn  authentication.Authenticator
	authz  authorization.Authorizer
	router *Router
	store  *badger.Store
}

// InitGenesis initializes the runtime with default state from modules which have genesis
func (r *Runtime) InitGenesis() error {
	// check if runtime is initialized
	const notInitialized uint32 = 0
	const initialized uint32 = 1
	if !atomic.CompareAndSwapUint32(&r.initialized, notInitialized, initialized) {
		return fmt.Errorf("already initialized")
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
	return convertStoreError(r.store.Create(object))
}

func (r *Runtime) Update(subject string, object meta.StateObject) error {
	return convertStoreError(r.store.Update(object))
}

func (r *Runtime) Delete(subject string, id meta.ID, object meta.StateObject) error {
	return convertStoreError(r.store.Delete(object))
}

func (r *Runtime) Deliver(subjects *authentication.Subjects, transition meta.StateTransition, opts ...DeliverOption) (err error) {
	// identity here should be used for authorization checks
	// ex: identity is module/user then can it call the state transition?
	// TODO
	return r.deliver(subjects, transition, opts...)
}

// deliver delivers a meta.StateTransition to the handling controller
// returns error in case of routing errors or execution errors.
func (r *Runtime) deliver(subjects *authentication.Subjects, stateTransition meta.StateTransition, opts ...DeliverOption) error {
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
			return fmt.Errorf("%w: %s", errors2.ErrBadRequest, err.Error())
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
			return fmt.Errorf("%w: %s", errors2.ErrBadRequest, err)
		}
	}
	return nil
}

func (r *Runtime) runTxPostAuthenticationChain(tx authentication.Tx) error {
	ctrls := r.router.GetTransactionPostAuthenticationControllers()
	for _, ctrl := range ctrls {
		_, err := ctrl.Deliver(authentication.DeliverRequest{Tx: tx})
		if err != nil {
			return fmt.Errorf("%w: %s", errors2.ErrBadRequest, err)
		}
	}
	return nil
}

// convertStoreError converts the store error to a runtime error
func convertStoreError(err error) error {
	if err == nil {
		return nil
	}
	switch {
	case errors.Is(err, badger.ErrNotFound):
		return fmt.Errorf("%w: %s", errors2.ErrNotFound, err)
	default:
		panic("unrecognized error type:" + err.Error())
	}
}
