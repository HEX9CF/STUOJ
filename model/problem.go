package model

import "time"

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
	CreateTime   time.Time `json:"create_time"`
	UpdateTime   time.Time `json:"update_time"`
}
