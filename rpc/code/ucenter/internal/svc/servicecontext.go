package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go-zero-micro/rpc/code/ucenter/internal/config"
	"go-zero-micro/rpc/database/sqlx/usermodel"
)

type ServiceContext struct {
	Config     config.Config
	UsersModel usermodel.ZeroUsersModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysqlConn := sqlx.NewMysql(c.MySQL.DataSource)

	return &ServiceContext{
		Config:     c,
		UsersModel: usermodel.NewZeroUsersModel(mysqlConn),
	}
}
