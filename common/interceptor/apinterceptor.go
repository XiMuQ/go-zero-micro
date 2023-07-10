package interceptor

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
)

// RpcClientInterceptor1 rpc的客户端拦截器
func RpcClientInterceptor1(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	fmt.Printf("RpcClientInterceptor1 ====> Start \n")
	fmt.Printf("req =====================> %+v \n", req)

	err := invoker(ctx, method, req, reply, cc, opts...)
	fmt.Printf("RpcClientInterceptor1 ====> End \n")
	if err != nil {
		return err
	}
	return nil
}

// RpcClientInterceptor2 rpc的客户端拦截器
func RpcClientInterceptor2(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	fmt.Printf("RpcClientInterceptor2 ====> Start \n")
	fmt.Printf("req =====================> %+v \n", req)

	err := invoker(ctx, method, req, reply, cc, opts...)
	fmt.Printf("RpcClientInterceptor2 ====> End \n")
	if err != nil {
		return err
	}
	return nil
}
