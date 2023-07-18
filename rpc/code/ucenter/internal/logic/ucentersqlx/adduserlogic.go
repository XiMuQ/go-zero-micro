package ucentersqlxlogic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go-zero-micro/common/errorx"
	"go-zero-micro/common/utils"
	sqlc_usermodel "go-zero-micro/rpc/database/sqlc/usermodel"
	"time"

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
	userId := utils.GetUidFromCtxInt64(l.ctx, "userId")
	currentTime := time.Now()
	/**
	  1、需求逻辑：User表保存账号信息，UserInfo表是子表，保存关联信息，比如：邮箱、手机号等
	  2、代码逻辑：先插入User表，后插入UserInfo表数据，插入UserInfo表时需要获取User表插入的id
	  3、无事务特性时：可能会出现主表有数据，但子表无数据的情况，导致数据不一致
	*/
	var InsertUserId int64

	//将对主子表的操作全部放到同一个事务中，每一步操作有错误就返回错误，没有错误最后就返回nil，事务遇到错误会回滚；
	if err := l.svcCtx.SqlcUsersModel.TransCtx(l.ctx, func(context context.Context, session sqlx.Session) error {
		userParam := &sqlc_usermodel.ZeroUsers{}
		copier.Copy(userParam, in)
		userParam.Password = utils.GeneratePassword(l.svcCtx.Config.DefaultConfig.DefaultPassword)
		userParam.CreatedBy = userId
		userParam.CreatedAt = currentTime
		userParam.DeletedFlag = 1
		dbUserRes, err := l.svcCtx.SqlcUsersModel.TransSaveCtx(l.ctx, session, userParam)
		if err != nil {
			return err
		}
		uid, err := dbUserRes.LastInsertId()
		if err != nil {
			return err
		}

		userInfoParam := &sqlc_usermodel.ZeroUserInfos{}
		copier.Copy(userInfoParam, in)
		userInfoParam.UserId = uid
		userInfoParam.CreatedBy = userId
		userInfoParam.CreatedAt = currentTime
		userInfoParam.DeletedFlag = 1
		_, err = l.svcCtx.SqlcUserInfosModel.TransSaveCtx(l.ctx, session, userInfoParam)
		if err != nil {
			return err
		}
		InsertUserId = uid
		return nil
	}); err != nil {
		return nil, errorx.NewDefaultError(errorx.DbAddErrorCode)
	}

	return &ucenter.BaseResp{
		Id: InsertUserId,
	}, nil
}
