// +build wireinject

package grpcservers

import (
	"github.com/google/wire"
	"github.com/Icedroid/go-grpc/pkg/config"
	"github.com/Icedroid/go-grpc/pkg/database"
	"github.com/Icedroid/go-grpc/pkg/log"
	"github.com/Icedroid/go-grpc/app/reviews/services"
)

var testProviderSet = wire.NewSet(
	log.ProviderSet,
	config.ProviderSet,
	database.ProviderSet,
	ProviderSet,
)

func CreateReviewsServer(cf string, service services.ReviewsService) (*ReviewsServer, error) {
	panic(wire.Build(testProviderSet))
}