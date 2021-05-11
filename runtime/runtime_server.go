package runtime

import (
	"github.com/fdymylja/tmos/runtime/authentication/user"
	"github.com/fdymylja/tmos/runtime/client"
	"github.com/fdymylja/tmos/runtime/meta"
)

var _ client.RuntimeServer = server{}

func newRuntimeAsServer(rt *Runtime) server {
	return server{
		rt: rt,
	}
}

// server exposes runtime as a client.RuntimeServer
type server struct {
	rt *Runtime
}

func (s server) Get(id meta.ID, object meta.StateObject) error {
	return s.rt.Get(id, object)
}

func (s server) Create(users user.Users, object meta.StateObject) error {
	return s.rt.Create(users, object)
}

func (s server) Update(users user.Users, object meta.StateObject) error {
	return s.rt.Update(users, object)
}

func (s server) Delete(users user.Users, object meta.StateObject) error {
	return s.rt.Delete(users, object)
}

func (s server) Deliver(users user.Users, transition meta.StateTransition) error {
	return s.rt.deliver(users, transition)
}
