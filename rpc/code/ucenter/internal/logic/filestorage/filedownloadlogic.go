package filestoragelogic

import (
	"context"
	"errors"
	"github.com/jinzhu/copier"
	"os"

	"go-zero-micro/rpc/code/ucenter/internal/svc"
	"go-zero-micro/rpc/code/ucenter/ucenter"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileDownloadLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFileDownloadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileDownloadLogic {
	return &FileDownloadLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 文件下载
func (l *FileDownloadLogic) FileDownload(in *ucenter.FileInfo, stream ucenter.FileStorage_FileDownloadServer) error {
	// todo: add your logic here and delete this line
	SavePath := l.svcCtx.Config.UploadFile.SavePath //上传文件的存储路径

	filePath := SavePath + in.FileUrl
	_, err := os.Stat(filePath)
	if err != nil || os.IsNotExist(err) {
		return errors.New("文件不存在")
	}
	bytes, err := os.ReadFile(filePath)
	if err != nil {
		return errors.New("读取文件失败")
	}

	response := &ucenter.FileInfo{}
	copier.Copy(response, in)
	response.FileName = "go-zero.png"
	response.FileData = bytes
	if err := stream.Send(response); err != nil {
		return err
	}
	return nil
}
