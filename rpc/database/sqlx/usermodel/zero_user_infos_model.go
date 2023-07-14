package usermodel

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go-zero-micro/common/utils"
	"reflect"
	"strings"
	"time"
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

/*
*
根据条件拼接的sql
*/
func userInfoSqlJoins(queryModel *ZeroUserInfos) string {
	typ := reflect.TypeOf(queryModel).Elem()  //指针类型需要加 Elem()
	val := reflect.ValueOf(queryModel).Elem() //指针类型需要加 Elem()
	fieldNum := val.NumField()

	sql := ""
	for i := 0; i < fieldNum; i++ {
		Field := val.Field(i)
		colType := Field.Type().String()
		colName := typ.Field(i).Tag.Get("db")

		if colType == "int64" {
			if Field.Int() > 0 {
				sql += fmt.Sprintf(" AND %s=%d", colName, Field.Int())
			}
		} else if colType == "string" {
			if Field.String() != "" {
				sql += fmt.Sprintf(" AND %s LIKE %s", colName, "'%"+Field.String()+"%'")
			}
		} else if colType == "time.Time" {
			value := Field.Interface().(time.Time)
			if !value.IsZero() {
				sql += fmt.Sprintf(" AND %s='%s'", colName, Field.String())
			}
		}
	}
	return sql
}

func (c customZeroUserInfosModel) Count(data *ZeroUserInfos, beginTime, endTime string) (int64, error) {
	sql := fmt.Sprintf("SELECT count(*) as count FROM %s WHERE deleted_flag = %d", c.table, utils.DelNo)
	joinSql := userInfoSqlJoins(data)
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
	joinSql := userInfoSqlJoins(data)
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
	joinSql := userInfoSqlJoins(data)
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
	joinSql := userInfoSqlJoins(data)
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
	typ := reflect.TypeOf(data).Elem()  //指针类型需要加 Elem()
	val := reflect.ValueOf(data).Elem() //指针类型需要加 Elem()
	fieldNum := val.NumField()

	names := ""
	values := ""
	for i := 1; i < fieldNum; i++ {
		Field := val.Field(i)
		colType := Field.Type().String()
		if colType == "int64" {
			if Field.Int() > 0 {
				names += fmt.Sprintf("`%s`,", typ.Field(i).Tag.Get("db"))
				values += fmt.Sprintf("%d,", Field.Int())
			}
		} else if colType == "string" {
			names += fmt.Sprintf("`%s`,", typ.Field(i).Tag.Get("db"))
			values += fmt.Sprintf("'%s',", Field.String())
		} else if colType == "time.Time" {
			value := Field.Interface().(time.Time)
			if !value.IsZero() {
				names += fmt.Sprintf("`%s`,", typ.Field(i).Tag.Get("db"))
				values += fmt.Sprintf("'%s',", value.Format(utils.DateTimeFormat))
			}
		}
	}
	names = strings.TrimRight(names, ",")
	values = strings.TrimRight(values, ",")
	saveSql := fmt.Sprintf("INSERT INTO %s(%s) VALUE(%s)", c.table, names, values)
	result, err := c.conn.ExecCtx(ctx, saveSql)
	return result, err
}

func (c customZeroUserInfosModel) Edit(ctx context.Context, data *ZeroUserInfos) (sql.Result, error) {
	typ := reflect.TypeOf(data).Elem()  //指针类型需要加 Elem()
	val := reflect.ValueOf(data).Elem() //指针类型需要加 Elem()
	fieldNum := val.NumField()

	names := ""
	for i := 1; i < fieldNum; i++ {
		Field := val.Field(i)
		colType := Field.Type().String()
		if colType == "int64" {
			if Field.Int() > 0 {
				names += fmt.Sprintf("`%s`=%d,", typ.Field(i).Tag.Get("db"), Field.Int())
			}
		} else if colType == "string" {
			names += fmt.Sprintf("`%s`='%s',", typ.Field(i).Tag.Get("db"), Field.String())
		} else if colType == "time.Time" {
			value := Field.Interface().(time.Time)
			if !value.IsZero() {
				names += fmt.Sprintf("`%s`='%s',", typ.Field(i).Tag.Get("db"), value.Format(utils.DateTimeFormat))
			}
		}
	}
	names = strings.TrimRight(names, ",")
	sql := fmt.Sprintf("UPDATE %s SET deleted_flag = %d, %s WHERE id = %d", c.table, utils.DelNo, names, data.Id)
	result, err := c.conn.ExecCtx(ctx, sql)
	return result, err
}

func (c customZeroUserInfosModel) DeleteData(ctx context.Context, data *ZeroUserInfos) error {
	UpdateTime := data.UpdatedAt.Format(utils.DateTimeFormat)
	sql := fmt.Sprintf("UPDATE %s SET deleted_flag = %d,deleted_at= %s WHERE id = %d", c.table, utils.DelYes, "'"+UpdateTime+"'", data.Id)
	_, err := c.conn.ExecCtx(ctx, sql)
	return err
}

func (c customZeroUserInfosModel) TransSaveCtx(ctx context.Context, session sqlx.Session, data *ZeroUserInfos) (sql.Result, error) {
	typ := reflect.TypeOf(data).Elem()  //指针类型需要加 Elem()
	val := reflect.ValueOf(data).Elem() //指针类型需要加 Elem()
	fieldNum := val.NumField()

	names := ""
	values := ""
	for i := 1; i < fieldNum; i++ {
		Field := val.Field(i)
		colType := Field.Type().String()
		if colType == "int64" {
			if Field.Int() > 0 {
				names += fmt.Sprintf("`%s`,", typ.Field(i).Tag.Get("db"))
				values += fmt.Sprintf("%d,", Field.Int())
			}
		} else if colType == "string" {
			names += fmt.Sprintf("`%s`,", typ.Field(i).Tag.Get("db"))
			values += fmt.Sprintf("'%s',", Field.String())
		} else if colType == "time.Time" {
			value := Field.Interface().(time.Time)
			if !value.IsZero() {
				names += fmt.Sprintf("`%s`,", typ.Field(i).Tag.Get("db"))
				values += fmt.Sprintf("'%s',", value.Format(utils.DateTimeFormat))
			}
		}
	}
	names = strings.TrimRight(names, ",")
	values = strings.TrimRight(values, ",")
	saveSql := fmt.Sprintf("INSERT INTO %s(%s) VALUE(%s)", c.table, names, values)

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
