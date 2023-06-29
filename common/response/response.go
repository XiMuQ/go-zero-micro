package response

import (
	"context"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

type Body struct {
	Code    int         `json:"code"`
	Msg     string      `json:"msg"`
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
}

func Response(ctx context.Context, w http.ResponseWriter, resp interface{}, err error) {
	if err != nil {
		httpx.OkJsonCtx(ctx, w, err)
	} else {
		var body Body
		body.Code = 200
		body.Msg = "success"
		body.Success = true
		body.Data = resp
		httpx.OkJsonCtx(ctx, w, body)
	}
}
