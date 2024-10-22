package model

// 单个评测点结果
type Judgement struct {
	Id           uint64 `json:"id"`
	SubmissionID uint64 `json:"submission_id"`
	Status       uint64 `json:"status"`
	Time         uint64 `json:"time"`
	Memory       uint64 `json:"memory"`
}
