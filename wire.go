// +build wireinject

package main

import (
	"github.com/Icedroid/go-grpc/handler"
	"github.com/Icedroid/go-grpc/internal/app/services"
	"github.com/Icedroid/go-grpc/internal/pkg/app"
	"github.com/Icedroid/go-grpc/internal/pkg/config"
	"github.com/Icedroid/go-grpc/internal/pkg/database"
	"github.com/Icedroid/go-grpc/internal/pkg/jaeger"
	"github.com/Icedroid/go-grpc/internal/pkg/log"
	"github.com/Icedroid/go-grpc/internal/pkg/transports/grpc"
	"github.com/Icedroid/go-grpc/internal/pkg/transports/http"
	"github.com/Icedroid/go-grpc/repositories"
	"github.com/google/wire"
)

var providerSet = wire.NewSet(
	log.ProviderSet,
	config.ProviderSet,
	database.ProviderSet,
	services.ProviderSet,
	// consul.ProviderSet,
	jaeger.ProviderSet,
	http.ProviderSet,
	grpc.ProviderSet,
	repositories.ProviderSet,
	// controllers.ProviderSet,
	handler.ProviderSet,
)

func CreateApp(cf string) (*app.Application, error) {
	panic(wire.Build(providerSet))
}
