// author: wsasydevc <asydev@163.com>
// date: 2021-08-11

package middlewares

import (
	"github.com/asydevc/log/v2"
	"time"

	"github.com/kataras/iris/v12"

	"github.com/asydevc/go-sketch/app"
)

// 控制器完成后.
//
// 控制器逻辑执行完成后触发本方法.
// 隐性调用, 禁止手工触发.
func AfterControllerMiddleware(ctx iris.Context) {
	if app.Config.LogHttpResponseBody {
		log.Infofc(ctx, "[d=%f] response end.",
			time.Now().Sub(ctx.Values().Get("RequestDispatchedTime").(time.Time)).Seconds(),
		)
	} else {
		log.Infofc(
			ctx, "[d=%f] request end.",
			time.Now().Sub(ctx.Values().Get("RequestDispatchedTime").(time.Time)).Seconds(),
		)
	}
	ctx.Next()
}
