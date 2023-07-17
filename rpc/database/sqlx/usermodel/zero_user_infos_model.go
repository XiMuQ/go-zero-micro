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
		CountCtx(ctx context.Context, data *ZeroUserInfos, beginTime, endTime string) (int64, error)
		FindPageListByParamCtx(ctx context.Context, data *ZeroUserInfos, beginTime, endTime string, current, pageSize int64) ([]*ZeroUserInfos, error)
		FindAllByParamCtx(ctx context.Context, data *ZeroUserInfos) ([]*ZeroUserInfos, error)
		FindOneByParamCtx(ctx context.Context, data *ZeroUserInfos) (*ZeroUserInfos, error)
		SaveCtx(ctx context.Context, data *ZeroUserInfos) (sql.Result, error)
		EditCtx(ctx context.Context, data *ZeroUserInfos) (sql.Result, error)
		DeleteDataCtx(ctx context.Context, data *ZeroUserInfos) error

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

func (c customZeroUserInfosModel) CountCtx(ctx context.Context, data *ZeroUserInfos, beginTime, endTime string) (int64, error) {
	querySql := fmt.Sprintf("SELECT count(*) as count FROM %s WHERE deleted_flag = %d", c.table, utils.DelNo)
	joinSql := utils.QuerySqlJoins(data)
	beginTimeSql := ""
	if beginTime != "" {
		beginTimeSql = fmt.Sprintf(" AND created_at >= %s", "'"+beginTime+"'")
	}
	endTimeSql := ""
	if endTime != "" {
		endTimeSql = fmt.Sprintf(" AND created_at <= %s", "'"+endTime+"'")
	}
	querySql = querySql + joinSql + beginTimeSql + endTimeSql

	var count int64
	err := c.conn.QueryRowCtx(ctx, &count, querySql)
	switch err {
	case nil:
		return count, nil
	case sqlc.ErrNotFound:
		return 0, ErrNotFound
	default:
		return 0, err
	}
}

func (c customZeroUserInfosModel) FindPageListByParamCtx(ctx context.Context, data *ZeroUserInfos, beginTime, endTime string, current, pageSize int64) ([]*ZeroUserInfos, error) {
	querySql := fmt.Sprintf("SELECT %s FROM %s WHERE deleted_flag = %d", zeroUsersRows, c.table, utils.DelNo)
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
	querySql = querySql + joinSql + beginTimeSql + endTimeSql + orderSql + limitSql

	var result []*ZeroUserInfos
	err := c.conn.QueryRowsCtx(ctx, &result, querySql)
	switch err {
	case nil:
		return result, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (c customZeroUserInfosModel) FindAllByParamCtx(ctx context.Context, data *ZeroUserInfos) ([]*ZeroUserInfos, error) {
	querySql := fmt.Sprintf("SELECT %s FROM %s WHERE deleted_flag = %d", zeroUsersRows, c.table, utils.DelNo)
	joinSql := utils.QuerySqlJoins(data)
	orderSql := " ORDER BY created_at DESC"
	querySql = querySql + joinSql + orderSql

	var result []*ZeroUserInfos
	err := c.conn.QueryRowsCtx(ctx, &result, querySql)
	switch err {
	case nil:
		return result, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (c customZeroUserInfosModel) FindOneByParamCtx(ctx context.Context, data *ZeroUserInfos) (*ZeroUserInfos, error) {
	querySql := fmt.Sprintf("SELECT %s FROM %s WHERE deleted_flag = %d", zeroUsersRows, c.table, utils.DelNo)
	joinSql := utils.QuerySqlJoins(data)
	orderSql := " ORDER BY created_at DESC"
	querySql = querySql + joinSql + orderSql

	var result ZeroUserInfos
	err := c.conn.QueryRowCtx(ctx, &result, querySql)
	switch err {
	case nil:
		return &result, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (c customZeroUserInfosModel) SaveCtx(ctx context.Context, data *ZeroUserInfos) (sql.Result, error) {
	saveSql := utils.SaveSqlJoins(data, c.table)
	result, err := c.conn.ExecCtx(ctx, saveSql)
	return result, err
}

func (c customZeroUserInfosModel) EditCtx(ctx context.Context, data *ZeroUserInfos) (sql.Result, error) {
	editSql := utils.EditSqlJoins(data, c.table, data.Id)
	result, err := c.conn.ExecCtx(ctx, editSql)
	return result, err
}

func (c customZeroUserInfosModel) DeleteDataCtx(ctx context.Context, data *ZeroUserInfos) error {
	UpdateTime := data.UpdatedAt.Format(utils.DateTimeFormat)
	deleteSql := fmt.Sprintf("UPDATE %s SET deleted_flag = %d,deleted_at= %s WHERE id = %d", c.table, utils.DelYes, "'"+UpdateTime+"'", data.Id)
	_, err := c.conn.ExecCtx(ctx, deleteSql)
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
