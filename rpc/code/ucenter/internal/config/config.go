package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf

	JWT struct {
		AccessSecret string
		AccessExpire int64
	}

	MySQL struct {
		DataSource string
	}

	DefaultConfig DefaultConfig

	UploadFile UploadFile
}

type UploadFile struct {
	MaxFileNum  int64
	MaxFileSize int64
	SavePath    string
}

// DefaultConfig 默认配置
type DefaultConfig struct {
	//默认密码
	DefaultPassword string
}
