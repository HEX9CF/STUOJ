package model

import "time"

// 按日期统计
type CountByDate struct {
	Date  time.Time `json:"date"`
	Count uint64    `json:"count"`
}

// 按日期统计映射
type MapCountByDate map[string]uint64

// Judge0统计信息
type Judge0Statistics struct {
	JudgeStatistics JudgeStatistics `json:"judge_statistics"`
	JudgeSystemInfo JudgeSystemInfo `json:"judge_system_info"`
}

// 用户统计信息
type UserStatistics struct {
	RegisterCount MapCountByDate `json:"register_count"`
}
