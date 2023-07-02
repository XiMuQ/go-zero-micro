package ucenter

import (
	"go-zero-micro/common/response"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-micro/api/code/ucenterapi/internal/logic/ucenter"
	"go-zero-micro/api/code/ucenterapi/internal/svc"
	"go-zero-micro/api/code/ucenterapi/internal/types"
)

func FileUploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FileUploadReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := ucenter.NewFileUploadLogic(r.Context(), svcCtx)
		resp, err := l.FileUpload(r, &req)
		//if err != nil {
		//	httpx.ErrorCtx(r.Context(), w, err)
		//} else {
		//	httpx.OkJsonCtx(r.Context(), w, resp)
		//}
		response.Response(r.Context(), w, resp, err)
	}
}
