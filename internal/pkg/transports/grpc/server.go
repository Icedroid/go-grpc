package grpc

import (
	"fmt"
	"log"
	"net"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	netutil2 "github.com/Icedroid/go-grpc/internal/pkg/utils/netutil"
)

const (
	defaultPort = 8081
)

type ServerOptions struct {
	Port int
}

func NewServerOptions(v *viper.Viper) (*ServerOptions, error) {
	var (
		err error
		o   = new(ServerOptions)
	)
	if err = v.UnmarshalKey("grpc", o); err != nil {
		return nil, err
	}

	return o, nil
}

type Server struct {
	o      *ServerOptions
	app    string
	host   string
	port   int
	logger *zap.Logger
	server *grpc.Server
}

type InitServers func(s *grpc.Server)

func NewServer(o *ServerOptions, logger *zap.Logger, init InitServers, tracer opentracing.Tracer) (*Server, error) {
	// initialize grpc server
	var gs *grpc.Server
	logger = logger.With(zap.String("type", "grpc"))
	{
		grpc_prometheus.EnableHandlingTimeHistogram()
		gs = grpc.NewServer(
			grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
				grpc_ctxtags.StreamServerInterceptor(),
				grpc_prometheus.StreamServerInterceptor,
				grpc_zap.StreamServerInterceptor(logger),
				grpc_recovery.StreamServerInterceptor(),
				otgrpc.OpenTracingStreamServerInterceptor(tracer),
			)),
			grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
				grpc_ctxtags.UnaryServerInterceptor(),
				grpc_prometheus.UnaryServerInterceptor,
				grpc_zap.UnaryServerInterceptor(logger),
				grpc_recovery.UnaryServerInterceptor(),
				otgrpc.OpenTracingServerInterceptor(tracer),
			)),
		)
		init(gs)
	}

	// Register reflection service on gRPC server.
	reflection.Register(gs)

	return &Server{
		o:      o,
		logger: logger.With(zap.String("type", "grpc.Server")),
		server: gs,
	}, nil
}

func (s *Server) Application(name string) {
	s.app = name
}

func (s *Server) Start() error {
	s.port = s.o.Port
	if s.port == 0 {
		// s.port = netutil.GetAvailablePort()
		s.port = defaultPort
	}

	s.host = netutil2.GetLocalIP4()

	if s.host == "" {
		return errors.New("get local ipv4 error")
	}

	addr := fmt.Sprintf("%s:%d", s.host, s.port)

	s.logger.Info("grpc server starting ...", zap.String("addr", addr))
	go func() {
		lis, err := net.Listen("tcp", addr)
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}

		if err := s.server.Serve(lis); err != nil {
			s.logger.Fatal("failed to serve: %v", zap.Error(err))
		}
	}()

	return nil
}

func (s *Server) Stop() error {
	s.logger.Info("grpc server stopping ...")

	s.server.GracefulStop()
	return nil
}
