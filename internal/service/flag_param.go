package service

import (
	"flag"
	"fmt"
	"strings"
)

// 自定义header 参数
type Header []string

// 将字符串切片转成字符串
func (h *Header) String() string {
	return fmt.Sprint(*h)
}

// 设置header 头信息
func (h *Header) Set(s string) error {
	*h = append(*h, s)
	return nil
}

// 定义命令行flag参数，便于解析
var (
	ConcurrentNumber uint64  // 并发数 启动n个协程
	PerNumber        uint64  // 请求数 每个协程/并发的处理的请求数
	URL              string  // 压测的url
	Headers          Header  // 自定义头信息
	Body             = ""    // http post 方式传输数据
	MaxCon           = 1     // 单个连接的最大请求数
	Code             = 200   // 成功状态码
	Http2            = false // 是否开启http2.0
	KeepAlive        = false // 是否开启长连接
	Method           = "GET" // 请求方式
)

func init() {
	flag.Uint64Var(&ConcurrentNumber, "c", ConcurrentNumber, "并发数")
	flag.Uint64Var(&PerNumber, "n", PerNumber, "请求数")
	flag.StringVar(&URL, "u", URL, "压测url")
	flag.IntVar(&MaxCon, "m", MaxCon, "单个host最大连接数")
	flag.IntVar(&Code, "code", Code, "请求成功的状态码")
	flag.BoolVar(&Http2, "http2", Http2, "是否开http2.0")
	flag.BoolVar(&KeepAlive, "k", KeepAlive, "是否开启长连接")
	flag.Var(&Headers, "H", "自定义头信息传递给服务器 示例:-H 'Content-Type: application/json'")
	flag.StringVar(&Method, "X", Method, "请求方式:GET POST ...")
	flag.StringVar(&Body, "d", Body, "传输数据")
	flag.Parse()
}

func CheckFlagPrarmIsOk() bool {
	Method = strings.ToUpper(Method)
	if ConcurrentNumber == 0 || PerNumber == 0 || URL == "" || (!validMethod(Method)) {
		fmt.Printf("示例: go run main.go -c 1 -n 1 -u https://www.baidu.com/ \n")
		fmt.Printf("压测地址或curl路径必填 \n")
		fmt.Printf("当前请求参数: -c %d -n %d  -u %s \n", ConcurrentNumber, PerNumber, URL)
		flag.Usage()
		return false
	}
	return true
}

// 校验合法的method
func validMethod(method string) bool {
	/*
	     Method         = "OPTIONS"                ; Section 9.2
	                    | "GET"                    ; Section 9.3
	                    | "HEAD"                   ; Section 9.4
	                    | "POST"                   ; Section 9.5
	                    | "PUT"                    ; Section 9.6
	                    | "DELETE"                 ; Section 9.7
	                    | "TRACE"                  ; Section 9.8
	                    | "CONNECT"                ; Section 9.9
	                    | extension-method
	   extension-method = token
	     token          = 1*<any CHAR except CTLs or separators>
	*/

	if len(method) < 0 {

		return false
	}

	m := []string{"OPTIONS", "GET", "POST", "PUT", "DELETE", "TRACE", "CONNECT"}

	for _, v := range m {
		if v == method {
			return true
		}
	}
	return false
}
