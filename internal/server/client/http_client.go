package client

import (
	"crypto/tls"
	"net/http"
	"stress-testing/internal/biz"
	"stress-testing/internal/data"
	"time"
)

//发送http 请求
func Request(sr *biz.StressRequest) (resp *http.Response, requestTime uint64, err error) {


	request, err := http.NewRequest(sr.Method, sr.URL, sr.GetBody())
	if err != nil {
		return
	}
	// 设置host
	if _, ok := sr.Headers["host"]; !ok {
		request.Host = sr.Headers["Host"]
	}

	// 默认utf-8 字符集
	if _, ok := sr.Headers["Content-Type"]; !ok {
		if sr.Headers == nil {
			sr.Headers = make(map[string]string)
		}
		sr.Headers["Content-Type"] = "application/x-www-form-urlencoded; charset=utf-8"
	}

	// 设置请求header
	for key, val := range sr.Headers {
		request.Header.Set(key, val)
	}
	tr := &http.Transport{}
	var client *http.Client

	if sr.HTTP2 {
		//todo
	} else {
		// 跳过证书验证
		tr = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}

	}

	client = &http.Client{
		Transport: tr,
		Timeout:   sr.TimeOut,
	}
	startTime := time.Now()
	resp, err = client.Do(request)
	// 计算请求所消耗的时间
	requestTime = uint64(time.Since(startTime))
	data.RequestTimeList =append(data.RequestTimeList,requestTime)
	return
}
