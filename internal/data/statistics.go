package data

import (
	"fmt"
	"stress-testing/internal/biz"
	"sync"
	"time"
)

// 接收请求结果并处理
func ReceivedStressResult(request *biz.StressRequest, ch <-chan *biz.StressResult, wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
	}()

	var stop = make(chan struct{})

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
		// 结果处理
		fmt.Println("通道内个数",len(ch))
		fmt.Println(stressResult.ID)
	//	//处理数据
	}
	stop <- struct{}{}
}
