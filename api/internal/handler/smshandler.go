package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"ms_user_account_service/api/internal/logic"
	"ms_user_account_service/api/internal/svc"
)

func SMSHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewSMSLogic(r.Context(), svcCtx)
		resp, err := l.SMS()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
