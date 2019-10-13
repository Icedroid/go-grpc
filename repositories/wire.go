// +build wireinject

package repositories

import (
	"github.com/google/wire"
	"github.com/Icedroid/go-grpc/pkg/config"
	"github.com/Icedroid/go-grpc/pkg/database"
	"github.com/Icedroid/go-grpc/pkg/log"
)



var testProviderSet = wire.NewSet(
	log.ProviderSet,
	config.ProviderSet,
	database.ProviderSet,
	ProviderSet,
)

func CreateReviewRepository(f string) (ReviewsRepository, error) {
	panic(wire.Build(testProviderSet))
}

