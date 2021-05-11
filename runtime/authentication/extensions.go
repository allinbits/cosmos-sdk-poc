package authentication

// AdmissionController is a controller for Authentication requests
// which does checks that MUST NOT change state but can read it
// for example:
//         was the minimum fee reached?
//         do the accounts exist?
type AdmissionController interface {
	Validate(tx Tx) error
}

type DeliverRequest struct {
	Tx Tx
}

type DeliverResponse struct {
}

type TransitionController interface {
	Deliver(DeliverRequest) (DeliverResponse, error)
}