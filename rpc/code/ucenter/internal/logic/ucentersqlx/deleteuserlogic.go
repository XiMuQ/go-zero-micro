package ucentersqlxlogic

import (
	"context"

	"go-zero-micro/rpc/code/ucenter/internal/svc"
	"go-zero-micro/rpc/code/ucenter/ucenter"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserLogic {
	return &DeleteUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 注销用户信息
func (l *DeleteUserLogic) DeleteUser(in *ucenter.BaseModel) (*ucenter.BaseResp, error) {
	// todo: add your logic here and delete this line

	return &ucenter.BaseResp{}, nil
}
