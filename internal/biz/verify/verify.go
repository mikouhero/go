package verify

import (
	"net/http"
	"stress-testing/internal/biz"
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
type VerifyHttp func( request *biz.StressRequest,response *http.Response) (code int, isSucceed bool)

var (
	// http校验函数 集合
	VerifyMapHttp = make(map[string]VerifyHttp)
)