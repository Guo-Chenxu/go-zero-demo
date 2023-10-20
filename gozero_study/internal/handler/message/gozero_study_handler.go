package message

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"gozero_study/internal/logic/message"
	"gozero_study/internal/svc"
	"gozero_study/internal/types"
)

func Gozero_studyHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := message.NewGozero_studyLogic(r.Context(), svcCtx)
		resp, err := l.Gozero_study(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
