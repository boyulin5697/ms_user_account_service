package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"ms_user_account_service/api/internal/logic"
	"ms_user_account_service/api/internal/svc"
)

func DetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewDetailLogic(r.Context(), svcCtx)
		resp, err := l.Detail()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
