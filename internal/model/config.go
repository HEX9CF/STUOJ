package model

import (
	"STUOJ/internal/conf"
)

// 系统设置
type Configuration struct {
	System conf.Config     `json:"system"`
	Judge  JudgeConfigInfo `json:"judge"`
}
