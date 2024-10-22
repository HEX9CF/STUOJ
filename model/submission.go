package model

import "time"

// 提交记录
type Submission struct {
	Id         uint64    `json:"id"`
	UserId     uint64    `json:"user_id"`
	ProblemId  uint64    `json:"problem_id"`
	Status     uint64    `json:"status"`
	Score      uint64    `json:"score"`
	SubmitTime time.Time `json:"submit_time"`
	LanguageId uint64    `json:"language_id"`
	Length     uint64    `json:"length"`
	Memory     uint64    `json:"memory"`
	Time       uint64    `json:"time"`
	SourceCode string    `json:"source_code"`
}
