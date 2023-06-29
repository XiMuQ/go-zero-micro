package utils

import (
	"github.com/zeromicro/go-zero/core/logx"
	"os"
)

// HasDir 判断文件夹是否存在
func HasDir(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// CreateDir 创建文件夹
func CreateDir(path string) {
	exist, err := HasDir(path)
	if err != nil {
		logx.Infof("获取文件夹异常: %+v", err)
		return
	}
	if !exist {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			logx.Infof("创建文件夹异常!: %+v", err)
		} else {
			logx.Infof("创建文件夹成功!")
		}
	} else {
		//logx.Infof("文件夹已存在! %+v", err)
	}
}
