package interceptor

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
)

// RpcServerInterceptor1 rpc的服务端拦截器
func RpcServerInterceptor1(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	fmt.Printf("RpcServerInterceptor1 ====> Start \n")
	fmt.Printf("req =====================> %+v \n", req)
	fmt.Printf("info =====================> %+v \n", info)
	resp, err = handler(ctx, req)
	fmt.Printf("resp =====================> %+v \n", resp)
	fmt.Printf("RpcServerInterceptor1 ====> End \n")
	return resp, err
}

// RpcServerInterceptor2 rpc的服务端拦截器
func RpcServerInterceptor2(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	fmt.Printf("RpcServerInterceptor2 ====> Start \n")
	fmt.Printf("req =====================> %+v \n", req)
	fmt.Printf("info =====================> %+v \n", info)
	resp, err = handler(ctx, req)
	fmt.Printf("resp =====================> %+v \n", resp)
	fmt.Printf("RpcServerInterceptor2 ====> End \n")
	return resp, err
}
