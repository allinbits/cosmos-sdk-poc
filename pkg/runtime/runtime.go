package runtime

import (
	"fmt"

	"google.golang.org/protobuf/types/known/anypb"

	runtime "github.com/fdymylja/tmos/apis/core/runtime/v1alpha1"
	tx "github.com/fdymylja/tmos/apis/core/tx/v1alpha1"
	"github.com/fdymylja/tmos/pkg/application"
	"github.com/fdymylja/tmos/pkg/runtime/store"
)

type deliverDescriptor struct {
	appID   string
	handler application.DeliverHandler
}

type checkDescriptor struct {
	appID   string
	handler application.DeliverHandler
}

type beginBlockDescriptor struct {
	appID   string
	handler application.BeginBlockHandler
}

type endBlockDescriptor struct {
	appID   string
	handler application.EndBlockHandler
}

type Runtime struct {
	db *store.Store
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
		return nil, fmt.Errorf("unknown tx: %s", rawTx.Body.TypeUrl)
	}
}

func NewRuntime() *Runtime {
	return &Runtime{
		db: store.NewStore(),
	}
}

func (r *Runtime) Mount(app application.Application) error {
	// register state objects
	app.RegisterStateObjects(r.db)
	// register routes
	app.RegisterHandlers()
}
