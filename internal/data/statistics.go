package data

import (
	"fmt"
	"sort"
	"stress-testing/internal/biz"
	"strings"
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
	speed            int64
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
				// todo why?
				endTime := uint64(time.Now().UnixNano())
				handleTotalTime = endTime - startTime
				calculate()
				//todo 输出结果
			case <-stop:

				// 处理完成
				return

			}
		}
	}()

	for stressResult := range ch {

		//处理数据
		requestTotalTime += stressResult.Time

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
	handleTotalTime = endTime - startTime
	calculate()

	//所有请求响应时间排序 ，底层包需要实现 Len ,Swap ，Less方法
	a := Uint64List{}
	a = RequestTimeList
	// 排序好的数组
	sort.Sort(a)

	fmt.Println("\n\n")
	fmt.Println("*************************  结果 stat  ****************************")
	fmt.Println("请求总数（并发数*请求数 ）:", successNum+failedNum, "总请求时间:",
		fmt.Sprintf("%.3f", float64(handleTotalTime)/1e9),
		"秒", "successNum:", successNum, "failedNum:", failedNum)

	fmt.Println("top80:", fmt.Sprintf("%.3f ms", float64(a[int(float64(len(a))*0.80)]/1e6)))
	fmt.Println("top90:", fmt.Sprintf("%.3f ms", float64(a[int(float64(len(a))*0.90)]/1e6)))
	fmt.Println("top95:", fmt.Sprintf("%.3f ms", float64(a[int(float64(len(a))*0.95)]/1e6)))
	fmt.Println("top99:", fmt.Sprintf("%.3f ms", float64(a[int(float64(len(a))*0.99)]/1e6)))

	fmt.Println("*************************  结果 end  ****************************")
}

var (
	qps              float64 // 平均每秒请求数
	averageTime      float64 // 平均请求时间
	maxTimeFloat     float64
	minTimeFloat     float64
	requestTimeFloat float64
	RequestTimeList  []uint64 //所有请求响应时间
)

func calculate() {
	//  每个协程成功数/总耗时(发送数据请求的总时间) (每秒)  每秒的响应请求数
	qps = float64(successNum*1e9) / float64(handleTotalTime)

	// 平均时长  成功个数/请求总耗时  毫秒
	if successNum == 0 {
		averageTime = 0
	} else {
		averageTime = float64(requestTotalTime) /float64(successNum*1e6)
	}
	// 最大请求时间（毫秒）
	maxTimeFloat = float64(maxTime) / 1e6
	// 最小请求时间（毫秒）
	minTimeFloat = float64(minTime) / 1e6
	// 所有请求总耗时（秒）
	requestTimeFloat = float64(requestTotalTime) / 1e9

	// 接收字节
	if requestTimeFloat > 0 {
		speed = int64(float64(receivedBytes) / requestTimeFloat)
	}
	s := fmt.Sprintf("%4ds│%7d│%7d│%7d│%8.2f│%8.2f│%8.2f│%8.2f│%8d│%8d│%v",
		handleTotalTime/1e9,
		chanIDNum,
		successNum,
		failedNum,
		qps,
		maxTimeFloat,
		minTimeFloat,
		averageTime,
		receivedBytes,
		speed,
		printMap(errCode),
	)
	fmt.Println(s)



}

// 将map 转为字符串
func printMap(errCode map[int]int) (mapStr string) {
	var (
		mapArr []string
	)
	for key, value := range errCode {
		mapArr = append(mapArr, fmt.Sprintf("%d:%d", key, value))
	}
	sort.Strings(mapArr)
	mapStr = strings.Join(mapArr, ";")
	return
}
