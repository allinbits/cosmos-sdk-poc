package controller

import (
	"github.com/fdymylja/tmos/runtime"
)

// StateTransitionRequest is the request forwarded to the StateTransition controller
type StateTransitionRequest struct {
	Transition runtime.StateTransition
}

// StateTransitionResponse is the response returned by StateTransition controller
type StateTransitionResponse struct {
}

// StateTransition identifies the state transition controller
type StateTransition interface {
	// Deliver forwards a request for state change
	Deliver(req StateTransitionRequest) (StateTransitionResponse, error)
}