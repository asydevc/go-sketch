package controllers

import (
	"github.com/asydevc/go-sketch/app"
	index "github.com/asydevc/go-sketch/app/logics"
	"github.com/kataras/iris/v12"
	"runtime"
)

type IndexController struct {
}

// Get
// 默认页面.
//
// 避免404请求的默认路由.
//
// @Ignore()
func (*IndexController) Get(ctx iris.Context) interface{} {
	return app.With.Success()
}

// GetPing
// 健康检查.
//
// 用于Consul, K8等检测服务健康状态.
//
// @Ignore()
// @Response(app/logics/index.PingResponse)
func (o *IndexController) GetPing(ctx iris.Context) interface{} {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return app.With.Data(index.PingResponse{
		Pid:         app.Config.Pid,
		StartedTime: app.Config.StartTime.Format("2006-01-02 15:04:05"),
		Goroutines:  runtime.NumGoroutine(),
		MemorySize:  float64(m.Sys) / 1024.0 / 1024.0,
	})
}
