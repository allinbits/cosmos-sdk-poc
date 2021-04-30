package authentication

// AdmissionController is a controller for Authentication requests
// which does checks that MUST NOT change state but can read it
// for example:
//         was the minimum fee reached?
//         do the accounts exist?
type AdmissionController interface {
	Validate(ValidateRequest) (ValidateResponse, error)
}

type ValidateRequest struct {
	Tx Tx
}

type ValidateResponse struct{}

type DeliverRequest struct {
}

type DeliverResponse struct {
}

type TransitionController interface {
	Deliver(DeliverRequest) (DeliverResponse, error)
}
