package handler

import (
	"context"

	"github.com/Icedroid/go-grpc/api/proto"
	services2 "github.com/Icedroid/go-grpc/internal/app/services"

	"github.com/golang/protobuf/ptypes"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type ReviewsServer struct {
	logger  *zap.Logger
	service services2.ReviewsService
}

func NewReviewsServer(logger *zap.Logger, ps services2.ReviewsService) (*ReviewsServer, error) {
	return &ReviewsServer{
		logger:  logger,
		service: ps,
	}, nil
}

func (s *ReviewsServer) Query(ctx context.Context, req *proto.QueryReviewsRequest) (*proto.QueryReviewsResponse, error) {
	rs, err := s.service.Query(req.ProductID)
	if err != nil {
		return nil, errors.Wrap(err, "reviews grpc service get reviews error")
	}

	resp := &proto.QueryReviewsResponse{
		Reviews: make([]*proto.Review, 0, len(rs)),
	}
	for _, r := range rs {
		ct, err := ptypes.TimestampProto(r.CreatedTime)
		if err != nil {
			return nil, errors.Wrap(err, "convert create time error")
		}

		pr := &proto.Review{
			Id:          uint64(r.ID),
			ProductID:   r.ProductID,
			Message:     r.Message,
			CreatedTime: ct,
		}

		resp.Reviews = append(resp.Reviews, pr)
	}

	return resp, nil
}
