// Code generated by goctl. DO NOT EDIT.

package usermodel

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	zeroUserInfosFieldNames          = builder.RawFieldNames(&ZeroUserInfos{})
	zeroUserInfosRows                = strings.Join(zeroUserInfosFieldNames, ",")
	zeroUserInfosRowsExpectAutoSet   = strings.Join(stringx.Remove(zeroUserInfosFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	zeroUserInfosRowsWithPlaceHolder = strings.Join(stringx.Remove(zeroUserInfosFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheZeroUserInfosIdPrefix = "cache:zeroUserInfos:id:"
)

type (
	zeroUserInfosModel interface {
		Insert(ctx context.Context, data *ZeroUserInfos) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*ZeroUserInfos, error)
		Update(ctx context.Context, data *ZeroUserInfos) error
		Delete(ctx context.Context, id int64) error
	}

	defaultZeroUserInfosModel struct {
		sqlc.CachedConn
		table string
	}

	ZeroUserInfos struct {
		Id          int64        `db:"id"`           // id
		UserId      int64        `db:"user_id"`      // 用户id
		Email       string       `db:"email"`        // 邮箱
		Phone       string       `db:"phone"`        // 手机号
		UpdatedBy   int64        `db:"updated_by"`   // 更新人
		UpdatedAt   time.Time    `db:"updated_at"`   // 更新时间
		CreatedBy   int64        `db:"created_by"`   // 创建人
		CreatedAt   time.Time    `db:"created_at"`   // 创建时间
		DeletedAt   sql.NullTime `db:"deleted_at"`   // 删除时间
		DeletedFlag int64        `db:"deleted_flag"` // 是否删除 1：正常  2：已删除
	}
)

func newZeroUserInfosModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultZeroUserInfosModel {
	return &defaultZeroUserInfosModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`zero_user_infos`",
	}
}

func (m *defaultZeroUserInfosModel) withSession(session sqlx.Session) *defaultZeroUserInfosModel {
	return &defaultZeroUserInfosModel{
		CachedConn: m.CachedConn.WithSession(session),
		table:      "`zero_user_infos`",
	}
}

func (m *defaultZeroUserInfosModel) Delete(ctx context.Context, id int64) error {
	zeroUserInfosIdKey := fmt.Sprintf("%s%v", cacheZeroUserInfosIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, zeroUserInfosIdKey)
	return err
}

func (m *defaultZeroUserInfosModel) FindOne(ctx context.Context, id int64) (*ZeroUserInfos, error) {
	zeroUserInfosIdKey := fmt.Sprintf("%s%v", cacheZeroUserInfosIdPrefix, id)
	var resp ZeroUserInfos
	err := m.QueryRowCtx(ctx, &resp, zeroUserInfosIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", zeroUserInfosRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultZeroUserInfosModel) Insert(ctx context.Context, data *ZeroUserInfos) (sql.Result, error) {
	zeroUserInfosIdKey := fmt.Sprintf("%s%v", cacheZeroUserInfosIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?)", m.table, zeroUserInfosRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.UserId, data.Email, data.Phone, data.UpdatedBy, data.CreatedBy, data.DeletedAt, data.DeletedFlag)
	}, zeroUserInfosIdKey)
	return ret, err
}

func (m *defaultZeroUserInfosModel) Update(ctx context.Context, data *ZeroUserInfos) error {
	zeroUserInfosIdKey := fmt.Sprintf("%s%v", cacheZeroUserInfosIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, zeroUserInfosRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.UserId, data.Email, data.Phone, data.UpdatedBy, data.CreatedBy, data.DeletedAt, data.DeletedFlag, data.Id)
	}, zeroUserInfosIdKey)
	return err
}

func (m *defaultZeroUserInfosModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheZeroUserInfosIdPrefix, primary)
}

func (m *defaultZeroUserInfosModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", zeroUserInfosRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultZeroUserInfosModel) tableName() string {
	return m.table
}