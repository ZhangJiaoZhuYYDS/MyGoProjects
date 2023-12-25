package handler
// 控制层（读取请求api，调用业务层，加载配置到上下文，自定义响应内容）

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-demo/greet/internal/logic"
	"go-zero-demo/greet/internal/svc"
	"go-zero-demo/greet/internal/types"
	"go-zero-demo/response" // 1 替换为你真实的response包名
	"net/http"
)

func GreetHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		// 实例化业务层对象
		l := logic.NewGreetLogic(r.Context(), svcCtx)
		// 调用业务层方法
		resp, err := l.Greet(&req)
		// 调用自定义响应内容
		response.Response(w, resp, err) //②  自定义模板内容
	}
}
