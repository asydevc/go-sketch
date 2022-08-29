// author: wsasydevc <asydev@163.com>
// date: 2021-08-10

package middlewares

import (
	"github.com/asydevc/log/v2"
	"time"

	"github.com/asydevc/log/v2/interfaces"
	"github.com/kataras/iris/v12"

	"github.com/asydevc/go-sketch/app"
)

// 请求链.
//
// IRIS框架收到HTTP请求时, 第1步先执行本中间件, 用于将
// OpenTracing标记加入到ctx中, 业务执行过程写log时会
// 提取Tracing信息, 保证同一个请求下的日志包含相同的请求
// 链路ID.
func BeforeTracingMiddleware(ctx iris.Context) {
	tracing := log.NewTracing().UseRequest(ctx.Request())
	ctx.Values().Set(interfaces.OpenTracingKey, tracing)
	ctx.Values().Set("RequestDispatchedTime", time.Now())
	ctx.ResponseWriter().Header().Set("request-id", tracing.GetTraceId())
	ctx.ResponseWriter().Header().Set("server", app.Config.Software)
	ctx.Next()
}
