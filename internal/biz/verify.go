package biz

import (
	"net/http"
	"sync"
)

const (
	// 请求成功
	HTTPOK = 200
	// 请求失败
	HTTPERR = 500
	//数据解析错误
	PARSEERR = 510
)

// 校验器
type Verify interface {
	GetCode() int
	GetResult() bool
}

// http 校验器类型
type VerifyHttp func(request *StressRequest, response *http.Response) (code int, isSucceed bool)

var (
	// http校验函数 集合
	VerifyMapHttp = make(map[string]VerifyHttp)
	// map 不是安全的
	VerifyMapHttpMutex sync.Mutex
)

//支持的协议
var (
	// http协议
	FormTypeHTTP = "http"
)
