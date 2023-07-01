package ucentergormlogic

import (
	"context"

	"go-zero-micro/rpc/code/ucenter/internal/svc"
	"go-zero-micro/rpc/code/ucenter/ucenter"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginUserLogic {
	return &LoginUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// LoginUser 用户登录
func (l *LoginUserLogic) LoginUser(in *ucenter.User) (*ucenter.UserLoginResp, error) {
	// todo: add your logic here and delete this line

	return &ucenter.UserLoginResp{}, nil
}
