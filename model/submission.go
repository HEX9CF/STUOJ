package model

import "time"

// 提交状态
type SubmitStatus uint64

const (
	SubmitStatusPending         SubmitStatus = 0
	SubmitStatusInQueue         SubmitStatus = 1
	SubmitStatusProcessing      SubmitStatus = 2
	SubmitStatusAC              SubmitStatus = 3
	SubmitStatusWA              SubmitStatus = 4
	SubmitStatusTLE             SubmitStatus = 5
	SubmitStatusCE              SubmitStatus = 6
	SubmitStatusRE_SIGSEGV      SubmitStatus = 7
	SubmitStatusRE_SIGXFSZ      SubmitStatus = 8
	SubmitStatusRE_SIGFPE       SubmitStatus = 9
	SubmitStatusRE_SIGABRT      SubmitStatus = 10
	SubmitStatusRE_NZEC         SubmitStatus = 11
	SubmitStatusRE_Other        SubmitStatus = 12
	SubmitStatusInternalError   SubmitStatus = 13
	SubmitStatusExecFormatError SubmitStatus = 14
)

func (s SubmitStatus) String() string {
	switch s {
	case SubmitStatusPending:
		return "Pending"
	case SubmitStatusInQueue:
		return "In Queue"
	case SubmitStatusProcessing:
		return "Processing"
	case SubmitStatusAC:
		return "Accepted"
	case SubmitStatusWA:
		return "Wrong Answer"
	case SubmitStatusTLE:
		return "Time Limit Exceeded"
	case SubmitStatusCE:
		return "Compilation Error"
	case SubmitStatusRE_SIGSEGV:
		return "Runtime Error (SIGSEGV)"
	case SubmitStatusRE_SIGXFSZ:
		return "Runtime Error (SIGXFSZ)"
	case SubmitStatusRE_SIGFPE:
		return "Runtime Error (SIGFPE)"
	case SubmitStatusRE_SIGABRT:
		return "Runtime Error (SIGABRT)"
	case SubmitStatusRE_NZEC:
		return "Runtime Error (NZEC)"
	case SubmitStatusRE_Other:
		return "Runtime Error (Other)"
	case SubmitStatusInternalError:
		return "Internal Error"
	case SubmitStatusExecFormatError:
		return "Exec Format Error"
	default:
		return "Unknown"
	}
}

// 提交信息
type Submission struct {
	Id         uint64       `json:"id,omitempty"`
	UserId     uint64       `json:"user_id,omitempty"`
	ProblemId  uint64       `json:"problem_id,omitempty"`
	Status     SubmitStatus `json:"status,omitempty"`
	Score      uint64       `json:"score,omitempty"`
	LanguageId uint64       `json:"language_id,omitempty"`
	Length     uint64       `json:"length,omitempty"`
	Memory     uint64       `json:"memory,omitempty"`
	Time       float64      `json:"time,omitempty"`
	SourceCode string       `json:"source_code,omitempty"`
	CreateTime time.Time    `json:"submit_time,omitempty"`
	UpdateTime time.Time    `json:"update_time,omitempty"`
}
