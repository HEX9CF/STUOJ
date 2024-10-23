package model

import "time"

type Status uint8

// 状态：0 作废，1 公开，2 出题中，3 调试中
const (
	StatusInvalid   Status = 0
	StatusPublic    Status = 1
	StatusEditing   Status = 2
	StatusDebugging Status = 3
)

func (s Status) String() string {
	switch s {
	case StatusInvalid:
		return "invalid"
	case StatusPublic:
		return "public"
	case StatusEditing:
		return "editing"
	case StatusDebugging:
		return "debugging"
	default:
		return "unknown"
	}
}

// 题目
type Problem struct {
	Id           uint64    `json:"id"`
	Title        string    `json:"title"`
	Source       string    `json:"source"`
	Difficulty   uint64    `json:"difficulty"`
	TimeLimit    uint64    `json:"time_limit"`
	MemoryLimit  uint64    `json:"memory_limit"`
	Description  string    `json:"description"`
	Input        string    `json:"input"`
	Output       string    `json:"output"`
	SampleInput  string    `json:"sample_input"`
	SampleOutput string    `json:"sample_output"`
	Hint         string    `json:"hint"`
	Status       Status    `json:"status"`
	CreateTime   time.Time `json:"create_time"`
	UpdateTime   time.Time `json:"update_time"`
}
