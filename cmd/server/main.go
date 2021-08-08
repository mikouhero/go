package main

import (
	"fmt"
	"stress-testing/internal/biz"
	"stress-testing/internal/service"
	_ "stress-testing/internal/service"
)

func main() {

	//验证必要参数
	if ok := service.CheckFlagPrarmIsOk(); !ok {
		return
	}

	// 获取
	sr, err := biz.NewRequest(service.URL, service.Code, 0, false, "", service.Headers, service.Body, service.MaxCon, service.Http2, service.KeepAlive,service.Method)

	if err != nil {
		fmt.Printf("参数不合法 %v \n", err)
		return
	}
	fmt.Println(sr)

	// 分发数据 开始处理

}
