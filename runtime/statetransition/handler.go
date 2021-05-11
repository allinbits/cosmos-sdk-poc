package statetransition

import (
	"github.com/fdymylja/tmos/runtime/authentication/user"
	"github.com/fdymylja/tmos/runtime/meta"
)

// Request is the request forwarded to the Handler controller
type Request struct {
	// Users contains information on the entities that have authorized the meta.StateTransition
	Users user.Users
	// Transition is the meta.StateTransition that needs to be processed
	Transition meta.StateTransition
}

// Response is the response returned by Handler controller
type Response struct {
}

// Handler identifies the state transition controller
type Handler interface {
	// Deliver is called when the Request needs to be processed
	Deliver(req Request) (Response, error)
}

type HandlerFunc func(req Request) (resp Response, err error)

func (s HandlerFunc) Deliver(req Request) (Response, error) {
	return s(req)
}
