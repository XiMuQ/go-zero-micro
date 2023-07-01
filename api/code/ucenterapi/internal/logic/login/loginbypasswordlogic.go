package login

import (
	"context"
	"github.com/jinzhu/copier"
	"go-zero-micro/common/errorx"
	"go-zero-micro/rpc/code/ucenter/ucenter"

	"go-zero-micro/api/code/ucenterapi/internal/svc"
	"go-zero-micro/api/code/ucenterapi/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginByPasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginByPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginByPasswordLogic {
	return &LoginByPasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginByPasswordLogic) LoginByPassword(req *types.UserLoginPasswordModel) (resp *types.UserLoginResp, err error) {
	// todo: add your logic here and delete this line
	param := &ucenter.User{}
	copier.Copy(param, req)
	loginRes, err := l.svcCtx.UcenterSqlxRpc.LoginUser(l.ctx, param)
	if err != nil {
		return nil, errorx.NewDefaultError(errorx.UserLoginPasswordErrorCode)
	}
	res := &types.UserLoginResp{}
	copier.Copy(res, loginRes)
	return res, nil
}
