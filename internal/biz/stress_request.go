package biz

import (
	"fmt"
	"io"
	"strings"
	"time"
)

// 压测请求的数据结构体
type StressRequest struct {
	URL              string            // 压测请求的url
	Method           string            // 请求的方式 GET/POST/PUT...
	Headers          map[string]string // 请求的头信息
	Body             string            // 请求的body体
	TimeOut          time.Duration     // 请求超时时间
	Debug            bool              // 调试模式
	MaxCon           int               //每个连接的请求数
	HTTP2            bool              // 是否使用http2.0
	KeepAlive        bool              // 是否开启长链接
	Code             int               // 验证状态码
	ConcurrentNumber uint64            // 并发数 启动n个协程
	PerNumber        uint64            // 请求数 每个协程/并发的处理的请求数
}

var (
	headers = make(map[string]string)
	body    string
)

func NewRequest(url string, code int, timeout time.Duration, debug bool, path string, reqHeaders []string,
	reqBody string, maxCon int, http2 bool, keepalive bool, mehtod string,perNumber,concurrentNumber uint64) (sr *StressRequest, err error) {

	// fixme  需要优化   请求判断
	if reqBody != "" {
		body = reqBody
	}
	for _, value := range reqHeaders {
		getHeaderValue(value, headers)
	}

	//fixme  兼容多种协议
	// 验证url  是否符合
	if !(strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https")) {
		err = fmt.Errorf("url:%s 不合法,必须是完整的http..", url)
		return
	}

	// fixme  读取配置
	if timeout == 0 {
		timeout = 15 * time.Second
	}
	sr = &StressRequest{
		URL:       url,
		Method:    mehtod,
		Headers:   headers,
		Body:      reqBody,
		TimeOut:   timeout,
		Debug:     debug,
		MaxCon:    maxCon,
		HTTP2:     http2,
		KeepAlive: keepalive,
		Code:      code,
		ConcurrentNumber :concurrentNumber,
		PerNumber        :perNumber,
	}
	return
}

// 解析headers 头信息
func getHeaderValue(v string, headers map[string]string) {
	index := strings.Index(v, ":")
	if index <= 0 {
		return
	}
	nextIndex := index + 1
	if len(v) >= nextIndex {
		// 值信息
		value := strings.TrimSuffix(v[nextIndex:], " ")
		fmt.Println(value)
		// 键
		key := v[:index]

		if _, ok := headers[key]; ok {
			//追加
			headers[key] = fmt.Sprintf("%s; %s", headers[key], value)
		} else {
			headers[key] = value

		}

	}

}

// 获取请求数据
func (r *StressRequest) GetBody() (body io.Reader) {
	return strings.NewReader(r.Body)
}

func (r *StressRequest) GetDebug() bool {
	return r.Debug
}
