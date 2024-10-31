package model

import "STUOJ/conf"

type Configuration struct {
	System conf.Config     `json:"system"`
	Judge  JudgeConfigInfo `json:"judge"`
}
