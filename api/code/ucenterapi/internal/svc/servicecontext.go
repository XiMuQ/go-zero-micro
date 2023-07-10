package svc

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"go-zero-micro/api/code/ucenterapi/internal/config"
	"go-zero-micro/api/code/ucenterapi/internal/middleware"
	"go-zero-micro/rpc/code/ucenter/client/filestorage"
	"go-zero-micro/rpc/code/ucenter/client/ucentergorm"
	"go-zero-micro/rpc/code/ucenter/client/ucentersqlx"
	"google.golang.org/grpc"
)

type ServiceContext struct {
	Config         config.Config
	Check          rest.Middleware
	UcenterGormRpc ucentergorm.UcenterGorm //gorm方式的接口
	UcenterSqlxRpc ucentersqlx.UcenterSqlx //sqlx方式的接口
	FileStorageRpc filestorage.FileStorage //文件存储相关接口
}

func NewServiceContext(c config.Config) *ServiceContext {
	MaxFileSize := int(c.UploadFile.MaxFileSize)
	//调整RPC客户端收到的消息体大小限制
	dialOption := grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(MaxFileSize))
	opt := zrpc.WithDialOption(dialOption)

	//声明拦截器
	interceptor1 := zrpc.WithUnaryClientInterceptor(func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		fmt.Printf("interceptor1 ====> Start \n")
		fmt.Printf("req =====================> %+v \n", req)

		err := invoker(ctx, method, req, reply, cc, opts...)
		fmt.Printf("interceptor1 ====> End \n")
		if err != nil {
			return err
		}
		return nil
	})

	//声明拦截器
	interceptor2 := zrpc.WithUnaryClientInterceptor(func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		fmt.Printf("interceptor2 ====> Start \n")
		fmt.Printf("req =====================> %+v \n", req)

		err := invoker(ctx, method, req, reply, cc, opts...)
		fmt.Printf("interceptor2 ====> End \n")
		if err != nil {
			return err
		}
		return nil
	})

	uCenterRpcClient := zrpc.MustNewClient(c.UCenterRpc, opt, interceptor1, interceptor2)

	return &ServiceContext{
		Config:         c,
		Check:          middleware.NewCheckMiddleware().Handle,
		UcenterGormRpc: ucentergorm.NewUcenterGorm(uCenterRpcClient),
		UcenterSqlxRpc: ucentersqlx.NewUcenterSqlx(uCenterRpcClient),
		FileStorageRpc: filestorage.NewFileStorage(uCenterRpcClient),
	}
}
