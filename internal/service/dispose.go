package service

import (
	"stress-testing/internal/biz"
	"stress-testing/internal/data"
	"stress-testing/internal/server/http"
	"sync"
	"time"
)

var (
	wg sync.WaitGroup // 发送数据
	wgReceiving sync.WaitGroup //数据处理
)
//处理请求数据
func Dispose(sr *biz.StressRequest)  {
	//fixme  ini 配置
	//设置接收压测结果缓存
	ch := make(chan *biz.StressResult, 1000)
	wgReceiving.Add(1)
	// 启动协程处理请求结果
	go data.ReceivedStressResult(sr,ch,&wgReceiving)

	//根据stressRequest 的concurrentNumber 启动协程

	for i:=uint64(0);i< sr.ConcurrentNumber;i++ {
		wg.Add(1)
		//todo  发送请求
		go http.Request(i,ch,sr,&wg)
	}

	// 所有的请求完成
	wg.Done();
	// 休息1s  确保所有数据接收完成
	time.Sleep(1*time.Second)
	//关闭通道
	close(ch)
	//等待数据处理
	wgReceiving.Wait();
	return
}