package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf

	JWT struct {
		AccessSecret string
		AccessExpire int64
	}
	UploadFile UploadFile
}

type UploadFile struct {
	MaxFileNum  int64
	MaxFileSize int64
	SavePath    string
}
