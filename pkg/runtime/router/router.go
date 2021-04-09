package router

import (
	"fmt"

	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/fdymylja/tmos/pkg/application"
)

func NewRouter() *Router {
	return &Router{
		deliver: make(map[protoreflect.FullName]application.DeliverHandler),
		check:   make(map[protoreflect.FullName]application.CheckHandler),
	}
}

// Router provides application.DeliverHandler and application.CheckHandler
type Router struct {
	deliver map[protoreflect.FullName]application.DeliverHandler
	check   map[protoreflect.FullName]application.CheckHandler
}

func (r *Router) DeliverHandlerFor(object application.StateTransitionObject) application.DeliverHandler {
	name := r.getName(object)
	deliver, exists := r.deliver[name]
	if !exists {
		return nil
	}
	return deliver
}

func (r *Router) RegisterDeliverHandler(object application.StateTransitionObject, deliverer application.DeliverHandler, checker application.CheckHandler, policy application.HandlerPolicy) {
	name := r.getName(object)
	_, exists := r.deliver[name]
	if exists {
		panic(fmt.Sprintf("registering multiple handlers for the same state transition object is not allowed: %s", name))
	}
	r.deliver[name] = deliverer
	// tODo checker
}

func (r *Router) RegisterBeginBlockHandler(handler application.BeginBlockHandler) {
	panic("implement me")
}

func (r *Router) RegisterEndBlockHandler(handler application.EndBlockHandler) {
	panic("implement me")
}

func (r *Router) RegisterHandlerHook(object application.StateTransitionObject, handler application.HookHandler) {
	panic("implement me")
}

func (r *Router) getName(object application.StateTransitionObject) protoreflect.FullName {
	return object.ProtoReflect().Descriptor().FullName()
}
