package model

import "time"

// 状态：0 作废，1 公开，2 出题中，3 调试中
type Status uint8

const (
	StatusInvalid   Status = 0
	StatusPublic    Status = 1
	StatusEditing   Status = 2
	StatusDebugging Status = 3
)

func (s Status) String() string {
	switch s {
	case StatusInvalid:
		return "作废"
	case StatusPublic:
		return "公开"
	case StatusEditing:
		return "出题中"
	case StatusDebugging:
		return "调试中"
	default:
		return "未知状态"
	}
}

// 难度： 0 暂无评定，1 普及−，2 普及/提高−，3 普及+/提高，4 提高+/省选− ，5 省选/NOI−，6 NOI/NOI+/CTSC
type Difficulty uint8

const (
	DifficultyUnknown Difficulty = 0
	Difficulty1       Difficulty = 1
	Difficulty2       Difficulty = 2
	Difficulty3       Difficulty = 3
	Difficulty4       Difficulty = 4
	Difficulty5       Difficulty = 5
	Difficulty6       Difficulty = 6
)

func (d Difficulty) String() string {
	switch d {
	case DifficultyUnknown:
		return "暂无评定"
	case Difficulty1:
		return "普及−"
	case Difficulty2:
		return "普及/提高−"
	case Difficulty3:
		return "普及+/提高"
	case Difficulty4:
		return "提高+/省选−"
	case Difficulty5:
		return "省选/NOI−"
	case Difficulty6:
		return "NOI/NOI+/CTSC"
	default:
		return "暂无评定"
	}
}

// 题目
type Problem struct {
	Id           uint64     `json:"id"`
	Title        string     `json:"title"`
	Source       string     `json:"source"`
	Difficulty   Difficulty `json:"difficulty"`
	TimeLimit    float64    `json:"time_limit"`
	MemoryLimit  float64    `json:"memory_limit"`
	Description  string     `json:"description"`
	Input        string     `json:"input"`
	Output       string     `json:"output"`
	SampleInput  string     `json:"sample_input"`
	SampleOutput string     `json:"sample_output"`
	Hint         string     `json:"hint"`
	Status       Status     `json:"status"`
	CreateTime   time.Time  `json:"create_time"`
	UpdateTime   time.Time  `json:"update_time"`
}
