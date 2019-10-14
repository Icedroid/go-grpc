package grpcclients

import (
	"github.com/pkg/errors"

	"github.com/Icedroid/go-grpc/api/proto"
	"github.com/Icedroid/go-grpc/internal/pkg/transports/grpc"
)

func NewReviewsClient(client *grpc.Client) (proto.ReviewsClient, error) {
	conn, err := client.Dial("")
	if err != nil {
		return nil, errors.Wrap(err, "detail client dial error")
	}
	c := proto.NewReviewsClient(conn)

	return c, nil
}
