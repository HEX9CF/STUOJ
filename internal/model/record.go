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

// 评测点结果
type Judgement struct {
	Id            uint64       `gorm:"primaryKey;autoIncrement;comment:评测点ID" json:"id,omitempty"`
	SubmissionId  uint64       `gorm:"not null;default:0;comment:提交记录ID" json:"submission_id,omitempty"`
	TestcaseId    uint64       `gorm:"not null;default:0;comment:评测点ID" json:"testcase_id,omitempty"`
	Time          float64      `gorm:"not null;default:0;comment:运行耗时（s）" json:"time,omitempty"`
	Memory        uint64       `gorm:"not null;default:0;comment:内存（kb）" json:"memory,omitempty"`
	Stdout        string       `gorm:"type:longtext;not null;comment:标准输出" json:"stdout,omitempty"`
	Stderr        string       `gorm:"type:longtext;not null;comment:标准错误输出" json:"stderr,omitempty"`
	CompileOutput string       `gorm:"type:longtext;not null;comment:编译输出" json:"compile_output,omitempty"`
	Message       string       `gorm:"type:longtext;not null;comment:信息" json:"message,omitempty"`
	Status        SubmitStatus `gorm:"not null;default:0;comment:状态" json:"status,omitempty"`
}

func (Judgement) TableName() string {
	return "tbl_judgement"
}

// 编程语言
type Language struct {
	Id   uint64 `gorm:"primaryKey;autoIncrement;comment:语言ID" json:"id,omitempty"`
	Name string `gorm:"type:varchar(255);not null;comment:语言名" json:"name,omitempty"`
}

func (Language) TableName() string {
	return "tbl_language"
}

// 提交记录（提交信息+评测结果）
type Record struct {
	Submission Submission  `json:"record,omitempty"`
	Judgements []Judgement `json:"judgements,omitempty"`
}
