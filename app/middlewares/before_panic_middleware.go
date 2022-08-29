// author: wsasydevc <asydev@163.com>
// date: 2021-08-10

package middlewares

import (
	"github.com/asydevc/log/v2"
	"net/http"

	"github.com/kataras/iris/v12"
)

// 恢复Panic.
func BeforePanicMiddleware(ctx iris.Context) {
	// 注册Panic.
	// 当请求结束时, 会走到业务过程中
	defer func() {
		err := recover()
		// 无panic.
		if err == nil {
			return
		}
		// 写入日志.
		log.Panicfc(ctx, "系统异常: %v.", err)
		renderHttpErrorCode(ctx, http.StatusInternalServerError)
		ctx.ResponseWriter().FlushResponse()
		ctx.Next()
	}()
	ctx.Next()
}
