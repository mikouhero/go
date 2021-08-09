package message

// message 接口
type  Message interface {
	Welcome()
	ShowParam(...interface{})
	HandleData()
	Finish()
}