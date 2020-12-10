package author

import (
	"github.com/fdymylja/cosmos-os/pkg/application"
)

type Application struct {
}

func (a Application) ID() application.ID {
	return "AUTHOR"
}

func (a Application) RegisterDeliverers(register application.RegisterDelivererFn) {
	register(&MsgRegisterAuthor{}, handleRegisterAuthor())
}

func (a Application) RegisterCheckers(register application.RegisterCheckerFn) {
}

func (a Application) RegisterQueriers(register application.RegisterQuerierFn) {
	register(&QueryAuthorRequest{}, &QueryAuthorResponse{}, queryAuthor())
}

func (a Application) InitGenesis() {
}

func (a Application) ExportGenesis() {
}

func (a Application) RegisterInvariant() {
}
