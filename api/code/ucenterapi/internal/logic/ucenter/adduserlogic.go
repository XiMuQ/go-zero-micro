package ucenter

import (
	"context"
	"github.com/jinzhu/copier"
	"go-zero-micro/common/errorx"
	"go-zero-micro/rpc/code/ucenter/ucenter"

	"go-zero-micro/api/code/ucenterapi/internal/svc"
	"go-zero-micro/api/code/ucenterapi/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserLogic {
	return &AddUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddUserLogic) AddUser(req *types.UserSimpleModel) (resp *types.BaseModel, err error) {
	//这里没做严格的参数校验
	if req.Account == "" || req.Username == "" {
		return nil, errorx.NewDefaultError(errorx.ParamErrorCode)
	}
	param := &ucenter.User{}
	copier.Copy(param, req)
	rpcRes, err := l.svcCtx.UcenterSqlxRpc.AddUser(l.ctx, param)
	if err != nil {
		return nil, errorx.NewDefaultError(errorx.ServerErrorCode)
	}
	resp = &types.BaseModel{
		Id:   rpcRes.Id,
		Data: "用户添加成功",
	}
	return resp, nil
}
