package grpcservers

import (
	"github.com/google/wire"
	"github.com/Icedroid/go-grpc/api/proto"
	"github.com/Icedroid/go-grpc/pkg/transports/grpc"
	stdgrpc "google.golang.org/grpc"
)



func CreateInitServersFn(
	ps *ReviewsServer,
) grpc.InitServers {
	return func(s *stdgrpc.Server) {
		proto.RegisterReviewsServer(s, ps)
	}
}

var ProviderSet = wire.NewSet(NewReviewsServer, CreateInitServersFn)