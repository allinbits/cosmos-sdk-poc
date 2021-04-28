package controller

import (
	"github.com/fdymylja/tmos/runtime/meta"
)

type MutatingAdmissionRequest struct {
	Transition meta.StateTransition
}

type MutatingAdmissionResponse struct {
}

// MutatingAdmission is a controller that does admission checks on state transitions
// but has actual access to state and can modify state
type MutatingAdmission interface {
	ValidateMutating(MutatingAdmissionRequest) (MutatingAdmissionResponse, error)
}
