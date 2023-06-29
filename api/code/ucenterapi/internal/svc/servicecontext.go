package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"go-zero-micro/api/code/ucenterapi/internal/config"
	"go-zero-micro/api/code/ucenterapi/internal/middleware"
)

type ServiceContext struct {
	Config config.Config
	Check  rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Check:  middleware.NewCheckMiddleware().Handle,
	}
}
