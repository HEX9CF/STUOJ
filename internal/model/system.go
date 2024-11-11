package model

import (
	"STUOJ/internal/conf"
)

// 统计数据
type Statistics struct {
	JudgeStatistics JudgeStatistics `json:"judge_statistics"`
	JudgeSystemInfo JudgeSystemInfo `json:"judge_system_info"`
}

// 系统设置
type Configuration struct {
	System conf.Config     `json:"system"`
	Judge  JudgeConfigInfo `json:"judge"`
}
