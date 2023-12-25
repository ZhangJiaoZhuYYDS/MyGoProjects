package handler

import (
	"book/service/search/api/internal/logic"
	"book/service/search/api/internal/svc"
	"book/service/user/api/response" // 1 替换为你真实的response包名
	"net/http"
)

func pingHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := logic.NewPingLogic(r.Context(), svcCtx)
		err := l.Ping()
		response.Response(w, nil, err) //②  自定义模板内容

	}
}
