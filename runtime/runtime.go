package runtime

import (
	"errors"
	"fmt"
	"github.com/fdymylja/tmos/runtime/store/badger"

	"github.com/fdymylja/tmos/module/meta"
	"github.com/fdymylja/tmos/runtime/authorization"
	"github.com/fdymylja/tmos/runtime/controller"
)

type Runtime struct {
	authorizer authorization.Authorizer
	router     *Router
	store      *badger.Store
}

func (r *Runtime) Get(object meta.StateObject) error {
	return convertStoreError(r.store.Get(object))
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

func (r *Runtime) Delete(user string, object meta.StateObject) error {
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
