package runtime

import (
	"errors"
	"fmt"

	"github.com/fdymylja/tmos/runtime/controller"
)

var ErrAlreadyRegistered = errors.New("router: transition already registered")
var ErrNotFound = errors.New("router: transition handler not found")

func NewRouter() *Router {
	return &Router{
		stateTransitionHandlers:     map[string]controller.StateTransition{},
		admissionControllerHandlers: map[string][]controller.Admission{}}
}

type Router struct {
	stateTransitionHandlers     map[string]controller.StateTransition
	admissionControllerHandlers map[string][]controller.Admission
}

func (r *Router) AddStateTransitionHandler(transition StateTransition, handler controller.StateTransition) error {
	name := Name(transition)
	if _, exists := r.stateTransitionHandlers[name]; exists {
		return fmt.Errorf("%w: %s", ErrAlreadyRegistered, name)
	}
	r.stateTransitionHandlers[name] = handler
	return nil
}

func (r *Router) GetStateTransitionController(transition StateTransition) (controller.StateTransition, error) {
	name := Name(transition)
	handler, exists := r.stateTransitionHandlers[name]
	if !exists {
		return nil, fmt.Errorf("%w: %s", ErrNotFound, name)
	}
	return handler, nil
}

func (r *Router) AddAdmissionController(transition StateTransition, handler controller.Admission) error {
	name := Name(transition)
	if _, exists := r.admissionControllerHandlers[name]; !exists {
		r.admissionControllerHandlers[name] = nil
	}
	r.admissionControllerHandlers[name] = append(r.admissionControllerHandlers[name], handler)
	return nil
}

func (r *Router) GetAdmissionControllers(transition StateTransition) ([]controller.Admission, error) {
	ctrls, exists := r.admissionControllerHandlers[Name(transition)]
	if !exists {
		return nil, nil
	}
	return ctrls, nil
}
