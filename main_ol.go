package main

import (
	"context"
	"errors"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"git.fogcdn.top/axe/ops-playbook/config"
	"git.fogcdn.top/axe/ops-playbook/dao"
	"git.fogcdn.top/axe/ops-playbook/dao/mysql"
	"git.fogcdn.top/axe/ops-playbook/handler"
	"git.fogcdn.top/axe/ops-playbook/utils/logger"

	pb "git.fogcdn.top/axe/protos/goout/playbook"
	pbtemplate "git.fogcdn.top/axe/protos/goout/template"
)

const (
	port     = ":8081"
	restPort = ":8082"
)

func main() {
	// 手动引入初始化操作，因为单元测试找不到配置文件
	config.Initial()
	// 引入logger初始化
	logger.Initial()

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()

	var dsn = config.Config().GetString("db.mysql")
	if dsn == "" {
		err := errors.New("mysql connect info is empty")
		panic(err)
	}
	db := dao.NewMySQL(dsn)
	playbookRepo := mysql.NewPlaybookRepo(db)
	playbookFileRepo := mysql.NewPlaybookFileRepo(db)
	playbookEntrypointRepo := mysql.NewPlaybookEntrypointRepo(db)
	pb.RegisterPlaybookServer(server, handler.NewPlaybook(playbookRepo, playbookFileRepo, playbookEntrypointRepo))

	templateRepo := mysql.NewTemplateRepo(db)
	pbtemplate.RegisterTemplateServer(server, handler.NewTemplate(templateRepo, playbookRepo, playbookFileRepo, playbookEntrypointRepo))

	// Register reflection service on gRPC server.
	reflection.Register(server)

	// restful api
	go runRest()

	log.Println("grpc listen on ", port)
	log.Println("restAPI listen on ", restPort)

	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

func runRest() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := http.NewServeMux()

	gwmux := runtime.NewServeMux(runtime.WithMarshalerOption(runtime.MIMEWildcard,
		&runtime.JSONPb{EnumsAsInts: true, OrigName: true, EmitDefaults: true}))

	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := pb.RegisterPlaybookHandlerFromEndpoint(ctx, gwmux, port, opts)
	if err != nil {
		goto Err
	}
	err = pbtemplate.RegisterTemplateHandlerFromEndpoint(ctx, gwmux, port, opts)
	if err != nil {
		goto Err
	}

	mux.Handle("/", gwmux)
	err = http.ListenAndServe(restPort, mux)

Err:
	if err != nil {
		panic(err)
	}
	return err
}
