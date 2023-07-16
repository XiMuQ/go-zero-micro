package utils

import (
	"fmt"
	"reflect"
	"strings"
	"time"
)

// DelNo 正常使用
var DelNo int64 = 1

// DelYes 软删除
var DelYes int64 = 2

// QuerySqlJoins 根据查询条件拼接sql，使用泛型更加通用
func QuerySqlJoins[T any](data *T) string {
	typ := reflect.TypeOf(data).Elem()  //指针类型需要加 Elem()
	val := reflect.ValueOf(data).Elem() //指针类型需要加 Elem()
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

// SaveSqlJoins 根据实际参数拼接sql，使用泛型更加通用
func SaveSqlJoins[T any](data *T, table string) string {
	typ := reflect.TypeOf(data).Elem()  //指针类型需要加 Elem()
	val := reflect.ValueOf(data).Elem() //指针类型需要加 Elem()
	fieldNum := val.NumField()

	names := ""
	values := ""
	for i := 1; i < fieldNum; i++ {
		Field := val.Field(i)
		colType := Field.Type().String()
		if colType == "int64" {
			//if Field.Int() > 0 {
			//	names += fmt.Sprintf("`%s`,", typ.Field(i).Tag.Get("db"))
			//	values += fmt.Sprintf("%d,", Field.Int())
			//}
			names += fmt.Sprintf("`%s`,", typ.Field(i).Tag.Get("db"))
			values += fmt.Sprintf("%d,", Field.Int())
		} else if colType == "string" {
			names += fmt.Sprintf("`%s`,", typ.Field(i).Tag.Get("db"))
			values += fmt.Sprintf("'%s',", Field.String())
		} else if colType == "time.Time" {
			value := Field.Interface().(time.Time)
			if !value.IsZero() {
				names += fmt.Sprintf("`%s`,", typ.Field(i).Tag.Get("db"))
				values += fmt.Sprintf("'%s',", value.Format(DateTimeFormat))
			}
		}
	}
	names = strings.TrimRight(names, ",")
	values = strings.TrimRight(values, ",")
	sql := fmt.Sprintf("INSERT INTO %s(%s) VALUE(%s)", table, names, values)
	return sql
}

// EditSqlJoins 根据实际参数拼接sql，使用泛型更加通用
func EditSqlJoins[T any](data *T, table string, Id int64) string {
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
				names += fmt.Sprintf("`%s`='%s',", typ.Field(i).Tag.Get("db"), value.Format(DateTimeFormat))
			}
		}
	}
	names = strings.TrimRight(names, ",")
	sql := fmt.Sprintf("UPDATE %s SET deleted_flag = %d, %s WHERE id = %d", table, DelNo, names, Id)
	return sql
}
