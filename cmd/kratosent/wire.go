//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"kratosent/internal/biz"
	"kratosent/internal/conf"
	"kratosent/internal/data"
	"kratosent/internal/middleware"
	"kratosent/internal/server"
	"kratosent/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, *conf.Register, *conf.MidConfig, *conf.Middleware, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp, middleware.ProviderSet))
}
