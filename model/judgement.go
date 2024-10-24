package model

// 评测状态
type JudgementStatus uint8

// 单个评测点结果
type Judgement struct {
	Id           uint64          `json:"id"`
	SubmissionID uint64          `json:"submission_id"`
	Status       JudgementStatus `json:"status"`
	Time         float64         `json:"time"`
	Memory       uint64          `json:"memory"`
}
