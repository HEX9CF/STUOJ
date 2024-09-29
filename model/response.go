package model

type Response struct {
	Code int         `json:"code"` // 0 失败，1 成功
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}
