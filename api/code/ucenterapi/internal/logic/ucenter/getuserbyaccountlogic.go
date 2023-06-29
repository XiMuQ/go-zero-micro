package ucenter

import (
	"context"

	"go-zero-micro/api/code/ucenterapi/internal/svc"
	"go-zero-micro/api/code/ucenterapi/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserByAccountLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserByAccountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByAccountLogic {
	return &GetUserByAccountLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserByAccountLogic) GetUserByAccount(req *types.UserSimpleModel) (resp *types.BaseModel, err error) {
	// todo: add your logic here and delete this line

	return
}
