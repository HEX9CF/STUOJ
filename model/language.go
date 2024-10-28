package model

// 编程语言
type Language struct {
	Id   uint64 `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
