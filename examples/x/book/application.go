package book

import (
	"github.com/fdymylja/cosmos-os/pkg/application"
)

type Application struct {
}

func (a Application) ID() application.ID {
	return "BOOK"
}

func (a Application) RegisterDeliverers(register application.RegisterDelivererFn) {
	register(&MsgRegisterBook{}, handleRegisterBook())
}

func (a Application) RegisterCheckers(register application.RegisterCheckerFn) {
	return
}

func (a Application) RegisterQueriers(register application.RegisterQuerierFn) {
	register(&QueryBookRequest{}, &QueryBookResponse{}, queryBook())
}

func (a Application) InitGenesis() {
}

func (a Application) ExportGenesis() {
}

func (a Application) RegisterInvariant() {
}
