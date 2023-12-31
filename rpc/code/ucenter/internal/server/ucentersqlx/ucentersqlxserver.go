// Code generated by goctl. DO NOT EDIT.
// Source: ucenter.proto

package server

import (
	"context"

	"go-zero-micro/rpc/code/ucenter/internal/logic/ucentersqlx"
	"go-zero-micro/rpc/code/ucenter/internal/svc"
	"go-zero-micro/rpc/code/ucenter/ucenter"
)

type UcenterSqlxServer struct {
	svcCtx *svc.ServiceContext
	ucenter.UnimplementedUcenterSqlxServer
}

func NewUcenterSqlxServer(svcCtx *svc.ServiceContext) *UcenterSqlxServer {
	return &UcenterSqlxServer{
		svcCtx: svcCtx,
	}
}

// 获取用户信息
func (s *UcenterSqlxServer) GetUser(ctx context.Context, in *ucenter.BaseModel) (*ucenter.User, error) {
	l := ucentersqlxlogic.NewGetUserLogic(ctx, s.svcCtx)
	return l.GetUser(in)
}

// 添加用户
func (s *UcenterSqlxServer) AddUser(ctx context.Context, in *ucenter.User) (*ucenter.BaseResp, error) {
	l := ucentersqlxlogic.NewAddUserLogic(ctx, s.svcCtx)
	return l.AddUser(in)
}

// 注销用户信息
func (s *UcenterSqlxServer) DeleteUser(ctx context.Context, in *ucenter.BaseModel) (*ucenter.BaseResp, error) {
	l := ucentersqlxlogic.NewDeleteUserLogic(ctx, s.svcCtx)
	return l.DeleteUser(in)
}

// 用户登录
func (s *UcenterSqlxServer) LoginUser(ctx context.Context, in *ucenter.User) (*ucenter.UserLoginResp, error) {
	l := ucentersqlxlogic.NewLoginUserLogic(ctx, s.svcCtx)
	return l.LoginUser(in)
}
