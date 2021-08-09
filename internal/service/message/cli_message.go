package message

import (
	"fmt"
	"stress-testing/internal/biz"
)

type CliMessage struct {
}

func NewCliMessage() *CliMessage {
	return &CliMessage{}
}

func (cm *CliMessage) Welcome() {
	fmt.Println(`
                     .::::.
                   .::::::::.
                  :::::::::::   欢迎使用GO压测工具
              ..:::::::::::'
            '::::::::::::'
              .::::::::::
         '::::::::::::::..
              ..::::::::::::.
             ''::::::::::::::::
             ::::'':::::::::'        .:::.
            ::::'   ':::::'       .::::::::.
          .::::'      ::::     .:::::::'::::.
         .:::'       :::::  .:::::::::' ':::::.
        .::'        :::::.:::::::::'      ':::::.
       .::'         ::::::::::::::'         ''::::.
   ...:::           ::::::::::::'              ''::.
  '''' ':.          ':::::::::'                  ::::..
                     '.:::::'                    ':'''''..

`)
}
func (cm *CliMessage) ShowParam(stressRequest *biz.StressRequest) {
	result := fmt.Sprint("输入的信息如下：\n")
	result = fmt.Sprintf("%s 请求地址:%s \n", result, stressRequest.URL)
	result = fmt.Sprintf("%s 请求方式:%s \n", result, stressRequest.Method)
	result = fmt.Sprintf("%s 请求数据内容:%s \n", result, stressRequest.Body)
	result = fmt.Sprintf("%s 请求的头信息:%s \n", result, stressRequest.Headers)
	result = fmt.Sprintf("%s 请求方式:%s \n", result, stressRequest.Method)
	fmt.Println(result)
	return
}

func(cm *CliMessage) Header() {
	fmt.Printf("\n\n")
	// 打印的时长都为毫秒 总请数
	fmt.Println("─────┬───────┬───────┬───────┬────────┬────────┬────────┬────────┬────────┬────────┬────────")
	fmt.Println(" 耗时│ 并发数│ 成功数│ 失败数│   qps  │最长耗时│最短耗时│平均耗时│下载字节│字节每秒│ 状态码")
	fmt.Println("─────┼───────┼───────┼───────┼────────┼────────┼────────┼────────┼────────┼────────┼────────")
	return
}
func (cm *CliMessage) HandleData() {

}
func (cm *CliMessage) Finish() {

}
