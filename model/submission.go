package model

import "time"

// 提交状态
type SubmissionStatus uint64

const (
	SubmissionStatusPending SubmissionStatus = 0
	SubmissionStatusAC      SubmissionStatus = 3
	SubmissionStatusWA      SubmissionStatus = 4
	SubmissionStatusTLE     SubmissionStatus = 5
	SubmissionStatusCE      SubmissionStatus = 6
)

func (s SubmissionStatus) String() string {
	switch s {
	case SubmissionStatusPending:
		return "Pending"
	case SubmissionStatusAC:
		return "AC"
	case SubmissionStatusWA:
		return "WA"
	case SubmissionStatusTLE:
		return "TLE"
	case SubmissionStatusCE:
		return "CE"
	default:
		return "Unknown"
	}
}

// 提交记录
type Submission struct {
	Id         uint64           `json:"id"`
	UserId     uint64           `json:"user_id"`
	ProblemId  uint64           `json:"problem_id"`
	Status     SubmissionStatus `json:"status"`
	Score      uint64           `json:"score"`
	SubmitTime time.Time        `json:"submit_time"`
	LanguageId uint64           `json:"language_id"`
	Length     uint64           `json:"length"`
	Memory     uint64           `json:"memory"`
	Time       float64          `json:"time"`
	SourceCode string           `json:"source_code"`
}
