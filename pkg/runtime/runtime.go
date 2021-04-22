package runtime

import (
	"github.com/fdymylja/tmos/apis/meta"
	basicctrl "github.com/fdymylja/tmos/pkg/controller/basic"
	"github.com/fdymylja/tmos/pkg/runtime/orm"
	"github.com/fdymylja/tmos/pkg/runtime/router"
)

type Runtime struct {
	router *router.Router
	store  *orm.Store
}

func (r *Runtime) Deliver(identities []string, transition meta.StateTransition) error {
	// identity here should be used for authorization checks
	// ex: identity is module/user then can it call the state transition?
	// TODO

	// get the handler
	handler, err := r.router.Handler(transition)
	if err != nil {
		return err
	}
	// deliver the request
	resp := handler.Deliver(basicctrl.DeliverRequest{StateTransition: transition})
	if resp.Error != nil {
		return err
	}
	return nil
}
