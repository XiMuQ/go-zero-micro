package svc

import (
	"fmt"
	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/zrpc"
	"go-zero-micro/asynq/zero_asynq/internal/config"
	"go-zero-micro/rpc/code/ucenter/client/ucentergorm"
	"time"
)

type ServiceContext struct {
	Config               config.Config
	AsynqServer          *asynq.Server           //asynq服务端
	AsynqClientScheduler *asynq.Scheduler        //asynq定时器类型的客户端
	UcenterGormRpc       ucentergorm.UcenterGorm //gorm方式的接口
}

func NewServiceContext(c config.Config) *ServiceContext {
	uCenterRpcClient := zrpc.MustNewClient(c.UCenterRpc)

	return &ServiceContext{
		Config:               c,
		AsynqServer:          NewAsynqServer(c),
		AsynqClientScheduler: NewAsynqClientScheduler(c),
		UcenterGormRpc:       ucentergorm.NewUcenterGorm(uCenterRpcClient),
	}
}

// NewAsynqServer 初始化asynq服务端
func NewAsynqServer(c config.Config) *asynq.Server {
	return asynq.NewServer(
		asynq.RedisClientOpt{Addr: c.Redis.Host, Password: c.Redis.Pass},
		asynq.Config{
			IsFailure: func(err error) bool {
				fmt.Printf("asynq server exec task IsFailure ======== >>>>>>>>>>>  err : %+v \n", err)
				return true
			},
			Concurrency: 20, //max concurrent process job task num
		},
	)
}

// NewAsynqClientScheduler 初始化asynq客户端
func NewAsynqClientScheduler(c config.Config) *asynq.Scheduler {
	location, _ := time.LoadLocation("Asia/Shanghai")
	return asynq.NewScheduler(
		asynq.RedisClientOpt{
			Addr:     c.Redis.Host,
			Password: c.Redis.Pass,
		}, &asynq.SchedulerOpts{
			Location: location,
			EnqueueErrorHandler: func(task *asynq.Task, opts []asynq.Option, err error) {
				fmt.Printf("Scheduler EnqueueErrorHandler <<<<<<<===>>>>> err : %+v , task : %+v", err, task)
			},
		})
}
