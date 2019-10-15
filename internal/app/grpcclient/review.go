package grpcclient

import (
	"fmt"

	"github.com/pkg/errors"

	"github.com/Icedroid/go-grpc/api/proto"
	"github.com/Icedroid/go-grpc/internal/pkg/transport/grpc"
)

func NewReviewsClient(client *grpc.Client, o *grpc.ServerOptions) (proto.ReviewsClient, error) {
	conn, err := client.Dial(fmt.Sprint(":%s", o.Port))
	if err != nil {
		return nil, errors.Wrap(err, "detail client dial error")
	}
	c := proto.NewReviewsClient(conn)

	return c, nil
}
