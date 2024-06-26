// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	user "class01/user-api/internal/handler/user"
	"class01/user-api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				// 获取用户信息
				Method:  http.MethodGet,
				Path:    "/user/info",
				Handler: user.UserInfoHandler(serverCtx),
			},
			{
				// 修改用户信息
				Method:  http.MethodPost,
				Path:    "/user/update",
				Handler: user.UserUpdateHandler(serverCtx),
			},
		},
		rest.WithPrefix("/userapi/v1"),
	)
}
