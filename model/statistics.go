package model

type Statistics struct {
	JudgeStatistics JudgeStatistics `json:"judge_statistics"`
	JudgeSystemInfo JudgeSystemInfo `json:"judge_system_info"`
}
