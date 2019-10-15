package repository

import (
	"github.com/Icedroid/go-grpc/internal/pkg/model"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type ReviewsRepository interface {
	Query(productID uint64) (p []*model.Review, err error)
}

type MysqlReviewsRepository struct {
	logger *zap.Logger
	db     *gorm.DB
}

func NewMysqlReviewsRepository(logger *zap.Logger, db *gorm.DB) ReviewsRepository {
	return &MysqlReviewsRepository{
		logger: logger.With(zap.String("type", "ReviewsRepository")),
		db:     db,
	}
}

func (s *MysqlReviewsRepository) Query(productID uint64) (rs []*model.Review, err error) {
	if err = s.db.Table("reviews").Where("product_id = ?", productID).Find(&rs).Error; err != nil {
		return nil, errors.Wrapf(err, "get review error[productID=%d]", productID)
	}
	return
}
