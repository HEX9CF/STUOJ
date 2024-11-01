package model

import "time"

// 操作：0 未知，1 添加，2 修改，3 删除
type Operation uint8

const (
	OperationUnknown Operation = 0
	OperationAdd     Operation = 1
	OperationUpdate  Operation = 2
	OperationDelete  Operation = 3
)

func (o Operation) String() string {
	switch o {
	case OperationUnknown:
		return "未知"
	case OperationAdd:
		return "添加"
	case OperationUpdate:
		return "修改"
	case OperationDelete:
		return "删除"
	default:
		return "未知"
	}
}

// 题目历史记录
type ProblemHistory struct {
	Id           uint64            `json:"id,omitempty"`
	UserId       uint64            `json:"user_id,omitempty"`
	ProblemId    uint64            `json:"problem_id,omitempty"`
	Title        string            `json:"title,omitempty"`
	Source       string            `json:"source,omitempty"`
	Difficulty   ProblemDifficulty `json:"difficulty,omitempty"`
	TimeLimit    float64           `json:"time_limit,omitempty"`
	MemoryLimit  uint64            `json:"memory_limit,omitempty"`
	Description  string            `json:"description,omitempty"`
	Input        string            `json:"input,omitempty"`
	Output       string            `json:"output,omitempty"`
	SampleInput  string            `json:"sample_input,omitempty"`
	SampleOutput string            `json:"sample_output,omitempty"`
	Hint         string            `json:"hint,omitempty"`
	Operation    Operation         `json:"operation,omitempty"`
	CreateTime   time.Time         `json:"create_time,omitempty"`
}
