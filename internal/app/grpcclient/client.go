package grpcclient

import (
	"context"

	"github.com/google/wire"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"

	"github.com/Icedroid/go-grpc/api/proto"
	transporthttp "github.com/Icedroid/go-grpc/internal/pkg/transport/http"
)

func CreateInitHttpRoutersFn(
	pc proto.ReviewsClient,
) transporthttp.InitRouter {
	return func(mux *runtime.ServeMux) {
		proto.RegisterReviewsHandlerClient(context.TODO(), mux, pc)
	}
}

var ProviderSet = wire.NewSet(NewReviewsClient, CreateInitHttpRoutersFn)
