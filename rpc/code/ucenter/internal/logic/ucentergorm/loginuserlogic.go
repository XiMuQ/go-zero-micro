package ucentergormlogic

import (
	"context"
	"errors"
	"fmt"
	"github.com/jinzhu/copier"
	"go-zero-micro/common/utils"
	gorm_usermodel "go-zero-micro/rpc/database/gorm/usermodel"
	"time"

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
	param := &gorm_usermodel.ZeroUsers{
		Id:      in.Id,
		Account: in.Account,
	}
	dbRes := &gorm_usermodel.ZeroUsers{}
	l.svcCtx.GormDb.Where(param).First(dbRes)

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
