package entity

import (
	"time"
)

// 提交信息
type Submission struct {
	Id         uint64      `gorm:"primaryKey;autoIncrement;comment:提交记录ID" json:"id,omitempty"`
	UserId     uint64      `gorm:"not null;default:0;comment:用户ID" json:"user_id,omitempty"`
	ProblemId  uint64      `gorm:"not null;default:0;comment:题目ID" json:"problem_id,omitempty"`
	Status     JudgeStatus `gorm:"not null;default:0;comment:状态" json:"status,omitempty"`
	Score      uint64      `gorm:"not null;default:0;comment:分数" json:"score,omitempty"`
	LanguageId uint64      `gorm:"not null;default:0;comment:语言ID" json:"language_id,omitempty"`
	Length     uint64      `gorm:"not null;default:0;comment:源代码长度" json:"length,omitempty"`
	Memory     uint64      `gorm:"not null;default:0;comment:内存（kb）" json:"memory,omitempty"`
	Time       float64     `gorm:"not null;default:0;comment:运行耗时（s）" json:"time,omitempty"`
	SourceCode string      `gorm:"type:longtext;not null;comment:源代码" json:"source_code,omitempty"`
	CreateTime time.Time   `gorm:"not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"create_time,omitempty"`
	UpdateTime time.Time   `gorm:"not null;default:CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP;comment:更新时间" json:"update_time,omitempty"`
}

func (Submission) TableName() string {
	return "tbl_submission"
}
