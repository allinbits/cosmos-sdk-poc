package router

import (
	"errors"
	"fmt"

	"github.com/fdymylja/tmos/apis/meta"
	basicctrl "github.com/fdymylja/tmos/pkg/controller/basic"
)

var ErrAlreadyRegistered = errors.New("router: transition already registered")
var ErrNotFound = errors.New("router: transition handler not found")

func NewRouter() *Router {
	return &Router{transitionHandlers: map[string]basicctrl.StateTransitionHandler{}}
}

type Router struct {
	transitionHandlers map[string]basicctrl.StateTransitionHandler
}

func (r *Router) AddHandler(transition meta.StateTransition, handler basicctrl.StateTransitionHandler) error {
	name := meta.Name(transition)
	if _, exists := r.transitionHandlers[name]; exists {
		return fmt.Errorf("%w: %s", ErrAlreadyRegistered, name)
	}
	r.transitionHandlers[name] = handler
	return nil
}

func (r *Router) Handler(transition meta.StateTransition) (basicctrl.StateTransitionHandler, error) {
	name := meta.Name(transition)
	handler, exists := r.transitionHandlers[name]
	if !exists {
		return nil, fmt.Errorf("%w: %s", ErrNotFound, name)
	}
	return handler, nil
}
