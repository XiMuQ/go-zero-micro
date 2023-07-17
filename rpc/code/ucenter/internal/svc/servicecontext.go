package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go-zero-micro/rpc/code/ucenter/internal/config"
	sqlx_usermodel "go-zero-micro/rpc/database/sqlx/usermodel"
)

type ServiceContext struct {
	Config             config.Config
	SqlxUsersModel     sqlx_usermodel.ZeroUsersModel
	SqlxUserInfosModel sqlx_usermodel.ZeroUserInfosModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysqlConn := sqlx.NewMysql(c.MySQL.DataSource)

	return &ServiceContext{
		Config:             c,
		SqlxUsersModel:     sqlx_usermodel.NewZeroUsersModel(mysqlConn),
		SqlxUserInfosModel: sqlx_usermodel.NewZeroUserInfosModel(mysqlConn),
	}
}
