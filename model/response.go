package model

type ResponseCode uint8

const (
	ResponseCodeError ResponseCode = 0
	ResponseCodeOk    ResponseCode = 1
	ResponseCodeRetry ResponseCode = 2
)

func (c ResponseCode) String() string {
	switch c {
	case ResponseCodeError:
		return "错误"
	case ResponseCodeOk:
		return "成功"
	case ResponseCodeRetry:
		return "重新请求"
	default:
		return "未知状态"
	}
}

// http响应体
type Response struct {
	Code ResponseCode `json:"code"` // 0 失败，1 成功，2 重新请求
	Msg  string       `json:"msg"`
	Data interface{}  `json:"data"`
}
