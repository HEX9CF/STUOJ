package model

import (
	"STUOJ/internal/conf"
)

type Statistics struct {
	JudgeStatistics JudgeStatistics `json:"judge_statistics"`
	JudgeSystemInfo JudgeSystemInfo `json:"judge_system_info"`
}

type Configuration struct {
	System conf.Config     `json:"system"`
	Judge  JudgeConfigInfo `json:"judge"`
}
