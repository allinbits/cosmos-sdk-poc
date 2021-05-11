package statetransition

import (
	"github.com/fdymylja/tmos/runtime/authentication/user"
	"github.com/fdymylja/tmos/runtime/meta"
)

// Request is the request forwarded to the Handler controller
type Request struct {
	Users      user.Users
	Transition meta.StateTransition
}

// Response is the response returned by Handler controller
type Response struct {
}

// Handler identifies the state transition controller
type Handler interface {
	// Deliver forwards a request for state change
	Deliver(req Request) (Response, error)
}

type HandlerFunc func(req Request) (resp Response, err error)

func (s HandlerFunc) Deliver(req Request) (Response, error) {
	return s(req)
}
