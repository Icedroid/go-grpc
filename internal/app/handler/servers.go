package handler

import (
	"github.com/google/wire"
	"google.golang.org/grpc"

	"github.com/Icedroid/go-grpc/api/proto"
	transportgrpc "github.com/Icedroid/go-grpc/internal/pkg/transports/grpc"
)

func CreateInitServersFn(
	ps *ReviewsServer,
) transportgrpc.InitServers {
	return func(s *grpc.Server) {
		proto.RegisterReviewsServer(s, ps)
	}
}

var ProviderSet = wire.NewSet(NewReviewsServer)
