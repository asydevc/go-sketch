// author: asydevc <asydev@163.com>
// date: 2021-08-04

package app

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

// Config
// APP配置.
//
// 包在初始化时, 读取配置文件`app.yaml`数据, 并赋值到
// 指定字段上.
//
//	// 在项目中用法:
//	println("host:", app.Config.Host) // 127.0.0.1
//	println("port:", app.Config.Port) // 8080
var Config *configuration

// 基础配置.
type configuration struct {
	Addr                string    `yaml:"-" json:"addr"`                   // 服务地址.
	Host                string    `yaml:"host" json:"host"`                // Host地址.
	Lang                string    `yaml:"lang" json:"lang"`                // 语言包.
	Name                string    `yaml:"name" json:"name"`                // 应用名称.
	Pid                 int       `yaml:"-" json:"pid"`                    // 操作系统进程ID.
	Port                int       `yaml:"port" json:"port"`                // 端口号.
	Software            string    `yaml:"-" json:"software"`               // 软件名称.
	StartTime           time.Time `yaml:"-" json:"start_time"`             // 启动时间
	Version             string    `yaml:"version" json:"version"`          // 应用版本号.
	LogHttpPayload      bool      `yaml:"log-http-payload" json:"-"`       // 日志中是否记录HTTP请求入参.
	LogHttpResponseBody bool      `yaml:"log-http-response-body" json:"-"` // 日志中是否记录HTTP返回结果.
}

// 填充默认值.
func (o *configuration) defaults() *configuration {
	o.Pid = os.Getpid()
	o.Addr = fmt.Sprintf("%s:%d", o.Host, o.Port)
	o.Software = fmt.Sprintf("%s/%s", o.Name, o.Version)
	o.StartTime = time.Now()

	if o.Lang == "" {
		o.Lang = "zh"
	}

	return o
}

// Required异常.
func (o *configuration) fatal() *configuration {
	if o.Name == "" || o.Version == "" {
		panic("application info not defined in app.yaml")
	}
	if o.Port < 1000 {
		panic("server port defined in app.yaml can not less than 1,000")
	}
	return o
}

// 初始化配置.
func (o *configuration) init() *configuration {
	return o.load().defaults().fatal()
}

// 从app.yaml中解析配置.
func (o *configuration) load() *configuration {
	for _, file := range []string{
		"tmp/app.yaml",
		"config/app.yaml",
		"../tmp/app.yaml",
		"../config/app.yaml",
	} {
		body, err := ioutil.ReadFile(file)
		if err != nil {
			continue
		}
		if err = yaml.Unmarshal(body, o); err == nil {
			break
		}
	}
	return o
}
