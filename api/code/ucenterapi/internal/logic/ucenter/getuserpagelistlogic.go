package ucenter

import (
	"context"

	"go-zero-micro/api/code/ucenterapi/internal/svc"
	"go-zero-micro/api/code/ucenterapi/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserPageListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserPageListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserPageListLogic {
	return &GetUserPageListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserPageListLogic) GetUserPageList(req *types.UserListReq) (resp *types.UserPageResp, err error) {
	// todo: add your logic here and delete this line

	return
}
