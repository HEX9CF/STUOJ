package model

// 标签
type Tag struct {
	Id   uint64 `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// 题目标签关系
type ProblemTag struct {
	Id        uint64 `json:"id,omitempty"`
	ProblemId uint64 `json:"problem_id,omitempty"`
	TagId     uint64 `json:"tag_id,omitempty"`
}
