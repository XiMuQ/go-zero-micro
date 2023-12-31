// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	fileStorage "go-zero-micro/api/code/ucenterapi/internal/handler/fileStorage"
	login "go-zero-micro/api/code/ucenterapi/internal/handler/login"
	ucenter "go-zero-micro/api/code/ucenterapi/internal/handler/ucenter"
	"go-zero-micro/api/code/ucenterapi/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Check},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/getUserPageList",
					Handler: ucenter.GetUserPageListHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/getUserById",
					Handler: ucenter.GetUserByIdHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/addUser",
					Handler: ucenter.AddUserHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/path/:id/:name",
					Handler: ucenter.QueryUserByPathHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/ucenter"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/getUserByAccount",
				Handler: ucenter.GetUserByAccountHandler(serverCtx),
			},
		},
		rest.WithPrefix("/ucenter"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/loginByPassword",
				Handler: login.LoginByPasswordHandler(serverCtx),
			},
		},
		rest.WithPrefix("/login"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/fileUpload",
				Handler: fileStorage.FileUploadHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/fileDownload",
				Handler: fileStorage.FileDownloadHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/filePreview",
				Handler: fileStorage.FilePreviewHandler(serverCtx),
			},
		},
		rest.WithPrefix("/fileStorage"),
	)
}
