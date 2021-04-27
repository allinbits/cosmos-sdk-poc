package controller

import "github.com/fdymylja/tmos/apis/meta"

type PostStateTransitionRequest struct {
	Transition meta.StateTransition
}

type PostStateTransitionResponse struct {
}

type PostStateTransition interface {
	PostTransitionExecute(PostStateTransitionRequest) (PostStateTransitionResponse, error)
}
