package runtime

import (
	"github.com/fdymylja/cosmos-os/pkg/application"
)

type deliverer struct {
	do            application.DeliverFunc
	applicationID application.ID
}

type checker struct {
	do            application.CheckFunc
	applicationID application.ID
}

type querier struct {
	do            application.QueryFunc
	applicationID application.ID
}


