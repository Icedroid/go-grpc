package services

import (
	"github.com/pkg/errors"
	"go.uber.org/zap"

	"github.com/Icedroid/go-grpc/internal/app/repositories"
	"github.com/Icedroid/go-grpc/internal/pkg/models"
)

type ReviewsService interface {
	Query(productID uint64) ([]*models.Review, error)
}

type DefaultReviewsService struct {
	logger     *zap.Logger
	Repository repositories.ReviewsRepository
}

func NewReviewService(logger *zap.Logger, Repository repositories.ReviewsRepository) ReviewsService {
	return &DefaultReviewsService{
		logger:     logger.With(zap.String("type", "DefaultReviewsService")),
		Repository: Repository,
	}
}

func (s *DefaultReviewsService) Query(productID uint64) (rs []*models.Review, err error) {
	if rs, err = s.Repository.Query(productID); err != nil {
		return nil, errors.Wrap(err, "get review error")
	}

	return
}
