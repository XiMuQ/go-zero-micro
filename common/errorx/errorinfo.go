package errorx

import "errors"

var (
	RedisErrRecordNotFound = errors.New("data not exist in redis")
)
