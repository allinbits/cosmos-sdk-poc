package runtime

import (
	"fmt"

	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/known/anypb"

	runtime "github.com/fdymylja/tmos/apis/core/runtime/v1alpha1"
	tx "github.com/fdymylja/tmos/apis/core/tx/v1alpha1"
	"github.com/fdymylja/tmos/pkg/application"
	"github.com/fdymylja/tmos/pkg/runtime/router"
	"github.com/fdymylja/tmos/pkg/runtime/store"
)

type Runtime struct {
	db     *store.Store
	router *router.Router
}

func (r *Runtime) decodeTx(rawTx *runtime.RawTx) ([]*anypb.Any, error) {
	switch {
	case rawTx.Body.MessageIs(&tx.Tx{}):
		t := new(tx.Tx)
		err := rawTx.Body.UnmarshalTo(t)
		if err != nil {
			return nil, err
		}
		return t.StateTransitions, nil
	default:
		return nil, fmt.Errorf("unknown tx type: %s", rawTx.Body.TypeUrl)
	}
}

func (r *Runtime) routeTransition(rawTransition *anypb.Any) error {
	// get underlying type from the transition TODO(fdymylja): maybe this should be handled in the router? IDK not sure yet.
	transitionType, err := protoregistry.GlobalTypes.FindMessageByURL(rawTransition.TypeUrl)
	if err != nil {
		return err
	}
	transition, ok := transitionType.New().Interface().(application.StateTransitionObject)
	if !ok {
		return fmt.Errorf("message is not a state transition: %s, %T", transitionType.Descriptor().FullName(), transitionType.New())
	}
	// we check if the transition exists in the router, it might come from some other protofile
	// no application can process. The router guarantees that it's going to work.
	handler := r.router.DeliverHandlerFor(transition)
	if handler == nil {
		return fmt.Errorf("unknown state transition for %s", transition.ProtoReflect().Descriptor().FullName())
	}
	// after we've asserted the transition is what we're looking for, we can unmarshal the raw transition bytes
	err = rawTransition.UnmarshalTo(transition)
	if err != nil {
		return err
	}
	// now deliver it
	return handler.Deliver(application.DeliverRequest{
		StateTransitionObject: transition,
		Client:                newExecutionContext(r.router, r.db),
	})
}

func newRuntime() *Runtime {
	return &Runtime{
		db:     store.NewStore(),
		router: router.NewRouter(),
	}
}

func (r *Runtime) mount(app application.Application) error {
	// register state objects
	app.RegisterStateObjects(r.db)
	// register routes
	app.RegisterHandlers(r.router)
	return nil
}
