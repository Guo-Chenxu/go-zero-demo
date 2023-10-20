package user

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"user_api/internal/logic/user"
	"user_api/internal/svc"
	"user_api/internal/types"
)

func SetUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SetUserReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewSetUserLogic(r.Context(), svcCtx)
		resp, err := l.SetUser(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
