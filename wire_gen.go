// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	app2 "github.com/Icedroid/go-grpc/internal/app"
	"github.com/Icedroid/go-grpc/internal/app/grpcclient"
	"github.com/Icedroid/go-grpc/internal/app/grpcserver"
	"github.com/Icedroid/go-grpc/internal/app/repository"
	"github.com/Icedroid/go-grpc/internal/app/service"
	"github.com/Icedroid/go-grpc/internal/pkg/app"
	"github.com/Icedroid/go-grpc/internal/pkg/config"
	"github.com/Icedroid/go-grpc/internal/pkg/database"
	"github.com/Icedroid/go-grpc/internal/pkg/jaeger"
	"github.com/Icedroid/go-grpc/internal/pkg/log"
	"github.com/Icedroid/go-grpc/internal/pkg/transport/grpc"
	"github.com/Icedroid/go-grpc/internal/pkg/transport/http"
	"github.com/google/wire"
)

// Injectors from wire.go:

func CreateApp(appName2 string) (*app.Application, error) {
	viper, err := config.New(appName2)
	if err != nil {
		return nil, err
	}
	options, err := log.NewOptions(viper)
	if err != nil {
		return nil, err
	}
	logger, err := log.New(options)
	if err != nil {
		return nil, err
	}
	httpOptions, err := http.NewOptions(viper)
	if err != nil {
		return nil, err
	}
	configuration, err := jaeger.NewConfiguration(viper, logger)
	if err != nil {
		return nil, err
	}
	tracer, err := jaeger.New(configuration)
	if err != nil {
		return nil, err
	}
	clientOptions, err := grpc.NewClientOptions(viper, tracer)
	if err != nil {
		return nil, err
	}
	client, err := grpc.NewClient(clientOptions)
	if err != nil {
		return nil, err
	}
	serverOptions, err := grpc.NewServerOptions(viper)
	if err != nil {
		return nil, err
	}
	reviewsClient, err := grpcclient.NewReviewsClient(client, serverOptions)
	if err != nil {
		return nil, err
	}
	initRouter := grpcclient.CreateInitHttpRoutersFn(reviewsClient)
	server, err := http.New(httpOptions, logger, initRouter)
	if err != nil {
		return nil, err
	}
	databaseOptions, err := database.NewOptions(viper, logger)
	if err != nil {
		return nil, err
	}
	db, err := database.New(databaseOptions)
	if err != nil {
		return nil, err
	}
	reviewsRepository := repository.NewMysqlReviewsRepository(logger, db)
	reviewsService := service.NewReviewService(logger, reviewsRepository)
	reviewsServer, err := grpcserver.NewReviewsServer(logger, reviewsService)
	if err != nil {
		return nil, err
	}
	initServers := grpcserver.CreateInitServersFn(reviewsServer)
	grpcServer, err := grpc.NewServer(serverOptions, logger, initServers, tracer)
	if err != nil {
		return nil, err
	}
	application, err := app2.NewApp(appName2, logger, server, grpcServer)
	if err != nil {
		return nil, err
	}
	return application, nil
}

// wire.go:

var providerSet = wire.NewSet(app2.ProviderSet, log.ProviderSet, config.ProviderSet, database.ProviderSet, service.ProviderSet, jaeger.ProviderSet, http.ProviderSet, grpc.ProviderSet, repository.ProviderSet, grpcclient.ProviderSet, grpcserver.ProviderSet)
