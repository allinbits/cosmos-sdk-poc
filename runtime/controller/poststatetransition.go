package controller

import (
	"github.com/fdymylja/tmos/runtime"
)

type PostStateTransitionRequest struct {
	Transition runtime.StateTransition
}

type PostStateTransitionResponse struct {
}

type PostStateTransition interface {
	PostTransitionExecute(PostStateTransitionRequest) (PostStateTransitionResponse, error)
}
