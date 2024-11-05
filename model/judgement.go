package model

// 评测点结果
type Judgement struct {
	ID            uint64  `gorm:"primaryKey;autoIncrement;comment:评测点ID" json:"id,omitempty"`
	SubmissionID  uint64  `gorm:"not null;default:0;comment:提交记录ID" json:"submission_id,omitempty"`
	TestcaseID    uint64  `gorm:"not null;default:0;comment:评测点ID" json:"testcase_id,omitempty"`
	Time          float64 `gorm:"not null;default:0;comment:运行耗时（s）" json:"time,omitempty"`
	Memory        uint64  `gorm:"not null;default:0;comment:内存（kb）" json:"memory,omitempty"`
	Stdout        string  `gorm:"type:longtext;not null;comment:标准输出" json:"stdout,omitempty"`
	Stderr        string  `gorm:"type:longtext;not null;comment:标准错误输出" json:"stderr,omitempty"`
	CompileOutput string  `gorm:"type:longtext;not null;comment:编译输出" json:"compile_output,omitempty"`
	Message       string  `gorm:"type:longtext;not null;comment:信息" json:"message,omitempty"`
	Status        uint64  `gorm:"not null;default:0;comment:状态" json:"status,omitempty"`
}

func (Judgement) TableName() string {
	return "tbl_judgement"
}
