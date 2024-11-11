package model

import "time"

// 提交状态
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
	Submission Submission  `json:"submission,omitempty"`
	Judgements []Judgement `json:"judgements,omitempty"`
}
