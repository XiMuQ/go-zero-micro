package utils

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
)

// GetUidFromCtx 从Context中获取指定key的信息，此处只处理了获取用户id
func GetUidFromCtx(ctx context.Context, key any) int64 {
	var uid int64
	if jsonUid, ok := ctx.Value(key).(json.Number); ok {
		if int64Uid, err := jsonUid.Int64(); err == nil {
			uid = int64Uid
		} else {
			logx.WithContext(ctx).Errorf("GetUidFromCtx err : %+v", err)
		}
	}
	return uid
}

// GetUidFromCtxInt64 从Context中获取指定key的信息，此处只处理了获取用户id
func GetUidFromCtxInt64(ctx context.Context, key any) int64 {
	var uid int64
	if jsonUid, ok := ctx.Value(key).(int64); ok {
		uid = jsonUid
	}
	return uid
}

// Base64Encode base64加密
func Base64Encode(message string) string {
	encodedMessage := base64.StdEncoding.EncodeToString([]byte(message))
	return encodedMessage
}

// Base64Decode base64解密
func Base64Decode(message string) (string, error) {
	decodedMessage, err := base64.StdEncoding.DecodeString(message)
	if err != nil {
		return "", err
	}
	return string(decodedMessage), nil
}
