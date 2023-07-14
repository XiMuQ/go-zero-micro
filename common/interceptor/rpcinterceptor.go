package interceptor

import (
	"context"
	"fmt"
	"go-zero-micro/common/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"strconv"
)

// RpcServerInterceptor1 rpc的服务端拦截器
func RpcServerInterceptor1(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	//fmt.Printf("RpcServerInterceptor1 ====> Start \n")
	//fmt.Printf("req =====================> %+v \n", req)
	//fmt.Printf("info =====================> %+v \n", info)
	resp, err = handler(ctx, req)
	//fmt.Printf("resp =====================> %+v \n", resp)
	//fmt.Printf("RpcServerInterceptor1 ====> End \n")
	return resp, err
}

// RpcServerInterceptor2 rpc的服务端拦截器
func RpcServerInterceptor2(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	//fmt.Printf("RpcServerInterceptor2 ====> Start \n")
	//fmt.Printf("req =====================> %+v \n", req)
	//fmt.Printf("info =====================> %+v \n", info)

	if md, ok := metadata.FromIncomingContext(ctx); ok {
		tmp := md.Get("userId")
		if len(tmp) > 0 {
			userId, _ := utils.Base64Decode(tmp[0])
			uid, _ := strconv.ParseInt(userId, 10, 64)
			fmt.Printf("userId：%d\n", uid)
		}
		//uname := md.Get("userName")
		//if len(tmp) > 0 {
		//	userName, _ := utils.Base64Decode(uname[0])
		//	fmt.Printf("userName：%s\n", userName)
		//}
	}

	resp, err = handler(ctx, req)
	//fmt.Printf("resp =====================> %+v \n", resp)
	//fmt.Printf("RpcServerInterceptor2 ====> End \n")
	return resp, err
}
