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
		Count(data *ZeroUsers, beginTime, endTime string) (int64, error)
		FindPageListByParam(data *ZeroUsers, beginTime, endTime string, current, pageSize int64) ([]*ZeroUsers, error)
		FindAllByParam(data *ZeroUsers) ([]*ZeroUsers, error)
		FindOneByParam(data *ZeroUsers) (*ZeroUsers, error)
		Save(ctx context.Context, data *ZeroUsers) (sql.Result, error)
		Edit(ctx context.Context, data *ZeroUsers) (sql.Result, error)
		DeleteData(ctx context.Context, data *ZeroUsers) error

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

func (c customZeroUsersModel) Count(data *ZeroUsers, beginTime, endTime string) (int64, error) {
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
	err := c.conn.QueryRow(&count, sql)
	switch err {
	case nil:
		return count, nil
	case sqlc.ErrNotFound:
		return 0, ErrNotFound
	default:
		return 0, err
	}
}

func (c customZeroUsersModel) FindPageListByParam(data *ZeroUsers, beginTime, endTime string, current, pageSize int64) ([]*ZeroUsers, error) {
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
	err := c.conn.QueryRows(&result, sql)
	switch err {
	case nil:
		return result, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (c customZeroUsersModel) FindAllByParam(data *ZeroUsers) ([]*ZeroUsers, error) {
	sql := fmt.Sprintf("SELECT %s FROM %s WHERE deleted_flag = %d", zeroUsersRows, c.table, utils.DelNo)
	joinSql := utils.QuerySqlJoins(data)
	orderSql := " ORDER BY created_at DESC"
	sql = sql + joinSql + orderSql

	var result []*ZeroUsers
	err := c.conn.QueryRows(&result, sql)
	switch err {
	case nil:
		return result, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (c customZeroUsersModel) FindOneByParam(data *ZeroUsers) (*ZeroUsers, error) {
	sql := fmt.Sprintf("SELECT %s FROM %s WHERE deleted_flag = %d", zeroUsersRows, c.table, utils.DelNo)
	joinSql := utils.QuerySqlJoins(data)
	orderSql := " ORDER BY created_at DESC"
	sql = sql + joinSql + orderSql

	var result ZeroUsers
	err := c.conn.QueryRow(&result, sql)
	switch err {
	case nil:
		return &result, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (c customZeroUsersModel) Save(ctx context.Context, data *ZeroUsers) (sql.Result, error) {
	saveSql := utils.SaveSqlJoins(data, c.table)
	result, err := c.conn.ExecCtx(ctx, saveSql)
	return result, err
}

func (c customZeroUsersModel) Edit(ctx context.Context, data *ZeroUsers) (sql.Result, error) {
	editSql := utils.EditSqlJoins(data, c.table, data.Id)
	result, err := c.conn.ExecCtx(ctx, editSql)
	return result, err
}

func (c customZeroUsersModel) DeleteData(ctx context.Context, data *ZeroUsers) error {
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
