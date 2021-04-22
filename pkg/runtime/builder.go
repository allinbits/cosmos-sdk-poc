package runtime

import (
	"fmt"

	"github.com/fdymylja/tmos/apis/meta"
	basicctrl "github.com/fdymylja/tmos/pkg/controller/basic"
	"github.com/fdymylja/tmos/pkg/runtime/client"
	"github.com/fdymylja/tmos/pkg/runtime/orm"
	"github.com/fdymylja/tmos/pkg/runtime/router"
	"k8s.io/klog/v2"
)

// NewBuilder creates a new Builder for the Runtime
func NewBuilder() *Builder {
	return &Builder{
		registeredApps: make(map[string]struct{}),
		router:         router.NewRouter(),
		store:          orm.NewStore(),
		rt:             &Runtime{},
	}
}

// Builder is used to create a new runtime from scratch
type Builder struct {
	registeredApps map[string]struct{}
	router         *router.Router
	store          *orm.Store
	rt             *Runtime
}

// MountApplication mounts a new basic onto the Runtime
func (b *Builder) MountApplication(app basicctrl.Controller) {
	identity := app.Name()
	if _, exists := b.registeredApps[identity]; exists {
		panic(fmt.Errorf("basic already registered: %s", identity))
	}

	appClient := client.NewClient(identity, b.store, b.rt)
	app.RegisterStateTransitions(appClient, func(transition meta.StateTransition, handler basicctrl.StateTransitionHandler) {
		err := b.router.AddHandler(transition, handler)
		if err != nil {
			panic(err)
		}
		klog.Infof("registered state transition %s owned by %s", meta.Name(transition), identity)
	})

	app.RegisterStateObjects(func(object meta.StateObject) {
		err := b.store.RegisterStateObject(app.Name(), object)
		if err != nil {
			panic(err)
		}
		klog.Infof("registered state object %s owned by %s", meta.Name(object), app.Name())
	})

	b.registeredApps[identity] = struct{}{}
}

// Build returns the built *Runtime
func (b *Builder) Build() *Runtime {
	b.rt.store = b.store
	b.rt.router = b.router
	return b.rt
}
