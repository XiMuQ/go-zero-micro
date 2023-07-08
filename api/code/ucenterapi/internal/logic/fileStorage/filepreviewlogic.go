package fileStorage

import (
	"context"

	"go-zero-micro/api/code/ucenterapi/internal/svc"
	"go-zero-micro/api/code/ucenterapi/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FilePreviewLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFilePreviewLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FilePreviewLogic {
	return &FilePreviewLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FilePreviewLogic) FilePreview(req *types.FileShowReq) error {
	// todo: add your logic here and delete this line

	return nil
}
