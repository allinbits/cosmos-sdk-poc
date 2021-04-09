package runtime

import (
	"fmt"

	"github.com/fdymylja/tmos/pkg/application"
	"github.com/fdymylja/tmos/pkg/runtime/router"
	"github.com/fdymylja/tmos/pkg/runtime/store"
)

func newExecutionContext(router *router.Router, db *store.Store) *executionContext {
	return &executionContext{
		router: router,
		store:  db,
	}
}

// executionContext contains information for context execution
type executionContext struct {
	router *router.Router
	store  *store.Store
}

func (e executionContext) Get(object application.StateObject) (exists bool) {
	return e.store.Get(object)
}

func (e executionContext) Set(object application.StateObject) {
	e.store.Set(object)
}

func (e executionContext) Delete(object application.StateObject) {
	e.store.Delete(object)
}

func (e executionContext) Deliver(object application.StateTransitionObject) error {
	deliverHandler := e.router.DeliverHandlerFor(object)
	if deliverHandler == nil {
		return fmt.Errorf("state transition handler not found for type: %T", object)
	}
	return deliverHandler.Deliver(application.DeliverRequest{
		StateTransitionObject: object,
		Client:                newExecutionContext(e.router, e.store),
	})
}
