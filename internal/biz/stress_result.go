package biz

import "fmt"

// 压测请求结果
type StressResult struct {
	ID            string // 消息id
	ChanID        uint64
	Time          uint64 // 请求时间纳秒
	IsSuccessed   bool   // 是否请求成功
	ErrCode       int    // 错误码
	ReceivedBytes int64  // 接收的字节数
}

//设置消息id 并发 + 第几次请求
func (r *StressResult) SetID(chanID, number uint64) {
	id := fmt.Sprintf("%d_%d", chanID, number)
	r.ID = id
	r.ChanID = chanID

}
