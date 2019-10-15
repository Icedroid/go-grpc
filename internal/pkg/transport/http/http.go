package http

import (
	"fmt"
	"net/http"

	"github.com/google/wire"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	"github.com/Icedroid/go-grpc/internal/pkg/utils/netutil"
)

type Options struct {
	Port int
}

type Server struct {
	o          *Options
	app        string
	host       string
	port       int
	logger     *zap.Logger
	router     *http.ServeMux
	httpServer http.Server
}

func NewOptions(v *viper.Viper) (*Options, error) {
	var (
		err error
		o   = new(Options)
	)

	if err = v.UnmarshalKey("http", o); err != nil {
		return nil, err
	}

	return o, err
}

type InitRouter func(r *runtime.ServeMux)

func New(o *Options, logger *zap.Logger, init InitRouter) (*Server, error) {
	mux := http.NewServeMux()

	gwmux := runtime.NewServeMux(runtime.WithMarshalerOption(runtime.MIMEWildcard,
		&runtime.JSONPb{EnumsAsInts: true, OrigName: true, EmitDefaults: true}))

	init(gwmux)

	mux.Handle("/", gwmux)

	var s = &Server{
		logger: logger.With(zap.String("type", "http.Server")),
		router: mux,
		o:      o,
	}

	return s, nil
}

func (s *Server) Application(name string) {
	s.app = name
}

func (s *Server) Start() error {
	s.port = s.o.Port
	if s.port == 0 {
		s.port = netutil.GetAvailablePort()
	}

	s.host = netutil.GetLocalIP4()
	if s.host == "" {
		return errors.New("get local ipv4 error")
	}

	addr := fmt.Sprintf("%s:%d", s.host, s.port)

	// s.httpServer = http.Server{Addr: addr, Handler: s.router}

	s.logger.Info("http server starting ...", zap.String("addr", addr))
	go func() {
		if err := http.ListenAndServe(addr, s.router); err != nil && err != http.ErrServerClosed {
			s.logger.Fatal("start http server err", zap.Error(err))
			return
		}
	}()

	return nil
}

func (s *Server) Stop() error {
	s.logger.Info("htp server stopping ...")

	return s.httpServer.Close()
}

var ProviderSet = wire.NewSet(New, NewOptions)
