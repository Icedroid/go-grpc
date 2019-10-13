package services

import (
	"github.com/pkg/errors"
	"github.com/Icedroid/go-grpc/app/reviews/repositories"
	"github.com/Icedroid/go-grpc/pkg/models"
	"go.uber.org/zap"
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
