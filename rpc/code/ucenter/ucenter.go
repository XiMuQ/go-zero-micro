package main

import (
	"flag"
	"fmt"

	"go-zero-micro/rpc/code/ucenter/internal/config"
	ucentergormServer "go-zero-micro/rpc/code/ucenter/internal/server/ucentergorm"
	ucentersqlxServer "go-zero-micro/rpc/code/ucenter/internal/server/ucentersqlx"
	"go-zero-micro/rpc/code/ucenter/internal/svc"
	"go-zero-micro/rpc/code/ucenter/ucenter"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/ucenter.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		ucenter.RegisterUcenterSqlxServer(grpcServer, ucentersqlxServer.NewUcenterSqlxServer(ctx))
		ucenter.RegisterUcenterGormServer(grpcServer, ucentergormServer.NewUcenterGormServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}