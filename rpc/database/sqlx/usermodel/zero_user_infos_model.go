package usermodel

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go-zero-micro/common/utils"
)

var _ ZeroUserInfosModel = (*customZeroUserInfosModel)(nil)

type (
	// ZeroUserInfosModel is an interface to be customized, add more methods here,
	// and implement the added methods in customZeroUserInfosModel.
	ZeroUserInfosModel interface {
		zeroUserInfosModel

		TransCtx(ctx context.Context, fn func(context context.Context, session sqlx.Session) error) error
		Count(data *ZeroUserInfos, beginTime, endTime string) (int64, error)
		FindPageListByParam(data *ZeroUserInfos, beginTime, endTime string, current, pageSize int64) ([]*ZeroUserInfos, error)
		FindAllByParam(data *ZeroUserInfos) ([]*ZeroUserInfos, error)
		FindOneByParam(data *ZeroUserInfos) (*ZeroUserInfos, error)
		Save(ctx context.Context, data *ZeroUserInfos) (sql.Result, error)
		Edit(ctx context.Context, data *ZeroUserInfos) (sql.Result, error)
		DeleteData(ctx context.Context, data *ZeroUserInfos) error

		TransSaveCtx(ctx context.Context, session sqlx.Session, data *ZeroUserInfos) (sql.Result, error)
	}

	customZeroUserInfosModel struct {
		*defaultZeroUserInfosModel
	}
)

func (c customZeroUserInfosModel) TransCtx(ctx context.Context, fn func(context context.Context, session sqlx.Session) error) error {
	return c.conn.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		return fn(ctx, session)
	})
}

func (c customZeroUserInfosModel) Count(data *ZeroUserInfos, beginTime, endTime string) (int64, error) {
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

func (c customZeroUserInfosModel) FindPageListByParam(data *ZeroUserInfos, beginTime, endTime string, current, pageSize int64) ([]*ZeroUserInfos, error) {
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

	var result []*ZeroUserInfos
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

func (c customZeroUserInfosModel) FindAllByParam(data *ZeroUserInfos) ([]*ZeroUserInfos, error) {
	sql := fmt.Sprintf("SELECT %s FROM %s WHERE deleted_flag = %d", zeroUsersRows, c.table, utils.DelNo)
	joinSql := utils.QuerySqlJoins(data)
	orderSql := " ORDER BY created_at DESC"
	sql = sql + joinSql + orderSql

	var result []*ZeroUserInfos
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

func (c customZeroUserInfosModel) FindOneByParam(data *ZeroUserInfos) (*ZeroUserInfos, error) {
	sql := fmt.Sprintf("SELECT %s FROM %s WHERE deleted_flag = %d", zeroUsersRows, c.table, utils.DelNo)
	joinSql := utils.QuerySqlJoins(data)
	orderSql := " ORDER BY created_at DESC"
	sql = sql + joinSql + orderSql

	var result ZeroUserInfos
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

func (c customZeroUserInfosModel) Save(ctx context.Context, data *ZeroUserInfos) (sql.Result, error) {
	saveSql := utils.SaveSqlJoins(data, c.table)
	result, err := c.conn.ExecCtx(ctx, saveSql)
	return result, err
}

func (c customZeroUserInfosModel) Edit(ctx context.Context, data *ZeroUserInfos) (sql.Result, error) {
	editSql := utils.EditSqlJoins(data, c.table, data.Id)
	result, err := c.conn.ExecCtx(ctx, editSql)
	return result, err
}

func (c customZeroUserInfosModel) DeleteData(ctx context.Context, data *ZeroUserInfos) error {
	UpdateTime := data.UpdatedAt.Format(utils.DateTimeFormat)
	sql := fmt.Sprintf("UPDATE %s SET deleted_flag = %d,deleted_at= %s WHERE id = %d", c.table, utils.DelYes, "'"+UpdateTime+"'", data.Id)
	_, err := c.conn.ExecCtx(ctx, sql)
	return err
}

func (c customZeroUserInfosModel) TransSaveCtx(ctx context.Context, session sqlx.Session, data *ZeroUserInfos) (sql.Result, error) {
	saveSql := utils.SaveSqlJoins(data, c.table)
	//result, err := c.conn.ExecCtx(ctx, saveSql)
	//return result, err
	result, err := session.ExecCtx(ctx, saveSql)
	return result, err
}

// NewZeroUserInfosModel returns a model for the database table.
func NewZeroUserInfosModel(conn sqlx.SqlConn) ZeroUserInfosModel {
	return &customZeroUserInfosModel{
		defaultZeroUserInfosModel: newZeroUserInfosModel(conn),
	}
}
