package main

import (
	"context"
	"flag"
	"fmt"
	filestorageServer "go-zero-micro/rpc/code/ucenter/internal/server/filestorage"

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

var configFile = flag.String("f", "conf/dev/rpc/ucenter.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		ucenter.RegisterUcenterSqlxServer(grpcServer, ucentersqlxServer.NewUcenterSqlxServer(ctx))
		ucenter.RegisterUcenterGormServer(grpcServer, ucentergormServer.NewUcenterGormServer(ctx))

		//新增的分组接口必须要在这里注册，根据proto生成时可能未新增，否则会报 unknown service ucenter.fileStorage
		ucenter.RegisterFileStorageServer(grpcServer, filestorageServer.NewFileStorageServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})

	MaxFileSize := int(c.UploadFile.MaxFileSize)
	//调整RPC服务端收到的消息体大小限制
	s.AddOptions(grpc.MaxRecvMsgSize(MaxFileSize))

	//拦截器
	s.AddUnaryInterceptors(func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		fmt.Printf("RpcServerInterceptor1 ====> Start \n")
		fmt.Printf("req =====================> %+v \n", req)
		fmt.Printf("info =====================> %+v \n", info)
		resp, err = handler(ctx, req)
		fmt.Printf("resp =====================> %+v \n", resp)
		fmt.Printf("RpcServerInterceptor1 ====> End \n")
		return resp, err
	})

	//拦截器
	s.AddUnaryInterceptors(func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		fmt.Printf("RpcServerInterceptor2 ====> Start \n")
		fmt.Printf("req =====================> %+v \n", req)
		fmt.Printf("info =====================> %+v \n", info)
		resp, err = handler(ctx, req)
		fmt.Printf("resp =====================> %+v \n", resp)
		fmt.Printf("RpcServerInterceptor2 ====> End \n")
		return resp, err
	})

	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
