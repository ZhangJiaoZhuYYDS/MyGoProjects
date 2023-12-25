package handler

import (
	"book/service/user/api/internal/logic"
	"book/service/user/api/internal/svc"
	"book/service/user/api/internal/types"
	"book/service/user/api/response" // 1 替换为你真实的response包名
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func loginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req)
		response.Response(w, resp, err) //②  自定义模板内容
	}
}
