package http

import (
	"fmt"
	"net/http"
	"stress-testing/internal/biz"
	"stress-testing/internal/biz/verify"
	"stress-testing/internal/server/client"
	"sync"
)

// http 请求处理
func Request(chanID uint64, ch chan<- *biz.StressResult, sr *biz.StressRequest, wg *sync.WaitGroup) {

	defer func() {
		wg.Done()
	}()

	// 同步发送单个协成的请求
	for i := uint64(0); i < sr.PerNumber; i++ {
		isSucceed, errCode, requestTime, contentLength := sendRequest(sr)
		result := &biz.StressResult{
			Time:          requestTime,
			IsSuccessed:   isSucceed,
			ErrCode:       errCode,
			ReceivedBytes: contentLength,
		}
		result.SetID(chanID, i)
		fmt.Println(result)
		ch <- result
	}
	return

}

// 发送请求
func sendRequest(sr *biz.StressRequest) (bool, int, uint64, int64) {

	var (
		isSuccessed   = false
		errCode       = verify.HTTPOK
		contentLength = int64(0)
		err           error
		resp          *http.Response
		requestTime   uint64
	)
	resp, requestTime, err = client.Request(sr)
	fmt.Println(resp,requestTime,err)
	if err != nil {
		errCode = verify.HTTPERR
	} else {
		contentLength = resp.ContentLength
		// todo 处理成功逻辑

	}
	return isSuccessed, errCode, requestTime, contentLength
}
