//go:build wireinject
// +build wireinject

package main

import (
	"valsea/src/config"
	"valsea/src/data"
	"valsea/src/handler"
	"valsea/src/server"
	"valsea/src/service"

	"github.com/google/wire"
)

type App struct {
	Config *config.Config
	Server *server.Server
}

func BuildApp() (*App, error) {
	wire.Build(
		config.NewConfig,
		server.NewServer,
		handler.NewAccount,
		handler.NewTransfer,
		service.NewAccount,
		service.NewTransfer,
		data.NewRepository,
		wire.Struct(new(App), "*"),
	)
	return &App{}, nil
}
