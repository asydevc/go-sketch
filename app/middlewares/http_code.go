// author: wsasydevc <asydev@163.com>
// date: 2021-08-06

package middlewares

import (
	"github.com/kataras/iris/v12"
)

func HttpCodeMiddleware(ctx iris.Context) {
	renderHttpErrorCode(ctx, ctx.GetStatusCode())
	ctx.Next()
}
