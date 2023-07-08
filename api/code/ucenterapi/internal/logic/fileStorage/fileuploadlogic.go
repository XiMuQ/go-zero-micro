package fileStorage

import (
	"bytes"
	"context"
	"fmt"
	"github.com/jinzhu/copier"
	uuid "github.com/satori/go.uuid"
	"go-zero-micro/common/errorx"
	"go-zero-micro/common/utils"
	"go-zero-micro/rpc/code/ucenter/ucenter"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strings"

	"go-zero-micro/api/code/ucenterapi/internal/svc"
	"go-zero-micro/api/code/ucenterapi/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileUploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileUploadLogic {
	return &FileUploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileUploadLogic) FileUpload(request *http.Request, req *types.FileUploadReq) (resp *types.BaseModel, err error) {
	// todo: add your logic here and delete this line
	//return LocalFileToByte(l, request, requestBody)
	return FileToByte(l, request, req)
}

// LocalFileToByte 方式1：ioutil.ReadFile()转换成byte需要知道文件路径，因此会生成临时文件，不适合处理文件上传的场景
func LocalFileToByte(l *FileUploadLogic, request *http.Request, requestBody *types.FileUploadReq) (resp *types.BaseModel, err error) {
	SavePath := l.svcCtx.Config.UploadFile.SavePath //上传文件的存储路径
	utils.CreateDir(SavePath)
	files := request.MultipartForm.File["fileList"]
	res := &types.BaseModel{
		Data: "上传成功",
	}
	param := &ucenter.FileList{}
	copier.Copy(param, requestBody)
	rpcFileList := make([]*ucenter.FileInfo, 0)
	typeId := fmt.Sprintf("%d", requestBody.Type)
	// 遍历所有文件
	for _, fileHeader := range files {
		//获取文件大小
		fileSize := fileHeader.Size
		//获取文件名称带后缀
		fileNameWithSuffix := path.Base(fileHeader.Filename)
		//获取文件的后缀(文件类型)
		fileType := path.Ext(fileNameWithSuffix)
		//生成UUID防止文件被覆盖
		uuidName := typeId + "_" + strings.Replace(uuid.NewV4().String(), "-", "", -1)

		saveName := uuidName + fileType
		saveFullPath := SavePath + saveName
		logx.Infof("upload file: %+v, file size: %d", fileNameWithSuffix, fileSize)
		file, err := fileHeader.Open()
		tempFile, err := os.Create(saveFullPath)
		if err != nil {
			return nil, err
		}
		io.Copy(tempFile, file)
		//关闭文件
		file.Close()
		tempFile.Close()

		//方式1：ioutil.ReadFile()转换成byte需要知道文件路径，不适合处理文件上传的场景
		content, err := ioutil.ReadFile(saveFullPath)
		fileInfo := &ucenter.FileInfo{
			FileId:   requestBody.Id,
			FileName: fileNameWithSuffix,
			FileType: fileType,
			FileSize: fileSize,
			FileData: content,
		}
		err = os.Remove(saveFullPath)
		if err != nil {
			logx.Infof("%s：删除失败", fileNameWithSuffix)
		}
		rpcFileList = append(rpcFileList, fileInfo)
	}
	param.File = rpcFileList
	uploadRes, err := l.svcCtx.FileStorageRpc.FileUpload(l.ctx, param)
	if err != nil {
		return nil, err
	}
	res.Data = uploadRes.Data
	return res, nil
}

// FileToByte 方式2：转换成byte适合上传文件的场景
func FileToByte(l *FileUploadLogic, request *http.Request, requestBody *types.FileUploadReq) (resp *types.BaseModel, err error) {
	files := request.MultipartForm.File["fileList"]
	res := &types.BaseModel{
		Data: "上传成功",
	}
	param := &ucenter.FileList{}
	copier.Copy(param, requestBody)
	rpcFileList := make([]*ucenter.FileInfo, 0)
	// 遍历所有文件
	for _, fileHeader := range files {
		//获取文件大小
		fileSize := fileHeader.Size
		//获取文件名称带后缀
		fileNameWithSuffix := path.Base(fileHeader.Filename)
		//获取文件的后缀(文件类型)
		fileType := path.Ext(fileNameWithSuffix)
		logx.Infof("upload file: %+v, file size: %d", fileNameWithSuffix, fileSize)
		file, err := fileHeader.Open()
		if err != nil {
			return nil, err
		}
		//方式2：转换成byte适合上传文件的场景
		fil := make([][]byte, 0)
		var b int64 = 0
		// 通过for循环写入
		for {
			buffer := make([]byte, 1024)
			n, err := file.ReadAt(buffer, b)
			b = b + int64(n)
			fil = append(fil, buffer)
			if err != nil {
				fmt.Println(err.Error())
				break
			}
		}
		// 生成最后的文件字节流
		content := bytes.Join(fil, []byte(""))
		fileInfo := &ucenter.FileInfo{
			FileId:   requestBody.Id,
			FileName: fileNameWithSuffix,
			FileType: fileType,
			FileSize: fileSize,
			FileData: content,
		}
		rpcFileList = append(rpcFileList, fileInfo)
	}
	param.File = rpcFileList
	uploadRes, err := l.svcCtx.FileStorageRpc.FileUpload(l.ctx, param)
	if err != nil {
		return nil, errorx.NewCodeError(1000, err.Error())
	}
	res.Data = uploadRes.Data
	return res, nil
}
