package controller

import (
	"github.com/fdymylja/tmos/runtime/meta"
)

type PostStateTransitionRequest struct {
	Transition meta.StateTransition
}

type PostStateTransitionResponse struct {
}

type PostStateTransition interface {
	PostTransitionExecute(PostStateTransitionRequest) (PostStateTransitionResponse, error)
}
