package router

import "github.com/fdymylja/tmos/pkg/application"

// Router provides application.DeliverHandler and application.CheckHandler
type Router struct {
	deliver map[string]application.DeliverHandler
	check   map[string]application.CheckHandler
}

func (r Router) RegisterDeliverHandler(object application.StateTransitionObject, deliverer application.DeliverHandler, checker application.CheckHandler, policy application.HandlerPolicy) {
	panic("implement me")
}

func (r Router) RegisterBeginBlockHandler(handler application.BeginBlockHandler) {
	panic("implement me")
}

func (r Router) RegisterEndBlockHandler(handler application.EndBlockHandler) {
	panic("implement me")
}

func (r Router) RegisterHandlerHook(object application.StateTransitionObject, handler application.HookHandler) {
	panic("implement me")
}
