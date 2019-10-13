// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/Icedroid/go-grpc/app/reviews"
	"github.com/Icedroid/go-grpc/app/reviews/controllers"
	"github.com/Icedroid/go-grpc/app/reviews/grpcservers"
	"github.com/Icedroid/go-grpc/app/reviews/repositories"
	"github.com/Icedroid/go-grpc/pkg/app"
	"github.com/Icedroid/go-grpc/pkg/config"
	"github.com/Icedroid/go-grpc/pkg/database"
	"github.com/Icedroid/go-grpc/pkg/log"
	"github.com/Icedroid/go-grpc/pkg/transports/grpc"
)

var providerSet = wire.NewSet(
	log.ProviderSet,
	config.ProviderSet,
	database.ProviderSet,
	// services.ProviderSet,
	// consul.ProviderSet,
	// jaeger.ProviderSet,
	// http.ProviderSet,
	grpc.ProviderSet,
	reviews.ProviderSet,
	repositories.ProviderSet,
	controllers.ProviderSet,
	grpcservers.ProviderSet,
)


func CreateApp(cf string) (*app.Application, error) {
	panic(wire.Build(providerSet))
}
