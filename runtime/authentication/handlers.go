package authentication

// AdmissionHandler is a handler for Authentication requests
// which does checks that MUST NOT change state but can read it
// for example:
//         was the minimum fee reached?
//         do the accounts exist?
type AdmissionHandler interface {
	Validate(tx Tx) error
}

type PostAuthenticationRequest struct {
	Tx Tx
}

type PostAuthenticationResponse struct{}

// PostAuthenticationHandler is executed after every AdmissionHandler
// contrary to AdmissionHandler it is allowed to modify state.
// For example:
//		- Deduct fees after a transaction was admitted
//		- Increase account nonce.. etc
type PostAuthenticationHandler interface {
	Exec(PostAuthenticationRequest) (PostAuthenticationResponse, error)
}
