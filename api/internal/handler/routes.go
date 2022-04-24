// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"ms_user_account_service/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/register",
				Handler: RegisterHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/login",
				Handler: LoginHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/getDetail",
				Handler: DetailHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/getEmailCode",
				Handler: EmailHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/getPhoneCode",
				Handler: SMSHandler(serverCtx),
			},
		},
	)
}
