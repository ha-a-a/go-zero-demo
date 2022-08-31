package handler

import (
	"go-zero-demo/book/user/api/internal/logic"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-demo/book/user/api/internal/svc"
	"go-zero-demo/book/user/api/internal/types"
)

func LoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
