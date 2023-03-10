// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"kratosent/internal/biz"
	"kratosent/internal/conf"
	"kratosent/internal/data"
	"kratosent/internal/middleware"
	"kratosent/internal/server"
	"kratosent/internal/service"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, register *conf.Register, midConfig *conf.MidConfig, confMiddleware *conf.Middleware, logger log.Logger) (*kratos.App, func(), error) {
	client, cleanup, err := data.NewEntCli(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	dataData, cleanup2, err := data.NewData(client, confData, logger)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	greeterRepo := data.NewGreeterRepo(dataData, logger)
	greeterUsecase := biz.NewGreeterUsecase(greeterRepo, logger)
	greeterService := service.NewGreeterService(greeterUsecase)
	departmenterRepo := data.NewDepartmentRepo(dataData, logger)
	departmentUsecase := biz.NewDepartmentUsecase(departmenterRepo, logger)
	departmentService := service.NewDepartmentService(departmentUsecase)
	roleRepo := data.NewRoleRepo(dataData, logger)
	roleUsecase := biz.NewRoleUsecase(roleRepo, logger)
	roleService := service.NewRoleService(roleUsecase)
	v := service.NewGRPCServiceSlice(greeterService, departmentService, roleService)
	histogramVec := middleware.NewPrometheusHistogramVec(confMiddleware)
	counterVec := middleware.NewPrometheusCounterVec(confMiddleware)
	grpcServer := server.NewGRPCServer(confServer, v, histogramVec, counterVec, logger)
	v2 := service.NewHTTPServiceSlice(greeterService, departmentService, roleService)
	httpServer := server.NewHTTPServer(confServer, v2, histogramVec, counterVec, logger)
	app := newApp(logger, grpcServer, httpServer, confData, register, midConfig)
	return app, func() {
		cleanup2()
		cleanup()
	}, nil
}
