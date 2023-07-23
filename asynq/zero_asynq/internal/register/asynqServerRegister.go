package register

import (
	"context"
	"github.com/hibiken/asynq"
	"go-zero-micro/asynq/zero_asynq/internal/handler"
	"go-zero-micro/asynq/zero_asynq/internal/svc"
	"go-zero-micro/common/consts"
)

type ZeroAsynqServer struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewZeroAsynqServer(ctx context.Context, svcCtx *svc.ServiceContext) *ZeroAsynqServer {
	return &ZeroAsynqServer{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// AsynqServerHandlerRegister Register Task Handler
func (l *ZeroAsynqServer) AsynqServerHandlerRegister() *asynq.ServeMux {
	mux := asynq.NewServeMux()
	//scheduler job
	mux.Handle(consts.ZeroAsynqDemo, handler.NewZeroAsynqServerHandler(l.svcCtx))
	//queue job , asynq support queue job
	// wait you fill..

	return mux
}
