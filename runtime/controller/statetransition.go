package controller

import (
	"github.com/fdymylja/tmos/runtime/authentication"
	"github.com/fdymylja/tmos/runtime/meta"
)

// StateTransitionRequest is the request forwarded to the StateTransition controller
type StateTransitionRequest struct {
	Subjects   *authentication.Subjects
	Transition meta.StateTransition
}

// StateTransitionResponse is the response returned by StateTransition controller
type StateTransitionResponse struct {
}

// StateTransition identifies the state transition controller
type StateTransition interface {
	// Deliver forwards a request for state change
	Deliver(req StateTransitionRequest) (StateTransitionResponse, error)
}

type StateTransitionFn func(req StateTransitionRequest) (resp StateTransitionResponse, err error)

func (s StateTransitionFn) Deliver(req StateTransitionRequest) (StateTransitionResponse, error) {
	return s(req)
}
