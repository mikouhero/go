package data

import (
	"fmt"
	"stress-testing/internal/biz"
	"sync"
	"time"
)

var (
	handleTotalTime  uint64   = 1 // 处理总时间
	requestTotalTime uint64       // 请求总时间
	responseTimeList []uint64     //所有请求响应时间
	maxTime          uint64       // 单个请求最长时间
	minTime          uint64       // 单个请求最小时间
	successNum       uint64       // 请求成功总数
	failedNum        uint64       // 请求失败总数
	chanIDNum        int          //并发数
	chanIDs          = make(map[uint64]bool)
	receivedBytes    int64                 // 接收总字节数
	errCode          = make(map[int]int)   // 错误码对应的错误个数
	stop             = make(chan struct{}) //结束标示
	concurrentNumber uint64
)

// 接收请求结果并处理
func ReceivedStressResult(request *biz.StressRequest, ch <-chan *biz.StressResult, wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
	}()
	concurrentNumber = request.ConcurrentNumber
	// 开始时间
	startTime := uint64(time.Now().UnixNano())
	// todo  配置多久输出
	ticker := time.NewTicker(1 * time.Second)
	go func() {
		for {
			select {
			case <-ticker.C:
				//todo 输出结果
				fmt.Println("adfadfadfadfadf")
			case <-stop:

				// 处理完成
				fmt.Println("over")
				return

			}
		}
	}()

	for stressResult := range ch {

		//处理数据
		handleTotalTime += stressResult.Time

		// 单个请求最长时间
		if maxTime < stressResult.Time {
			maxTime = stressResult.Time
		}
		//单个请求最短时间

		if minTime == 0 {
			minTime = stressResult.Time
		} else if minTime > stressResult.Time {
			minTime = stressResult.Time
		}
		// 请求成功/失败总数
		if stressResult.IsSuccessed {
			successNum += 1
		} else {
			failedNum += 1
		}

		//返回状态码收集
		if _, ok := errCode[stressResult.ErrCode]; ok {
			errCode[stressResult.ErrCode] += 1
		} else {
			errCode[stressResult.ErrCode] = 1
		}

		// 接收总字节数
		receivedBytes += stressResult.ReceivedBytes

		if _, ok := chanIDs[stressResult.ChanID]; !ok {
			chanIDs[stressResult.ChanID] = true
			chanIDNum = len(chanIDs)
		}

	}
	stop <- struct{}{}
	endTime := uint64(time.Now().UnixNano())
	requestTotalTime = endTime - startTime
	fmt.Println(requestTotalTime)
	calculate()
}

func calculateData(concurrent, handleTotleTime, requestTotleTime, maxTime, minTime, successNum, failureNum uint64,
	chanIDLen int, errCode map[int]int, receivedBytes int64) {

}

var (
	qps              float64 // 平均每秒请求数
	averageTime      float64 // 平均请求时间
	maxTimeFloat     float64
	minTimeFloat     float64
	requestTimeFloat float64
)

func calculate() {
	//  每个协程成功数/总耗时(发送数据请求的总时间) (每秒)  每秒的响应请求数
	qps = float64(successNum*1e9) / float64(handleTotalTime)

	// 平均时长 总耗时/总请求数/并发数 纳秒=>毫秒
	if successNum == 0 {
		averageTime = 0
	} else {
		averageTime = float64(handleTotalTime) / float64(successNum*1e6*concurrentNumber)
	}
	maxTimeFloat = float64(maxTime) / 1e6
	minTimeFloat = float64(minTime) / 1e6
	requestTimeFloat = float64(requestTotalTime) / 1e9

	fmt.Println(qps, maxTimeFloat, minTimeFloat, requestTimeFloat)
}
