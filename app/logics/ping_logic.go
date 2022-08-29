package index

type (
	// PingResponse
	// 健康检查返回结果.
	PingResponse struct {
		Pid         int     `json:"pid" label:"进程ID" mock:"1001"`
		StartedTime string  `json:"started_time" label:"服务启动时间" mock:"2021-12-23 34:56:08"`
		Goroutines  int     `json:"goroutines" label:"协程数" mock:"5"`
		MemorySize  float64 `json:"memory_size" label:"内存占用量" mock:"17.89"`
	}
)
