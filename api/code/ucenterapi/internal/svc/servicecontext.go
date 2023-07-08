package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"go-zero-micro/api/code/ucenterapi/internal/config"
	"go-zero-micro/api/code/ucenterapi/internal/middleware"
	"go-zero-micro/rpc/code/ucenter/client/filestorage"
	"go-zero-micro/rpc/code/ucenter/client/ucentergorm"
	"go-zero-micro/rpc/code/ucenter/client/ucentersqlx"
)

type ServiceContext struct {
	Config         config.Config
	Check          rest.Middleware
	UcenterGormRpc ucentergorm.UcenterGorm //gorm方式的接口
	UcenterSqlxRpc ucentersqlx.UcenterSqlx //sqlx方式的接口
	FileStorageRpc filestorage.FileStorage //文件存储相关接口
}

func NewServiceContext(c config.Config) *ServiceContext {
	uCenterRpcClient := zrpc.MustNewClient(c.UCenterRpc)

	return &ServiceContext{
		Config:         c,
		Check:          middleware.NewCheckMiddleware().Handle,
		UcenterGormRpc: ucentergorm.NewUcenterGorm(uCenterRpcClient),
		UcenterSqlxRpc: ucentersqlx.NewUcenterSqlx(uCenterRpcClient),
		FileStorageRpc: filestorage.NewFileStorage(uCenterRpcClient),
	}
}
