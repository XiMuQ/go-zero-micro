package register

import (
	"context"
	"encoding/json"
	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-micro/asynq/zero_asynq/internal/svc"
	"go-zero-micro/common/consts"
	"time"
)

type ZeroAsynqClientScheduler struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewZeroAsynqClientScheduler(ctx context.Context, svcCtx *svc.ServiceContext) *ZeroAsynqClientScheduler {
	return &ZeroAsynqClientScheduler{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// ZeroAsynqClientSchedulerRegister Register Task Handler
func (l *ZeroAsynqClientScheduler) ZeroAsynqClientSchedulerRegister() {
	param := consts.ZeroAsynqPayload{
		Id:       "123qwe",
		UserName: "测试",
		Password: "123456",
	}
	payload, err := json.Marshal(param)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("json.Unmarshal failed: %v", err)
	}

	task := asynq.NewTask(consts.ZeroAsynqDemo, payload)
	// every one minute exec
	entryID, err := l.svcCtx.AsynqClientScheduler.Register("* * * * *", task, asynq.MaxRetry(5), asynq.Timeout(1*time.Minute))
	if err != nil {
		logx.WithContext(l.ctx).Errorf("!!!AsynqClientScheduler!!! ====> 【ZeroAsynqClientSchedulerRegister】 registered  err:%+v , task:%+v", err, task)
	}
	logx.WithContext(l.ctx).Infof("【ZeroAsynqClientSchedulerRegister】 registered an  entry: %q \n", entryID)
}
