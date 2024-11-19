package model

import "time"

// Judge0统计信息
type Judge0Statistics struct {
	LanguageCount   int64           `json:"language_count,omitempty"`
	JudgeStatistics JudgeStatistics `json:"judge_statistics"`
	JudgeSystemInfo JudgeSystemInfo `json:"judge_system_info"`
}

// 用户统计信息
type UserStatistics struct {
	UserCount           int64    `json:"user_count,omitempty"`
	UserCountByRole     MapCount `json:"user_count_by_role,omitempty"`
	RegisterCountByDate MapCount `json:"register_count_by_date,omitempty"`
}

func (s *UserStatistics) FillZero(startDate time.Time, endDate time.Time) {
	s.UserCountByRole.FillZero(startDate, endDate)
	s.RegisterCountByDate.FillZero(startDate, endDate)
}

// 提交记录统计信息
type RecordStatistics struct {
	SubmissionCount           int64    `json:"submission_count,omitempty"`
	JudgementCount            int64    `json:"judgement_count,omitempty"`
	SubmissionCountByLanguage MapCount `json:"submission_count_by_language,omitempty"`
	SubmissionCountByStatus   MapCount `json:"submission_count_by_status,omitempty"`
	JudgementCountByStatus    MapCount `json:"judgement_count_by_status,omitempty"`
	SubmissionCountByDate     MapCount `json:"submission_count_by_date,omitempty"`
}

func (s *RecordStatistics) FillZero(startDate time.Time, endDate time.Time) {
	s.SubmissionCountByLanguage.FillZero(startDate, endDate)
	s.SubmissionCountByStatus.FillZero(startDate, endDate)
	s.JudgementCountByStatus.FillZero(startDate, endDate)
	s.SubmissionCountByDate.FillZero(startDate, endDate)
}

// 题目统计信息
type ProblemStatistics struct {
	ProblemCount      int64    `json:"problem_count,omitempty"`
	TestcaseCount     int64    `json:"testcase_count,omitempty"`
	SolutionCount     int64    `json:"solution_count,omitempty"`
	InsertCountByDate MapCount `json:"insert_count_by_date,omitempty"`
	UpdateCountByDate MapCount `json:"update_count_by_date,omitempty"`
	DeleteCountByDate MapCount `json:"delete_count_by_date,omitempty"`
}

func (p *ProblemStatistics) FillZero(startDate time.Time, endDate time.Time) {
	p.InsertCountByDate.FillZero(startDate, endDate)
	p.UpdateCountByDate.FillZero(startDate, endDate)
	p.DeleteCountByDate.FillZero(startDate, endDate)
}

// 标签统计信息
type TagStatistics struct {
	TagCount          int64    `json:"tag_count,omitempty"`
	ProblemCountByTag MapCount `json:"problem_count_by_tag,omitempty"`
}
