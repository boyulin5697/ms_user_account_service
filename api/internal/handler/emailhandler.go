package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"ms_user_account_service/api/internal/logic"
	"ms_user_account_service/api/internal/svc"
)

func EmailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewEmailLogic(r.Context(), svcCtx)
		resp, err := l.Email()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
