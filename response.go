package timeout

type Response struct {
	Code       int64       `json:"code"`
	Success    bool        `json:"success"`
	Message    string      `json:"msg"`
	Data       interface{} `json:"data"`
	Version    string      `json:"version"`
	ServerTime int64       `json:"server_time"`
	Timestamp  string      `json:"timestamp"`
	ExecTime   int64       `json:"exec_time"`
}

