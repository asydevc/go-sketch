// author: wsasydevc <asydev@163.com>
// date: 2021-08-10

package middlewares

import (
	"github.com/asydevc/log/v2"
	"net/http"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"

	"github.com/asydevc/go-sketch/app"
)

// 控制器执行前.
//
// 执行控制器逻辑前触发本方法.
// 隐性调用, 禁止手工触发.
func BeforeControllerMiddleware(ctx iris.Context) {
	// 1. 基础日志.
	log.Infofc(ctx, "request start, headers: %s.", ctx.Request().Header)
	// 2. 记录入参.
	//    Post/Put/Delete.
	if app.Config.LogHttpPayload {
		switch ctx.Method() {
		case http.MethodPost, http.MethodPut, http.MethodDelete:
			if body, err := context.GetBody(ctx.Request(), true); err == nil {
				log.Infofc(ctx, "request payload: %s.", body)
			}
		}
	}
	// 3. 下探.
	ctx.Next()
}
