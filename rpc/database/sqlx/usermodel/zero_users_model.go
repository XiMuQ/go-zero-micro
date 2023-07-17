package usermodel

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go-zero-micro/common/utils"
)

var _ ZeroUsersModel = (*customZeroUsersModel)(nil)

type (
	// ZeroUsersModel is an interface to be customized, add more methods here,
	// and implement the added methods in customZeroUsersModel.
	ZeroUsersModel interface {
		zeroUsersModel

		TransCtx(ctx context.Context, fn func(context context.Context, session sqlx.Session) error) error
		CountCtx(ctx context.Context, data *ZeroUsers, beginTime, endTime string) (int64, error)
		FindPageListByParamCtx(ctx context.Context, data *ZeroUsers, beginTime, endTime string, current, pageSize int64) ([]*ZeroUsers, error)
		FindAllByParamCtx(ctx context.Context, data *ZeroUsers) ([]*ZeroUsers, error)
		FindOneByParamCtx(ctx context.Context, data *ZeroUsers) (*ZeroUsers, error)
		SaveCtx(ctx context.Context, data *ZeroUsers) (sql.Result, error)
		EditCtx(ctx context.Context, data *ZeroUsers) (sql.Result, error)
		DeleteDataCtx(ctx context.Context, data *ZeroUsers) error

		TransSaveCtx(ctx context.Context, session sqlx.Session, data *ZeroUsers) (sql.Result, error)
	}

	customZeroUsersModel struct {
		*defaultZeroUsersModel
	}
)

func (c customZeroUsersModel) TransCtx(ctx context.Context, fn func(context context.Context, session sqlx.Session) error) error {
	return c.conn.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		return fn(ctx, session)
	})
}

func (c customZeroUsersModel) CountCtx(ctx context.Context, data *ZeroUsers, beginTime, endTime string) (int64, error) {
	sql := fmt.Sprintf("SELECT count(*) as count FROM %s WHERE deleted_flag = %d", c.table, utils.DelNo)
	joinSql := utils.QuerySqlJoins(data)
	beginTimeSql := ""
	if beginTime != "" {
		beginTimeSql = fmt.Sprintf(" AND created_at >= %s", "'"+beginTime+"'")
	}
	endTimeSql := ""
	if endTime != "" {
		endTimeSql = fmt.Sprintf(" AND created_at <= %s", "'"+endTime+"'")
	}
	sql = sql + joinSql + beginTimeSql + endTimeSql

	var count int64
	err := c.conn.QueryRowCtx(ctx, &count, sql)
	switch err {
	case nil:
		return count, nil
	case sqlc.ErrNotFound:
		return 0, ErrNotFound
	default:
		return 0, err
	}
}

func (c customZeroUsersModel) FindPageListByParamCtx(ctx context.Context, data *ZeroUsers, beginTime, endTime string, current, pageSize int64) ([]*ZeroUsers, error) {
	sql := fmt.Sprintf("SELECT %s FROM %s WHERE deleted_flag = %d", zeroUsersRows, c.table, utils.DelNo)
	joinSql := utils.QuerySqlJoins(data)
	beginTimeSql := ""
	if beginTime != "" {
		beginTimeSql = fmt.Sprintf(" AND created_at >= %s", "'"+beginTime+"'")
	}
	endTimeSql := ""
	if endTime != "" {
		endTimeSql = fmt.Sprintf(" AND created_at <= %s", "'"+endTime+"'")
	}
	orderSql := " ORDER BY created_at DESC"
	limitSql := fmt.Sprintf(" LIMIT %d,%d", (current-1)*pageSize, pageSize)
	sql = sql + joinSql + beginTimeSql + endTimeSql + orderSql + limitSql

	var result []*ZeroUsers
	err := c.conn.QueryRowsCtx(ctx, &result, sql)
	switch err {
	case nil:
		return result, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (c customZeroUsersModel) FindAllByParamCtx(ctx context.Context, data *ZeroUsers) ([]*ZeroUsers, error) {
	sql := fmt.Sprintf("SELECT %s FROM %s WHERE deleted_flag = %d", zeroUsersRows, c.table, utils.DelNo)
	joinSql := utils.QuerySqlJoins(data)
	orderSql := " ORDER BY created_at DESC"
	sql = sql + joinSql + orderSql

	var result []*ZeroUsers
	err := c.conn.QueryRowsCtx(ctx, &result, sql)
	switch err {
	case nil:
		return result, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (c customZeroUsersModel) FindOneByParamCtx(ctx context.Context, data *ZeroUsers) (*ZeroUsers, error) {
	sql := fmt.Sprintf("SELECT %s FROM %s WHERE deleted_flag = %d", zeroUsersRows, c.table, utils.DelNo)
	joinSql := utils.QuerySqlJoins(data)
	orderSql := " ORDER BY created_at DESC"
	sql = sql + joinSql + orderSql

	var result ZeroUsers
	err := c.conn.QueryRowCtx(ctx, &result, sql)
	switch err {
	case nil:
		return &result, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (c customZeroUsersModel) SaveCtx(ctx context.Context, data *ZeroUsers) (sql.Result, error) {
	saveSql := utils.SaveSqlJoins(data, c.table)
	result, err := c.conn.ExecCtx(ctx, saveSql)
	return result, err
}

func (c customZeroUsersModel) EditCtx(ctx context.Context, data *ZeroUsers) (sql.Result, error) {
	editSql := utils.EditSqlJoins(data, c.table, data.Id)
	result, err := c.conn.ExecCtx(ctx, editSql)
	return result, err
}

func (c customZeroUsersModel) DeleteDataCtx(ctx context.Context, data *ZeroUsers) error {
	UpdateTime := data.UpdatedAt.Format(utils.DateTimeFormat)
	sql := fmt.Sprintf("UPDATE %s SET deleted_flag = %d,deleted_at= %s WHERE id = %d", c.table, utils.DelYes, "'"+UpdateTime+"'", data.Id)
	_, err := c.conn.ExecCtx(ctx, sql)
	return err
}

func (c customZeroUsersModel) TransSaveCtx(ctx context.Context, session sqlx.Session, data *ZeroUsers) (sql.Result, error) {
	saveSql := utils.SaveSqlJoins(data, c.table)
	//result, err := c.conn.ExecCtx(ctx, saveSql)
	//return result, err
	result, err := session.ExecCtx(ctx, saveSql)
	return result, err
}

// NewZeroUsersModel returns a model for the database table.
func NewZeroUsersModel(conn sqlx.SqlConn) ZeroUsersModel {
	return &customZeroUsersModel{
		defaultZeroUsersModel: newZeroUsersModel(conn),
	}
}
