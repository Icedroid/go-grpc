package app

import (
	"github.com/google/wire"
	"github.com/pkg/errors"
	"go.uber.org/zap"

	"github.com/Icedroid/go-grpc/internal/pkg/app"
	"github.com/Icedroid/go-grpc/internal/pkg/transport/grpc"
	"github.com/Icedroid/go-grpc/internal/pkg/transport/http"
)

func NewApp(appName string, logger *zap.Logger, hs *http.Server, gs *grpc.Server) (*app.Application, error) {
	a, err := app.New(appName, logger, app.GrpcServerOption(gs), app.HttpServerOption(hs))

	if err != nil {
		return nil, errors.Wrap(err, "new app error")
	}

	return a, nil
}

var ProviderSet = wire.NewSet(NewApp)
