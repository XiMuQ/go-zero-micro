package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go-zero-micro/rpc/code/ucenter/internal/config"
	sqlc_usermodel "go-zero-micro/rpc/database/sqlc/usermodel"
	sqlx_usermodel "go-zero-micro/rpc/database/sqlx/usermodel"
)

type ServiceContext struct {
	Config             config.Config
	SqlxUsersModel     sqlx_usermodel.ZeroUsersModel
	SqlxUserInfosModel sqlx_usermodel.ZeroUserInfosModel

	SqlcUsersModel     sqlc_usermodel.ZeroUsersModel
	SqlcUserInfosModel sqlc_usermodel.ZeroUserInfosModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysqlConn := sqlx.NewMysql(c.MySQL.DataSource)

	return &ServiceContext{
		Config:             c,
		SqlxUsersModel:     sqlx_usermodel.NewZeroUsersModel(mysqlConn),
		SqlxUserInfosModel: sqlx_usermodel.NewZeroUserInfosModel(mysqlConn),

		SqlcUsersModel:     sqlc_usermodel.NewZeroUsersModel(mysqlConn, c.CacheRedis),
		SqlcUserInfosModel: sqlc_usermodel.NewZeroUserInfosModel(mysqlConn, c.CacheRedis),
	}
}
