package usermodel

import (
	"database/sql"
	"time"
)

type (
	ZeroUsers struct {
		Id          int64        // id
		Account     string       // 账号
		Username    string       // 用户名
		Password    string       // 密码
		Gender      int64        // 性别 1：未设置；2：男性；3：女性
		UpdatedBy   int64        // 更新人
		UpdatedAt   time.Time    // 更新时间
		CreatedBy   int64        // 创建人
		CreatedAt   time.Time    // 创建时间
		DeletedAt   sql.NullTime // 删除时间
		DeletedFlag int64        // 是否删除 1：正常  2：已删除
	}

	ZeroUserInfos struct {
		Id          int64        // id
		UserId      int64        // 用户id
		Email       string       // 邮箱
		Phone       string       // 手机号
		UpdatedBy   int64        // 更新人
		UpdatedAt   time.Time    // 更新时间
		CreatedBy   int64        // 创建人
		CreatedAt   time.Time    // 创建时间
		DeletedAt   sql.NullTime // 删除时间
		DeletedFlag int64        // 是否删除 1：正常  2：已删除
	}
)
