package main

import (
	"stress-testing/internal/service"
	_ "stress-testing/internal/service"
)

func main() {


	//验证必要参数
	if ok := service.CheckFlagPrarmIsOk();!ok {
		return
	}

	//返回 request 对象

	// 分发数据 开始处理

}
