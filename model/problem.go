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
	Id           uint64            `json:"id"`
	Title        string            `json:"title"`
	Source       string            `json:"source"`
	Difficulty   ProblemDifficulty `json:"difficulty"`
	TimeLimit    float64           `json:"time_limit"`
	MemoryLimit  uint64            `json:"memory_limit"`
	Description  string            `json:"description"`
	Input        string            `json:"input"`
	Output       string            `json:"output"`
	SampleInput  string            `json:"sample_input"`
	SampleOutput string            `json:"sample_output"`
	Hint         string            `json:"hint"`
	Status       ProblemStatus     `json:"status"`
	CreateTime   time.Time         `json:"create_time"`
	UpdateTime   time.Time         `json:"update_time"`
}
