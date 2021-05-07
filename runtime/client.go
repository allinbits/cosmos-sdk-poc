package runtime

import (
	"github.com/fdymylja/tmos/runtime/authentication"
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

func (s server) Create(subject string, object meta.StateObject) error {
	return s.rt.Create(subject, object)
}

func (s server) Update(subject string, object meta.StateObject) error {
	return s.rt.Update(subject, object)
}

func (s server) Delete(subject string, object meta.StateObject) error {
	return s.rt.Delete(subject, object)
}

func (s server) Deliver(subjects *authentication.Subjects, transition meta.StateTransition) error {
	return s.rt.deliver(subjects, transition)
}
