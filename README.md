# go-zero-micro 代码示例

# 1 相关命令
**1.1 API服务模块**
1. `goctl`使用`api`文件生成`api服务`命令：

```bash
\go-zero-micro\api> goctl api go -api ./doc/all.api -dir ./code/ucenterapi
```

**1.2 RPC服务模块**
1. `goctl`使用`protoc`文件生成`rpc服务`命令：

```bash
goctl rpc protoc ./proto/ucenter.proto --go_out=./code/ucenter --go-grpc_out=./code/ucenter --zrpc_out=./code/ucenter --multiple
```
注意：`--go_out`、`--go-grpc_out`、`--zrpc_out` 三者配置的路径需要完全一致，否则会报下列错误。
```bash
the output of pb.go and _grpc.pb.go must not be the same with --zrpc_out
```

**1.3 model服务模块**

1、生成sqlx代码命令

单表：
```bash
goctl model mysql ddl -src="./rpc/database/sql/user/zero_users.sql" -dir="./rpc/database/sqlx/usermodel" -style=go_zero
```
多表：
```bash
goctl model mysql ddl -src="./rpc/database/sql/user/zero_*.sql" -dir="./rpc/database/sqlx/usermodel" -style=go_zero
```
使用数据库连接方式：
```bash
goctl model mysql datasource -url="root:root@tcp(127.0.0.1:3357)/go-zero-micro" -table="zero_users" -dir="./rpc/database/sqlx/usermodel"
```
