package model

import "time"

// 题解
type Solution struct {
	Id         uint64    `gorm:"primaryKey;autoIncrement;comment:题解ID" json:"id,omitempty"`
	LanguageId uint64    `gorm:"not null;default:0;comment:语言ID" json:"language_id,omitempty"`
	ProblemId  uint64    `gorm:"not null;default:0;comment:题目ID" json:"problem_id,omitempty"`
	SourceCode string    `gorm:"type:longtext;not null;comment:源代码" json:"source_code,omitempty"`
	CreateTime time.Time `gorm:"not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"create_time,omitempty"`
	UpdateTime time.Time `gorm:"not null;default:CURRENT_TIMESTAMP;comment:更新时间" json:"update_time,omitempty"`
}

func (Solution) TableName() string {
	return "tbl_solution"
}
