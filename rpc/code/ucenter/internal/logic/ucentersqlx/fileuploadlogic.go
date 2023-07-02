package ucentersqlxlogic

import (
	"context"
	uuid "github.com/satori/go.uuid"
	"go-zero-micro/common/utils"
	"io/ioutil"
	"path"
	"strings"

	"go-zero-micro/rpc/code/ucenter/internal/svc"
	"go-zero-micro/rpc/code/ucenter/ucenter"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileUploadLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFileUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileUploadLogic {
	return &FileUploadLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 文件上传
func (l *FileUploadLogic) FileUpload(in *ucenter.FileList) (*ucenter.BaseResp, error) {
	// todo: add your logic here and delete this line
	SavePath := l.svcCtx.Config.UploadFile.SavePath //上传文件的存储路径
	utils.CreateDir(SavePath)
	for _, fileInfo := range in.File {
		//获取文件名称带后缀
		fileNameWithSuffix := path.Base(fileInfo.FileName)
		//获取文件的后缀(文件类型)
		fileType := path.Ext(fileNameWithSuffix)
		//生成UUID防止文件被覆盖
		uuidName := strings.Replace(uuid.NewV4().String(), "-", "", -1)
		saveName := uuidName + fileType

		saveFullPath := SavePath + saveName
		err := ioutil.WriteFile(saveFullPath, fileInfo.FileData, 0644)
		if err != nil {
			// handle error
		}
	}
	return &ucenter.BaseResp{
		Data: "upload success",
	}, nil
}
