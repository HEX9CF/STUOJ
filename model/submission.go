package model

import "time"

// 提交状态
type SubmitStatus uint64

const (
	SubmitStatusPending SubmitStatus = 0
	SubmitStatusAC      SubmitStatus = 3
	SubmitStatusWA      SubmitStatus = 4
	SubmitStatusTLE     SubmitStatus = 5
	SubmitStatusCE      SubmitStatus = 6
)

func (s SubmitStatus) String() string {
	switch s {
	case SubmitStatusPending:
		return "Pending"
	case SubmitStatusAC:
		return "AC"
	case SubmitStatusWA:
		return "WA"
	case SubmitStatusTLE:
		return "TLE"
	case SubmitStatusCE:
		return "CE"
	default:
		return "Unknown"
	}
}

// 提交记录
type Submission struct {
	Id         uint64       `json:"id"`
	UserId     uint64       `json:"user_id"`
	ProblemId  uint64       `json:"problem_id"`
	Status     SubmitStatus `json:"status"`
	Score      uint64       `json:"score"`
	LanguageId uint64       `json:"language_id"`
	Length     uint64       `json:"length"`
	Memory     uint64       `json:"memory"`
	Time       float64      `json:"time"`
	SourceCode string       `json:"source_code"`
	CreateTime time.Time    `json:"submit_time"`
	UpdateTime time.Time    `json:"update_time"`
}
