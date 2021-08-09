package data

import (
	"stress-testing/internal/biz"
	"sync"
)

// 接收请求结果并处理
func ReceivedStressResult(request *biz.StressRequest,ch <-chan *biz.StressResult,wg *sync.WaitGroup)  {
	
}