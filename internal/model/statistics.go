package model

import (
	"time"
)

// 按日期统计
type CountByDate struct {
	Date  time.Time `json:"date"`
	Count int64     `json:"count"`
}

// 按日期统计映射
type MapCountByDate map[string]int64

const LayoutCountByDate = "2006-01-02"

// Judge0统计信息
type Judge0Statistics struct {
	LanguageCount   int64           `json:"language_count,omitempty"`
	JudgeStatistics JudgeStatistics `json:"judge_statistics"`
	JudgeSystemInfo JudgeSystemInfo `json:"judge_system_info"`
}

// 用户统计信息
type UserStatistics struct {
	UserCount           int64          `json:"user_count,omitempty"`
	RegisterCountByDate MapCountByDate `json:"register_count_by_date,omitempty"`
}

// 提交记录统计信息
type RecordStatistics struct {
	SubmissionCount   int64          `json:"submission_count,omitempty"`
	JudgementCount    int64          `json:"judgement_count,omitempty"`
	SubmitCountByDate MapCountByDate `json:"submit_count_by_date,omitempty"`
}

// 题目统计信息
type ProblemStatistics struct {
	ProblemCount   int64          `json:"problem_count,omitempty"`
	TagCount       int64          `json:"tag_count,omitempty"`
	TestcaseCount  int64          `json:"testcase_count,omitempty"`
	SolutionCount  int64          `json:"solution_count,omitempty"`
	AddCountByDate MapCountByDate `json:"add_count_by_date,omitempty"`
}
