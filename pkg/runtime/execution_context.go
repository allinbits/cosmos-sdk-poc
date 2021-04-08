package runtime

import (
	"github.com/fdymylja/tmos/pkg/runtime/router"
	"github.com/fdymylja/tmos/pkg/runtime/store"
)

// executionContext contains information for context execution
type executionContext struct {
	router router.Router
	store  *store.Store
}
