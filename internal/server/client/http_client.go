package client

import (
	"fmt"
	"net/http"
	"stress-testing/internal/biz"
	"time"
)

func Request(sr *biz.StressRequest) (resp *http.Response, requestTime uint64, err error) {
	// todo
	startTime := time.Now()
	resp, err = http.Get("http://www.baidu.com")
	requestTime = uint64(time.Since(startTime))

	fmt.Println("消耗时间",float64(requestTime)/1e9)
	return
}
