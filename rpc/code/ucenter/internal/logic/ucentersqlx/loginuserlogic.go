package ucentersqlxlogic

import (
	"context"
	"errors"
	"fmt"
	"go-zero-micro/common/utils"
	"go-zero-micro/rpc/database/sqlx/usermodel"
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
	param := &usermodel.ZeroUsers{
		Account: in.Account,
	}
	dbRes, err := l.svcCtx.UsersModel.FindOneByParamCtx(l.ctx, param)
	if err != nil {
		logx.Error(err)
		errInfo := fmt.Sprintf("LoginUser:FindOneByParam:db err:%v , in : %+v", err, in)
		return nil, errors.New(errInfo)
	}
	if utils.ComparePassword(in.Password, dbRes.Password) {
		copier.Copy(in, dbRes)
		return l.LoginSuccess(in)
	} else {
		errInfo := fmt.Sprintf("LoginUser:user password error:in : %+v", in)
		return nil, errors.New(errInfo)
	}
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
