// author: wsasydevc <asydev@163.com>
// date: 2021-08-05

// Package middleware.
package middlewares

import (
	"fmt"
	"net/http"

	"github.com/kataras/iris/v12"

	"github.com/asydevc/go-sketch/app"
)

// Render for error status code.
// Example: 400, 403, 404, 500.
func renderHttpErrorCode(ctx iris.Context, code int) {
	ctx.StatusCode(http.StatusOK)
	_, _ = ctx.JSON(app.With.ErrorCode(fmt.Errorf("HTTP %d %s", code, http.StatusText(code)), app.ErrCode(code)))
}
