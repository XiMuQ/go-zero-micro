package svc

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go-zero-micro/rpc/code/ucenter/internal/config"
	sqlc_usermodel "go-zero-micro/rpc/database/sqlc/usermodel"
	sqlx_usermodel "go-zero-micro/rpc/database/sqlx/usermodel"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type ServiceContext struct {
	Config      config.Config
	RedisClient *redis.Redis

	SqlxUsersModel     sqlx_usermodel.ZeroUsersModel
	SqlxUserInfosModel sqlx_usermodel.ZeroUserInfosModel

	SqlcUsersModel     sqlc_usermodel.ZeroUsersModel
	SqlcUserInfosModel sqlc_usermodel.ZeroUserInfosModel
	GormDb             *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysqlConn := sqlx.NewMysql(c.MySQL.DataSource)

	gormDb, err := gorm.Open(mysql.Open(c.MySQL.DataSource), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			//TablePrefix:   "tech_", // 表名前缀，`User` 的表名应该是 `t_users`
			//SingularTable: true, // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `t_user`
		},
	})
	if err != nil {
		errInfo := fmt.Sprintf("Gorm connect database err:%v", err)
		panic(errInfo)
	}
	//自动同步更新表结构,不要建表了O(∩_∩)O哈哈~
	//db.AutoMigrate(&models.User{})

	redisConn := redis.New(c.Redis.Host, func(r *redis.Redis) {
		r.Type = c.Redis.Type
		r.Pass = c.Redis.Pass
	})
	return &ServiceContext{
		Config:             c,
		RedisClient:        redisConn,
		SqlxUsersModel:     sqlx_usermodel.NewZeroUsersModel(mysqlConn),
		SqlxUserInfosModel: sqlx_usermodel.NewZeroUserInfosModel(mysqlConn),

		SqlcUsersModel:     sqlc_usermodel.NewZeroUsersModel(mysqlConn, c.CacheRedis),
		SqlcUserInfosModel: sqlc_usermodel.NewZeroUserInfosModel(mysqlConn, c.CacheRedis),
		GormDb:             gormDb,
	}
}
