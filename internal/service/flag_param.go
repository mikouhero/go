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
	ConcurrentNumber uint64   = 1 // 并发数 启动n个协程
	PerNumber        uint64   = 1 // 请求数 每个协程/并发的处理的请求数
	URL              string       // 压测的url
	Headers          Header       // 自定义头信息
	Body             = ""         // http post 方式传输数据
	MaxCon           = 1          // 单个连接的最大请求数
	Code             = 200        // 成功状态码
	Http2            = false      // 是否开启http2.0
	KeepAlive        = false      // 是否开启长连接
	Method           = "GET"      // 请求方式
	Verify           = "code"     // http相应状态码 json
)

func init() {
	/*
		fmt.Print("请输入要压测的url (http://baidu.com):")
		fmt.Scanln(&URL)
		fmt.Println()
		//校验
		fmt.Print("请输入要压测的并发数 (非数字默认：1):")
		fmt.Scanln(&ConcurrentNumber)
		fmt.Println()

		fmt.Print("请输入单个并发的请求数 (非数字默认：1):")
		fmt.Scanln(&PerNumber)
		fmt.Println()

		s := fmt.Sprintf("请输入请求方式：\n"+
			"%10s\n"+
			"%11s\n"+
			"%10s\n"+
			"%13s\n",
			"1: GET", "2: POST", "3: PUT", "4: DELETE")
		fmt.Println(s)
		var m int = 1
		fmt.Scanln(&m)

		fmt.Println()

		fmt.Println("请输入请求的参数(参照CURL数据格式)")
		fmt.Scanln(&Body)
		fmt.Println()

	*/

	header()
}

func header() {
	var h string = "Y"
	fmt.Print("是否需要添加头信息 (Y/N): ")

	fmt.Scanln(&h)

	if "N" == strings.ToUpper(h) {
		return
	}

	if "Y" != strings.ToUpper(h) && "N" != strings.ToUpper(h) {
		fmt.Println()
		s := fmt.Sprintf("  %25s", "输入有误, 请输入 `Y` 或者 `N`")
		fmt.Println(s)
		fmt.Println()

		header()
		return
	}

	var s string
	//todo 处理header 逻辑


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
