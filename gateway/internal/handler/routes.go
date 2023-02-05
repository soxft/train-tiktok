// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	gateway "train-tiktok/gateway/internal/handler/gateway"
	identity "train-tiktok/gateway/internal/handler/identity"
	user "train-tiktok/gateway/internal/handler/user"
	video "train-tiktok/gateway/internal/handler/video"
	"train-tiktok/gateway/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/ping",
				Handler: gateway.PingHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/douyin/user/register",
				Handler: identity.RegisterHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/douyin/user/login",
				Handler: identity.LoginHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Auth},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/douyin/publish/action",
					Handler: video.PublishActionHandler(serverCtx),
				},
			}...,
		),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Auth},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/douyin/user",
					Handler: user.UserHandler(serverCtx),
				},
			}...,
		),
	)
}
