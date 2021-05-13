package statetransition

import "github.com/fdymylja/tmos/runtime/authentication/user"

type PostExecutionHandler interface {
	PostExec(req PostExecutionRequest) error
}

type PostExecutionRequest struct {
	Users      user.Users
	Transition StateTransition
	Response   ExecutionResponse
}

type PostExecutionHandlerFunc func(req PostExecutionRequest) error

func (f PostExecutionHandlerFunc) PostExec(req PostExecutionRequest) error {
	return f(req)
}
