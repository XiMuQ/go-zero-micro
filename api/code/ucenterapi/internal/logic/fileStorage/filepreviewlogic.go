package fileStorage

import (
	"context"
	"errors"
	"github.com/jinzhu/copier"
	"go-zero-micro/rpc/code/ucenter/ucenter"
	"net/http"

	"go-zero-micro/api/code/ucenterapi/internal/svc"
	"go-zero-micro/api/code/ucenterapi/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FilePreviewLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	writer http.ResponseWriter
}

func NewFilePreviewLogic(ctx context.Context, svcCtx *svc.ServiceContext, writer http.ResponseWriter) *FilePreviewLogic {
	return &FilePreviewLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		writer: writer,
	}
}

func (l *FilePreviewLogic) FilePreview(req *types.FileShowReq) error {
	// todo: add your logic here and delete this line
	param := &ucenter.FileInfo{}
	copier.Copy(param, req)
	downloadRes, err := l.svcCtx.FileStorageRpc.FileDownload(l.ctx, param)
	if err != nil {
		return errors.New("文件服务异常")
	}
	fileInfo, err := downloadRes.Recv()
	if err != nil {
		return errors.New("打开文件失败")
	}
	//fileName := fileInfo.FileName
	byteArr := fileInfo.FileData

	//如果是下载，则需要在Header中设置这两个参数
	//l.writer.Header().Add("Content-Type", "application/octet-stream")
	//l.writer.Header().Add("Content-Disposition", "attachment; filename= "+fileName)
	l.writer.Write(byteArr)
	return nil
}
