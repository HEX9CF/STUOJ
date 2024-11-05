package model

import "time"

// 状态：0 作废，1 公开，2 出题中，3 调试中
type ProblemStatus uint8

const (
	ProblemStatusInvalid   ProblemStatus = 0
	ProblemStatusPublic    ProblemStatus = 1
	ProblemStatusEditing   ProblemStatus = 2
	ProblemStatusDebugging ProblemStatus = 3
)

func (s ProblemStatus) String() string {
	switch s {
	case ProblemStatusInvalid:
		return "作废"
	case ProblemStatusPublic:
		return "公开"
	case ProblemStatusEditing:
		return "出题中"
	case ProblemStatusDebugging:
		return "调试中"
	default:
		return "未知状态"
	}
}

// 难度： 0 暂无评定，1 普及−，2 普及/提高−，3 普及+/提高，4 提高+/省选− ，5 省选/NOI−，6 NOI/NOI+/CTSC
type ProblemDifficulty uint8

const (
	ProblemDifficultyUnknown ProblemDifficulty = 0
	ProblemDifficulty1       ProblemDifficulty = 1
	ProblemDifficulty2       ProblemDifficulty = 2
	ProblemDifficulty3       ProblemDifficulty = 3
	ProblemDifficulty4       ProblemDifficulty = 4
	ProblemDifficulty5       ProblemDifficulty = 5
	ProblemDifficulty6       ProblemDifficulty = 6
)

func (d ProblemDifficulty) String() string {
	switch d {
	case ProblemDifficultyUnknown:
		return "暂无评定"
	case ProblemDifficulty1:
		return "普及−"
	case ProblemDifficulty2:
		return "普及/提高−"
	case ProblemDifficulty3:
		return "普及+/提高"
	case ProblemDifficulty4:
		return "提高+/省选−"
	case ProblemDifficulty5:
		return "省选/NOI−"
	case ProblemDifficulty6:
		return "NOI/NOI+/CTSC"
	default:
		return "暂无评定"
	}
}

// 题目
type Problem struct {
	Id           uint64            `gorm:"primaryKey;autoIncrement;comment:题目ID" json:"id,omitempty"`
	Title        string            `gorm:"type:text;not null;comment:标题" json:"title,omitempty"`
	Source       string            `gorm:"type:text;not null;comment:题目来源" json:"source,omitempty"`
	Difficulty   ProblemDifficulty `gorm:"not null;default:0;comment:难度" json:"difficulty,omitempty"`
	TimeLimit    float64           `gorm:"not null;default:1;comment:时间限制（s）" json:"time_limit,omitempty"`
	MemoryLimit  uint64            `gorm:"not null;default:131072;comment:内存限制（kb）" json:"memory_limit,omitempty"`
	Description  string            `gorm:"type:longtext;not null;comment:题面" json:"description,omitempty"`
	Input        string            `gorm:"type:longtext;not null;comment:输入说明" json:"input,omitempty"`
	Output       string            `gorm:"type:longtext;not null;comment:输出说明" json:"output,omitempty"`
	SampleInput  string            `gorm:"type:longtext;not null;comment:输入样例" json:"sample_input,omitempty"`
	SampleOutput string            `gorm:"type:longtext;not null;comment:输出样例" json:"sample_output,omitempty"`
	Hint         string            `gorm:"type:longtext;not null;comment:提示" json:"hint,omitempty"`
	Status       ProblemStatus     `gorm:"not null;default:1;comment:状态" json:"status,omitempty"`
	ProblemTag   []*Tag            `gorm:"many2many:tbl_problem_tag;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;association_jointable_foreignkey:tag_id;jointable_foreignkey:problem_id" json:"problem_tag,omitempty"`
	CreateTime   time.Time         `gorm:"not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"create_time,omitempty"`
	UpdateTime   time.Time         `gorm:"not null;default:CURRENT_TIMESTAMP;comment:更新时间" json:"update_time,omitempty"`
}

func (Problem) TableName() string {
	return "tbl_problem"
}
