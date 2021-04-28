package runtime

import (
	"errors"
	"fmt"
	"sync/atomic"

	"github.com/fdymylja/tmos/runtime/authentication"
	"github.com/fdymylja/tmos/runtime/authorization"
	"github.com/fdymylja/tmos/runtime/controller"
	"github.com/fdymylja/tmos/runtime/meta"
	"github.com/fdymylja/tmos/runtime/store/badger"
	"k8s.io/klog/v2"
)

type Runtime struct {
	modules     []*ModuleDescriptor
	initialized uint32

	authenticator authentication.Authenticator
	authorizer    authorization.Authorizer
	router        *Router
	store         *badger.Store
}

// Initialize initializes the runtime with default state from modules which have genesis
func (r *Runtime) Initialize() error {
	const notInitialized uint32 = 0
	const initialized uint32 = 1
	if !atomic.CompareAndSwapUint32(&r.initialized, notInitialized, initialized) {
		return fmt.Errorf("already initialized")
	}
	klog.Infof("initializing default genesis state for modules")
	// iterate through modules and call the genesis
	for _, m := range r.modules {
		klog.Infof("initializing genesis state for %s", m.Name)
		if err := m.Genesis.Handler.SetDefault(); err != nil {
			return fmt.Errorf("runtime: failed genesis initalization for module %s: %w", m.Name, err)
		}
	}
	klog.Infof("default genesis initialization completed")
	return nil
}

func (r *Runtime) Get(id meta.ID, object meta.StateObject) error {
	return convertStoreError(r.store.Get(id, object))
}

func (r *Runtime) List() {
	panic("implement me")
}

func (r *Runtime) Create(user string, object meta.StateObject) error {
	return convertStoreError(r.store.Create(object))
}

func (r *Runtime) Update(user string, object meta.StateObject) error {
	return convertStoreError(r.store.Update(object))
}

func (r *Runtime) Delete(user string, id meta.ID, object meta.StateObject) error {
	return convertStoreError(r.store.Delete(object))
}

func (r *Runtime) Deliver(identities []string, transition meta.StateTransition, skipAdmissionControllers bool) (err error) {
	// identity here should be used for authorization checks
	// ex: identity is module/user then can it call the state transition?
	// TODO

	// run the admission controllers
	if !skipAdmissionControllers {
		err = r.runAdmissionControllers(transition)
		if err != nil {
			return err
		}
	}

	// get the handler
	handler, err := r.router.GetStateTransitionController(transition)
	if err != nil {
		return err
	}

	// deliver the request
	_, err = handler.Deliver(controller.StateTransitionRequest{Transition: transition})
	if err != nil {
		return err
	}

	return nil
}

func (r *Runtime) runAdmissionControllers(transition meta.StateTransition) error {
	ctrls, err := r.router.GetAdmissionControllers(transition)
	if err != nil {
		return fmt.Errorf("unable to execute request: %s", meta.Name(transition))
	}
	for _, ctrl := range ctrls {
		_, err = ctrl.Validate(controller.AdmissionRequest{Transition: transition})
		if err != nil {
			return fmt.Errorf("%w: %s", BadRequest, err.Error())
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
		return fmt.Errorf("%w: %s", NotFound, err)
	default:
		panic("unrecognized error type:" + err.Error())
	}
}
