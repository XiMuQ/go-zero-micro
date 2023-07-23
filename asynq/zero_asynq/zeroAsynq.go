package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"go-zero-micro/asynq/zero_asynq/internal/config"
	"go-zero-micro/asynq/zero_asynq/internal/register"
	"go-zero-micro/asynq/zero_asynq/internal/svc"
	"os"
	"time"

	"github.com/zeromicro/go-zero/core/conf"
)

var configFile = flag.String("f", "conf/dev/asynq/zero-asynq.yaml", "the config file")

func main() {
	flag.Parse()
	var c config.Config

	conf.MustLoad(*configFile, &c, conf.UseEnv())

	// log、prometheus、trace、metricsUrl
	if err := c.SetUp(); err != nil {
		panic(err)
	}

	svcContext := svc.NewServiceContext(c)
	ctx := context.Background()

	waitTime := time.Second * 2
	if c.Mode == service.DevMode {
		waitTime = time.Millisecond * 300
	}
	time.AfterFunc(waitTime, func() {
		AsynqClientSchedulerSetUp(ctx, svcContext)
	})

	AsynqServerSetUp(ctx, svcContext)
}

// AsynqServerSetUp asynq服务端启动
func AsynqServerSetUp(ctx context.Context, svcContext *svc.ServiceContext) {
	asynqServer := register.NewZeroAsynqServer(ctx, svcContext)
	mux := asynqServer.AsynqServerHandlerRegister()

	fmt.Printf("[asynqServer] starting server\n\n")
	if err := svcContext.AsynqServer.Run(mux); err != nil {
		logx.WithContext(ctx).Errorf("[AsynqServer] run err : %+v", err)
		os.Exit(1)
	}
}

// AsynqClientSchedulerSetUp asynq客户端启动
func AsynqClientSchedulerSetUp(ctx context.Context, svcContext *svc.ServiceContext) {
	asynqClientScheduler := register.NewZeroAsynqClientScheduler(ctx, svcContext)
	asynqClientScheduler.ZeroAsynqClientSchedulerRegister()

	fmt.Printf("[asynqClientScheduler] starting server\n\n")
	if err := svcContext.AsynqClientScheduler.Run(); err != nil {
		logx.WithContext(ctx).Errorf("[asynqClientScheduler] run err : %+v", err)
		os.Exit(1)
	}
}
