// author: asydevc <asydev@163.com>
// date: 2021-05-24

package app

import (
	"fmt"

	"github.com/asydevc/log/v2"
)

// Catch
// 捕获Panic错误.
//
//	func (o *Logic) Run (ctx iris.Context) interface{} {
//
//	    defer CatchWithMessageAndCallbacks(ctx)
//
//	}
func Catch(ctx interface{}) {
	CatchWithMessageAndCallbacks(ctx, "")
}

// CatchWithCallbacks
// 捕获Panic错误并执行回调.
//
//	func (o *Logic) Run (ctx iris.Context) interface{} {
//
//	    defer CatchWithMessageAndCallbacks(ctx, func(){
//	        println("callback", 1)
//	    }, func(){
//	        println("callback", 2)
//	    })
//
//	}
func CatchWithCallbacks(ctx interface{}, callbacks ...func()) {
	CatchWithMessageAndCallbacks(ctx, "", callbacks...)
}

// CatchWithMessage
// 捕获Panic错误并指定错误内容.
//
//	func (o *Logic) Run (ctx iris.Context) interface{} {
//
//	    defer CatchWithMessage(ctx, "message")
//
//	}
func CatchWithMessage(ctx interface{}, message string) {
	CatchWithMessageAndCallbacks(ctx, message)
}

// CatchWithMessageAndCallbacks
// 捕获Panic错误并指定错误内容, 同时触发回调.
//
//	func (o *Logic) Run (ctx iris.Context) interface{} {
//
//	    defer CatchWithMessageAndCallbacks(ctx, "message", func(){
//	        println("callback", 1)
//	    }, func(){
//	        println("callback", 2)
//	    })
//
//	}
func CatchWithMessageAndCallbacks(ctx interface{}, message string, callbacks ...func()) {
	// 1. 检查异常.
	//    以下代码仅在运行中出现Panic时, 才会触发.
	if r := recover(); r != nil {
		// 读取原因.
		if message == "" {
			message = fmt.Sprintf("%v", r)
		} else {
			message = fmt.Sprintf(message+": %v", r)
		}

		// 写入日志.
		// 日志包括运行期的栈信息.
		log.Panicfc(ctx, message)
	}

	// 2. 捕获异常.
	//    回调方法中出现Panic异常时, 走入此过程.
	defer func() {
		if r := recover(); r != nil {
			log.Panicfc(ctx, fmt.Sprintf("callback panic: %v", r))
		}
	}()

	// 3. 触发回调.
	//    不论是否有Panic异常, 回调都会触发, 且是串行执行.
	for _, callback := range callbacks {
		callback()
	}
}
