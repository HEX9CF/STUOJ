package model

type ResponseCode uint8

const (
	ResponseCodeError   ResponseCode = 0
	ResponseCodeSuccess ResponseCode = 1
)

func (c ResponseCode) String() string {
	switch c {
	case ResponseCodeError:
		return "错误"
	case ResponseCodeSuccess:
		return "成功"
	default:
		return "未知状态"
	}
}

// http响应体
type Response struct {
	Code ResponseCode `json:"code"`
	Msg  string       `json:"msg,omitempty"`
	Data interface{}  `json:"data,omitempty"`
}
