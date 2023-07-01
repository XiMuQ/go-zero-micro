package ucentersqlxlogic

import (
	"context"

	"go-zero-micro/rpc/code/ucenter/internal/svc"
	"go-zero-micro/rpc/code/ucenter/ucenter"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserLogic {
	return &AddUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddUser 添加用户
func (l *AddUserLogic) AddUser(in *ucenter.User) (*ucenter.BaseResp, error) {
	// todo: add your logic here and delete this line

	return &ucenter.BaseResp{}, nil
}
