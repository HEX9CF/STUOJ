package entity

import (
	"time"
)

// 提交状态：0 Pend, 1 In Queue, 2 Proc, 3 AC, 4 WA, 5 TLE, 6 CE, 7 RE(SIGSEGV), 8 RE(SIGXFSZ), 9 RE(SIGFPE), 10 RE(SIGABRT), 11 RE(NZEC), 12 RE(Other), 13 IE, 14 EFE
type SubmitStatus uint64

const (
	SubmitStatusPend      SubmitStatus = 0
	SubmitStatusIQ        SubmitStatus = 1
	SubmitStatusProc      SubmitStatus = 2
	SubmitStatusAC        SubmitStatus = 3
	SubmitStatusWA        SubmitStatus = 4
	SubmitStatusTLE       SubmitStatus = 5
	SubmitStatusCE        SubmitStatus = 6
	SubmitStatusRESIGSEGV SubmitStatus = 7
	SubmitStatusRESIGXFSZ SubmitStatus = 8
	SubmitStatusRESIGFPE  SubmitStatus = 9
	SubmitStatusRESIGABRT SubmitStatus = 10
	SubmitStatusRENZEC    SubmitStatus = 11
	SubmitStatusREOther   SubmitStatus = 12
	SubmitStatusIE        SubmitStatus = 13
	SubmitStatusEFE       SubmitStatus = 14
)

func (s SubmitStatus) String() string {
	switch s {
	case SubmitStatusPend:
		return "Pending"
	case SubmitStatusIQ:
		return "In Queue"
	case SubmitStatusProc:
		return "Processing"
	case SubmitStatusAC:
		return "Accepted"
	case SubmitStatusWA:
		return "Wrong Answer"
	case SubmitStatusTLE:
		return "Time Limit Exceeded"
	case SubmitStatusCE:
		return "Compilation Error"
	case SubmitStatusRESIGSEGV:
		return "Runtime Error (SIGSEGV)"
	case SubmitStatusRESIGXFSZ:
		return "Runtime Error (SIGXFSZ)"
	case SubmitStatusRESIGFPE:
		return "Runtime Error (SIGFPE)"
	case SubmitStatusRESIGABRT:
		return "Runtime Error (SIGABRT)"
	case SubmitStatusRENZEC:
		return "Runtime Error (NZEC)"
	case SubmitStatusREOther:
		return "Runtime Error (Other)"
	case SubmitStatusIE:
		return "Internal Error"
	case SubmitStatusEFE:
		return "Exec Format Error"
	default:
		return "Unknown"
	}
}

// 提交信息
type Submission struct {
	Id         uint64       `gorm:"primaryKey;autoIncrement;comment:提交记录ID" json:"id,omitempty"`
	UserId     uint64       `gorm:"not null;default:0;comment:用户ID" json:"user_id,omitempty"`
	ProblemId  uint64       `gorm:"not null;default:0;comment:题目ID" json:"problem_id,omitempty"`
	Status     SubmitStatus `gorm:"not null;default:0;comment:状态" json:"status,omitempty"`
	Score      uint64       `gorm:"not null;default:0;comment:分数" json:"score,omitempty"`
	LanguageId uint64       `gorm:"not null;default:0;comment:语言ID" json:"language_id,omitempty"`
	Length     uint64       `gorm:"not null;default:0;comment:源代码长度" json:"length,omitempty"`
	Memory     uint64       `gorm:"not null;default:0;comment:内存（kb）" json:"memory,omitempty"`
	Time       float64      `gorm:"not null;default:0;comment:运行耗时（s）" json:"time,omitempty"`
	SourceCode string       `gorm:"type:longtext;not null;comment:源代码" json:"source_code,omitempty"`
	CreateTime time.Time    `gorm:"not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"create_time,omitempty"`
	UpdateTime time.Time    `gorm:"not null;default:CURRENT_TIMESTAMP;comment:更新时间" json:"update_time,omitempty"`
}

func (Submission) TableName() string {
	return "tbl_submission"
}
