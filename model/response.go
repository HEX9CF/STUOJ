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
	Code ResponseCode `json:"code"` // 0 失败，1 成功，2 重新请求
	Msg  string       `json:"msg"`
	Data interface{}  `json:"data"`
}
