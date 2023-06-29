package ucenter

import (
	"context"

	"go-zero-micro/api/code/ucenterapi/internal/svc"
	"go-zero-micro/api/code/ucenterapi/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryUserByPathLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryUserByPathLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryUserByPathLogic {
	return &QueryUserByPathLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryUserByPathLogic) QueryUserByPath(req *types.PathReq) (resp *types.BaseModel, err error) {
	// todo: add your logic here and delete this line

	return
}
