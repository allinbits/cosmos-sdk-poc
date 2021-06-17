package runtime

import (
	"errors"
	"fmt"

	"github.com/fdymylja/tmos/core/meta"
	"github.com/fdymylja/tmos/runtime/authentication"
	"github.com/fdymylja/tmos/runtime/statetransition"
)

var ErrTransitionAlreadyRegistered = errors.New("router: state transition already registered")
var ErrTransitionNotFound = errors.New("router: state transition not found")

func NewRouter() *Router {
	return &Router{
		authAdmissionHandlers:                 nil,
		transactionPostAuthenticationHandlers: nil,

		stateTransitionAdmissionHandlers:     map[string][]statetransition.AdmissionHandler{},
		stateTransitionPreExecutionHandlers:  map[string][]statetransition.PreExecutionHandler{},
		stateTransitionExecutionHandlers:     map[string]statetransition.ExecutionHandler{},
		stateTransitionPostExecutionHandlers: map[string][]statetransition.PostExecutionHandler{},
	}
}

type Router struct {
	authAdmissionHandlers                 []authentication.AdmissionHandler
	transactionPostAuthenticationHandlers []authentication.PostAuthenticationHandler

	stateTransitionAdmissionHandlers     map[string][]statetransition.AdmissionHandler
	stateTransitionPreExecutionHandlers  map[string][]statetransition.PreExecutionHandler
	stateTransitionExecutionHandlers     map[string]statetransition.ExecutionHandler
	stateTransitionPostExecutionHandlers map[string][]statetransition.PostExecutionHandler
}

// Admission handlers

func (r *Router) AddStateTransitionAdmissionHandler(transition meta.StateTransition, handler statetransition.AdmissionHandler) error {
	name := meta.Name(transition)

	if !r.stateTransitionAdmissionHandlerExists(transition) {
		r.stateTransitionAdmissionHandlers[name] = nil
	}

	r.stateTransitionAdmissionHandlers[name] = append(r.stateTransitionAdmissionHandlers[name], handler)
	return nil
}

func (r *Router) GetStateTransitionAdmissionHandlers(transition meta.StateTransition) ([]statetransition.AdmissionHandler, error) {
	if !r.knownStateTransition(transition) {
		return nil, fmt.Errorf(
			"%w: unable to provide state transition admission handlers for unknown state transition %s",
			ErrTransitionNotFound,
			meta.Name(transition),
		)
	}
	ctrls, exists := r.stateTransitionAdmissionHandlers[meta.Name(transition)]
	if !exists {
		return nil, nil
	}
	return ctrls, nil
}

// Pre execution handlers

func (r *Router) AddStateTransitionPreExecutionHandler(transition meta.StateTransition, handler statetransition.PreExecutionHandler) error {
	if !r.knownStateTransition(transition) {
		return fmt.Errorf("%w: unable to register state transition pre execution handler for unknown state transition %s", ErrTransitionNotFound, meta.Name(transition))
	}
	name := meta.Name(transition)
	// initialize slice if it does not exist
	if _, exists := r.stateTransitionPreExecutionHandlers[name]; !exists {
		r.stateTransitionPreExecutionHandlers[name] = nil
	}
	// register handler
	r.stateTransitionPreExecutionHandlers[name] = append(r.stateTransitionPreExecutionHandlers[name], handler)
	return nil
}

func (r *Router) GetStateTransitionPreExecutionHandlers(transition meta.StateTransition) ([]statetransition.PreExecutionHandler, error) {
	name := meta.Name(transition)
	if !r.knownStateTransition(transition) {
		return nil, fmt.Errorf("%w: %s", ErrTransitionNotFound, name)
	}
	preExecHandlers, exists := r.stateTransitionPreExecutionHandlers[name]
	if !exists {
		return nil, nil
	}
	return preExecHandlers, nil
}

// Execution Handlers

func (r *Router) AddStateTransitionExecutionHandler(transition meta.StateTransition, handler statetransition.ExecutionHandler) error {
	name := meta.Name(transition)
	if r.knownStateTransition(transition) {
		return fmt.Errorf("%w: %s", ErrTransitionAlreadyRegistered, name)
	}
	r.stateTransitionExecutionHandlers[name] = handler
	return nil
}

func (r *Router) GetStateTransitionExecutionHandler(transition meta.StateTransition) (statetransition.ExecutionHandler, error) {
	name := meta.Name(transition)
	handler, exists := r.stateTransitionExecutionHandlers[name]
	if !exists {
		return nil, fmt.Errorf("%w: %s", ErrTransitionNotFound, name)
	}
	return handler, nil
}

// Post Execution handlers

func (r *Router) AddStateTransitionPostExecutionHandler(transition meta.StateTransition, handler statetransition.PostExecutionHandler) error {
	name := meta.Name(transition)
	if !r.knownStateTransition(transition) {
		return fmt.Errorf("%w: unable to register state transition post execution handler for unknown state transition %s", ErrTransitionNotFound, name)
	}
	if _, exists := r.stateTransitionPostExecutionHandlers[name]; !exists {
		r.stateTransitionPostExecutionHandlers[name] = nil
	}
	r.stateTransitionPostExecutionHandlers[name] = append(r.stateTransitionPostExecutionHandlers[name], handler)
	return nil
}

func (r *Router) GetStateTransitionPostExecutionHandlers(transition meta.StateTransition) ([]statetransition.PostExecutionHandler, error) {
	name := meta.Name(transition)
	if !r.knownStateTransition(transition) {
		return nil, fmt.Errorf("%w: %s", ErrTransitionNotFound, name)
	}
	postExecHandlers, exists := r.stateTransitionPostExecutionHandlers[name]
	if !exists {
		return nil, nil
	}
	return postExecHandlers, nil
}

func (r *Router) GetAuthAdmissionHandlers() []authentication.AdmissionHandler {
	return r.authAdmissionHandlers
}

func (r *Router) AddAuthAdmissionHandler(ctrl authentication.AdmissionHandler) {
	r.authAdmissionHandlers = append(r.authAdmissionHandlers, ctrl)
}

func (r *Router) GetTransactionPostAuthenticationHandlers() []authentication.PostAuthenticationHandler {
	return r.transactionPostAuthenticationHandlers
}

func (r *Router) AddTransactionPostAuthenticationHandler(ctrl authentication.PostAuthenticationHandler) {
	r.transactionPostAuthenticationHandlers = append(r.transactionPostAuthenticationHandlers, ctrl)
}

func (r *Router) ListStateTransitions() []string {
	sts := make([]string, 0, len(r.stateTransitionExecutionHandlers))
	for st := range r.stateTransitionExecutionHandlers {
		sts = append(sts, st)
	}
	return sts
}

func (r *Router) knownStateTransition(transition meta.StateTransition) bool {
	_, known := r.stateTransitionExecutionHandlers[meta.Name(transition)]
	return known
}

func (r *Router) stateTransitionAdmissionHandlerExists(transition meta.StateTransition) bool {
	name := meta.Name(transition)
	_, exists := r.stateTransitionAdmissionHandlers[name]
	return exists
}
