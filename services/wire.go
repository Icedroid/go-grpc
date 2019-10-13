// +build wireinject

package services

import (
	"github.com/google/wire"
	"github.com/Icedroid/go-grpc/pkg/config"
	"github.com/Icedroid/go-grpc/pkg/database"
	"github.com/Icedroid/go-grpc/pkg/log"
	"github.com/Icedroid/go-grpc/app/reviews/repositories"
)

var testProviderSet = wire.NewSet(
	log.ProviderSet,
	config.ProviderSet,
	database.ProviderSet,
	ProviderSet,
)

func CreateReviewsService(cf string, sto repositories.ReviewsRepository) (ReviewsService, error) {
	panic(wire.Build(testProviderSet))
}
