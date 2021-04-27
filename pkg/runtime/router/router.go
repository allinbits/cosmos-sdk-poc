package router

import (
	"errors"
	"fmt"

	"github.com/fdymylja/tmos/apis/meta"
	"github.com/fdymylja/tmos/pkg/controller"
)

var ErrAlreadyRegistered = errors.New("router: transition already registered")
var ErrNotFound = errors.New("router: transition handler not found")

func NewRouter() *Router {
	return &Router{stateTransitionHandlers: map[string]controller.StateTransition{}}
}

type Router struct {
	stateTransitionHandlers map[string]controller.StateTransition
}

func (r *Router) AddStateTransitionHandler(transition meta.StateTransition, handler controller.StateTransition) error {
	name := meta.Name(transition)
	if _, exists := r.stateTransitionHandlers[name]; exists {
		return fmt.Errorf("%w: %s", ErrAlreadyRegistered, name)
	}
	r.stateTransitionHandlers[name] = handler
	return nil
}

func (r *Router) GetStateTransitionController(transition meta.StateTransition) (controller.StateTransition, error) {
	name := meta.Name(transition)
	handler, exists := r.stateTransitionHandlers[name]
	if !exists {
		return nil, fmt.Errorf("%w: %s", ErrNotFound, name)
	}
	return handler, nil
}
