package runtime

import (
	"github.com/fdymylja/tmos/apis/meta"
	"github.com/fdymylja/tmos/pkg/controller"
	"github.com/fdymylja/tmos/pkg/runtime/orm"
	"github.com/fdymylja/tmos/pkg/runtime/router"
)

type Runtime struct {
	router *router.Router
	store  *orm.Store
}

func (r *Runtime) Get(object meta.StateObject) error {
	return r.store.Get(object)
}

func (r *Runtime) List() {
	panic("implement me")
}

func (r *Runtime) Create(user string, object meta.StateObject) error {
	return r.store.Create(object)
}

func (r *Runtime) Update(user string, object meta.StateObject) error {
	return r.store.Update(object)
}

func (r *Runtime) Delete(user string, object meta.StateObject) error {
	return r.store.Delete(object)
}

func (r *Runtime) Deliver(identities []string, transition meta.StateTransition) error {
	// identity here should be used for authorization checks
	// ex: identity is module/user then can it call the state transition?
	// TODO

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
