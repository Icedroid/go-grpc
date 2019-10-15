package service

import (
	"github.com/pkg/errors"
	"go.uber.org/zap"

	"github.com/Icedroid/go-grpc/internal/app/repository"
	"github.com/Icedroid/go-grpc/internal/pkg/model"
)

type ReviewsService interface {
	Query(productID uint64) ([]*model.Review, error)
}

type DefaultReviewsService struct {
	logger     *zap.Logger
	Repository repository.ReviewsRepository
}

func NewReviewService(logger *zap.Logger, Repository repository.ReviewsRepository) ReviewsService {
	return &DefaultReviewsService{
		logger:     logger.With(zap.String("type", "DefaultReviewsService")),
		Repository: Repository,
	}
}

func (s *DefaultReviewsService) Query(productID uint64) (rs []*model.Review, err error) {
	if rs, err = s.Repository.Query(productID); err != nil {
		return nil, errors.Wrap(err, "get review error")
	}

	return
}
