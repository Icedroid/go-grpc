package app

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/google/wire"
	"github.com/pkg/errors"
	"go.uber.org/zap"

	grpc2 "github.com/Icedroid/go-grpc/internal/pkg/transports/grpc"
	http2 "github.com/Icedroid/go-grpc/internal/pkg/transports/http"
)

type Application struct {
	name       string
	logger     *zap.Logger
	httpServer *http2.Server
	grpcServer *grpc2.Server
}

type Option func(app *Application) error

func HttpServerOption(svr *http2.Server) Option {
	return func(app *Application) error {
		svr.Application(app.name)
		app.httpServer = svr

		return nil
	}
}

func GrpcServerOption(svr *grpc2.Server) Option {
	return func(app *Application) error {
		svr.Application(app.name)
		app.grpcServer = svr
		return nil
	}
}

func New(name string, logger *zap.Logger, options ...Option) (*Application, error) {
	app := &Application{
		name:   name,
		logger: logger.With(zap.String("type", "Application")),
	}

	for _, option := range options {
		if err := option(app); err != nil {
			return nil, err
		}
	}

	return app, nil
}

func (a *Application) Start() error {
	if a.httpServer != nil {
		if err := a.httpServer.Start(); err != nil {
			return errors.Wrap(err, "http server start error")
		}
	}

	if a.grpcServer != nil {
		if err := a.grpcServer.Start(); err != nil {
			return errors.Wrap(err, "grpc server start error")
		}
	}

	return nil
}

func (a *Application) AwaitSignal() {
	c := make(chan os.Signal, 1)
	signal.Reset(syscall.SIGTERM, syscall.SIGINT)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)
	select {
	case s := <-c:
		a.logger.Info("receive a signal", zap.String("signal", s.String()))
		if a.httpServer != nil {
			if err := a.httpServer.Stop(); err != nil {
				a.logger.Warn("stop http server error", zap.Error(err))
			}
		}

		if a.grpcServer != nil {
			if err := a.grpcServer.Stop(); err != nil {
				a.logger.Warn("stop grpc server error", zap.Error(err))
			}
		}

		os.Exit(0)
	}
}

var ProviderSet = wire.NewSet(New)
