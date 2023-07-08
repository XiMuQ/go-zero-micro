package ucentersqlxlogic

import (
	"context"
	"go-zero-micro/common/utils"
	"time"

	"go-zero-micro/rpc/code/ucenter/internal/svc"
	"go-zero-micro/rpc/code/ucenter/ucenter"

	"github.com/jinzhu/copier"
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
	//return &ucenter.UserLoginResp{}, nil

	//模拟耗时 20秒钟
	sleepTime := 20 * time.Second
	time.Sleep(sleepTime)
	return l.LoginSuccess(in)
}

func (l *LoginUserLogic) LoginSuccess(in *ucenter.User) (*ucenter.UserLoginResp, error) {
	AccessSecret := l.svcCtx.Config.JWT.AccessSecret
	AccessExpire := l.svcCtx.Config.JWT.AccessExpire
	now := time.Now().Unix()

	jwtToken, err := utils.GenerateJwtToken(AccessSecret, now, AccessExpire, in.Id)
	if err != nil {
		return nil, err
	}
	resp := &ucenter.UserLoginResp{}
	copier.Copy(resp, in)
	resp.AccessToken = jwtToken
	resp.AccessExpire = now + AccessExpire
	resp.RefreshAfter = now + AccessExpire/2
	return resp, nil
}
