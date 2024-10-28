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
	SourceCode string       `json:"source_code,omitempty"`
	CreateTime time.Time    `json:"submit_time"`
	UpdateTime time.Time    `json:"update_time"`
}
