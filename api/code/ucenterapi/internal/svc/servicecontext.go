package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"go-zero-micro/api/code/ucenterapi/internal/config"
	"go-zero-micro/api/code/ucenterapi/internal/middleware"
	"go-zero-micro/rpc/code/ucenter/ucenter"
)

type ServiceContext struct {
	Config config.Config
	Check  rest.Middleware
	//UcenterGormRpc ucentergorm.UcenterGorm //gorm方式的接口
	//UcenterSqlxRpc ucentersqlx.UcenterSqlx //sqlx方式的接口
	UcenterSqlxDC ucenter.UcenterSqlxClient //sqlx方式的GRPC接口
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := zrpc.MustNewClient(zrpc.RpcClientConf{
		Target: "dns:///127.0.0.1:8080",
	})
	return &ServiceContext{
		Config: c,
		Check:  middleware.NewCheckMiddleware().Handle,
		//UcenterGormRpc: ucentergorm.NewUcenterGorm(zrpc.MustNewClient(c.UCenterRpc)),
		//UcenterSqlxRpc: ucentersqlx.NewUcenterSqlx(zrpc.MustNewClient(c.UCenterRpc)),
		UcenterSqlxDC: ucenter.NewUcenterSqlxClient(conn.Conn()),
	}
}
