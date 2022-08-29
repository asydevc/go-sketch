// author: asydevc <asydev@163.com>
// date: 2021-08-04

package main

import (
	"github.com/asydevc/console/v2"
	"github.com/asydevc/console/v2/base"
	"github.com/asydevc/console/v2/i"
	"github.com/asydevc/log/v2"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"

	"github.com/asydevc/go-sketch/app"
	"github.com/asydevc/go-sketch/app/controllers"
	"github.com/asydevc/go-sketch/app/middlewares"
)

func main() {
	// 1. 启动命令.
	st := &base.Command{}
	st.Initialize()
	st.SetName("start")
	st.SetHandler(start)
	// 2. 加入Console.
	cs := console.Default()
	cs.Add(st)
	cs.Run()
}

func start(_ i.IConsole) {
	// 1. 创建IRIS服务.
	//    iris application.
	ia := iris.New()

	// 2. 注册全局前置中间件.
	//    进入控制器前顺序执行.
	ia.UseGlobal(
		middlewares.BeforeTracingMiddleware,
		middlewares.BeforePanicMiddleware,
		middlewares.BeforeControllerMiddleware,
	)

	// 3. 注册全局后置中间件.
	//    控制器完成后顺序执行.
	ia.DoneGlobal(
		middlewares.AfterControllerMiddleware,
	)

	// 4. 监听HTTP错误状态码.
	//    包括400, 403, 404, 500等.
	ia.OnAnyErrorCode(middlewares.HttpCodeMiddleware)

	// 5. 注册控制器.
	mo := iris.ExecutionOptions{Force: true}
	mr := iris.ExecutionRules{Done: mo}
	// 5.1 Controller.
	mvc.Configure(ia.Party("/"), func(m *mvc.Application) {
		m.Router.SetExecutionRules(mr)
		m.Handle(&controllers.IndexController{})
	})

	// 6. 配置IRIS启动项.
	cfg := iris.WithConfiguration(iris.Configuration{
		DisableBodyConsumptionOnUnmarshal: true,
	})

	// 7. 启动服务.
	if err := ia.Configure(cfg).Run(iris.Addr(app.Config.Addr)); err != nil {
		log.Panicf("start fatal: %v.", err)
	}
}
