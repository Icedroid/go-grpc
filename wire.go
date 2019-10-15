// +build wireinject

package main

import (
	"github.com/google/wire"

	"github.com/Icedroid/go-grpc/internal/app"
	"github.com/Icedroid/go-grpc/internal/app/grpcclient"
	"github.com/Icedroid/go-grpc/internal/app/grpcserver"
	"github.com/Icedroid/go-grpc/internal/app/repository"
	"github.com/Icedroid/go-grpc/internal/app/service"
	pkgapp "github.com/Icedroid/go-grpc/internal/pkg/app"
	"github.com/Icedroid/go-grpc/internal/pkg/config"
	"github.com/Icedroid/go-grpc/internal/pkg/database"
	"github.com/Icedroid/go-grpc/internal/pkg/jaeger"
	"github.com/Icedroid/go-grpc/internal/pkg/log"
	"github.com/Icedroid/go-grpc/internal/pkg/transport/grpc"
	"github.com/Icedroid/go-grpc/internal/pkg/transport/http"
)

var providerSet = wire.NewSet(
	app.ProviderSet,
	log.ProviderSet,
	config.ProviderSet,
	database.ProviderSet,
	service.ProviderSet,
	jaeger.ProviderSet,
	http.ProviderSet,
	grpc.ProviderSet,
	repository.ProviderSet,
	grpcclient.ProviderSet,
	grpcserver.ProviderSet,
)

func CreateApp(appName string) (*pkgapp.Application, error) {
	panic(wire.Build(providerSet))
}
