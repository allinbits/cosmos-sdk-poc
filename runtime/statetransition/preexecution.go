package statetransition

import "github.com/fdymylja/tmos/runtime/authentication/user"

type PreExecutionRequest struct {
	Users      user.Users
	Transition StateTransition
}

type PreExecutionHandler interface {
	PreExec(req PreExecutionRequest) error
}

type PreExecutionHandlerFunc func(req PreExecutionRequest) error

func (f PreExecutionHandlerFunc) PreExec(req PreExecutionRequest) error {
	return f(req)
}
