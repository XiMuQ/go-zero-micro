package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hibiken/asynq"
	"go-zero-micro/asynq/zero_asynq/internal/svc"
	"go-zero-micro/common/consts"
	"go-zero-micro/common/utils"
	"time"
)

// ZeroAsynqServerHandler 服务端
type ZeroAsynqServerHandler struct {
	svcCtx *svc.ServiceContext
}

func NewZeroAsynqServerHandler(svcCtx *svc.ServiceContext) *ZeroAsynqServerHandler {
	return &ZeroAsynqServerHandler{
		svcCtx: svcCtx,
	}
}

// ProcessTask 服务端处理任务
func (l *ZeroAsynqServerHandler) ProcessTask(ctx context.Context, task *asynq.Task) error {
	var payload consts.ZeroAsynqPayload
	if err := json.Unmarshal(task.Payload(), &payload); err != nil {
		//return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
		return err
	}
	// TODO: 模拟处理数据
	fmt.Printf("\nZeroAsync Server：ProcessTask：%s Start handle AsyncTask!\n", time.Now().Format(utils.DateTimeFormat))
	fmt.Printf("ZeroAsync Server Received Message: %v\n", payload)
	fmt.Println("ZeroAsync Server：ProcessTask：End handle AsyncTask!")
	return nil
}
