package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
	UploadFile UploadFile
	UCenterRpc zrpc.RpcClientConf
}

type UploadFile struct {
	MaxFileNum  int64
	MaxFileSize int64
	SavePath    string
}
