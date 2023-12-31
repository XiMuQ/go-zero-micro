// Code generated by goctl. DO NOT EDIT.
// Source: ucenter.proto

package server

import (
	"context"

	"go-zero-micro/rpc/code/ucenter/internal/logic/filestorage"
	"go-zero-micro/rpc/code/ucenter/internal/svc"
	"go-zero-micro/rpc/code/ucenter/ucenter"
)

type FileStorageServer struct {
	svcCtx *svc.ServiceContext
	ucenter.UnimplementedFileStorageServer
}

func NewFileStorageServer(svcCtx *svc.ServiceContext) *FileStorageServer {
	return &FileStorageServer{
		svcCtx: svcCtx,
	}
}

// 文件上传
func (s *FileStorageServer) FileUpload(ctx context.Context, in *ucenter.FileList) (*ucenter.BaseResp, error) {
	l := filestoragelogic.NewFileUploadLogic(ctx, s.svcCtx)
	return l.FileUpload(in)
}

// 文件下载
func (s *FileStorageServer) FileDownload(in *ucenter.FileInfo, stream ucenter.FileStorage_FileDownloadServer) error {
	l := filestoragelogic.NewFileDownloadLogic(stream.Context(), s.svcCtx)
	return l.FileDownload(in, stream)
}
