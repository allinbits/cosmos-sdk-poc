package controller

import "github.com/fdymylja/tmos/apis/meta"

// AdmissionRequest is the request sent to admission controllers
type AdmissionRequest struct {
	// Transition defines the meta.StateTransition that needs to be validated
	Transition meta.StateTransition
}

// AdmissionResponse is the response returned by Admission controler
type AdmissionResponse struct {
}

// Admission defines the admission controller
// its duty is to verify without accessing state
// that the provided state transition is valid
type Admission interface {
	Validate(AdmissionRequest) (AdmissionResponse, error)
}
