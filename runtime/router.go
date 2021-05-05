package runtime

import (
	"errors"
	"fmt"

	"github.com/fdymylja/tmos/runtime/authentication"
	"github.com/fdymylja/tmos/runtime/controller"
	"github.com/fdymylja/tmos/runtime/meta"
)

var ErrTransitionAlreadyRegistered = errors.New("router: state transition already registered")
var ErrTransitionNotFound = errors.New("router: state transition not found")

func NewRouter() *Router {
	return &Router{
		stateTransitionControllers:          map[string]controller.StateTransition{},
		stateTransitionAdmissionControllers: map[string][]controller.Admission{}}
}

type Router struct {
	transactionAdmissionControllers          []authentication.AdmissionController
	transactionPostAuthenticationControllers []authentication.TransitionController
	stateTransitionControllers               map[string]controller.StateTransition
	stateTransitionAdmissionControllers      map[string][]controller.Admission
}

func (r *Router) AddStateTransitionController(transition meta.StateTransition, handler controller.StateTransition) error {
	name := meta.Name(transition)
	if _, exists := r.stateTransitionControllers[name]; exists {
		return fmt.Errorf("%w: %s", ErrTransitionAlreadyRegistered, name)
	}
	r.stateTransitionControllers[name] = handler
	return nil
}

func (r *Router) GetStateTransitionController(transition meta.StateTransition) (controller.StateTransition, error) {
	name := meta.Name(transition)
	handler, exists := r.stateTransitionControllers[name]
	if !exists {
		return nil, fmt.Errorf("%w: %s", ErrTransitionNotFound, name)
	}
	return handler, nil
}

func (r *Router) AddStateTransitionAdmissionController(transition meta.StateTransition, handler controller.Admission) error {
	name := meta.Name(transition)
	if _, exists := r.stateTransitionAdmissionControllers[name]; !exists {
		r.stateTransitionAdmissionControllers[name] = nil
	}
	r.stateTransitionAdmissionControllers[name] = append(r.stateTransitionAdmissionControllers[name], handler)
	return nil
}

func (r *Router) GetAdmissionControllers(transition meta.StateTransition) ([]controller.Admission, error) {
	ctrls, exists := r.stateTransitionAdmissionControllers[meta.Name(transition)]
	if !exists {
		return nil, nil
	}
	return ctrls, nil
}

func (r *Router) GetTransactionAdmissionControllers() []authentication.AdmissionController {
	return r.transactionAdmissionControllers
}

func (r *Router) AddTransactionAdmissionController(ctrl authentication.AdmissionController) {
	r.transactionAdmissionControllers = append(r.transactionAdmissionControllers, ctrl)
}

func (r *Router) GetTransactionPostAuthenticationControllers() []authentication.TransitionController {
	return r.transactionPostAuthenticationControllers
}

func (r *Router) AddTransactionPostAuthenticationController(ctrl authentication.TransitionController) {
	r.transactionPostAuthenticationControllers = append(r.transactionPostAuthenticationControllers, ctrl)
}

func (r *Router) ListStateTransitions() []string {
	sts := make([]string, len(r.stateTransitionControllers))
	for st := range r.stateTransitionControllers {
		sts = append(sts, st)
	}
	return sts
}
