package model

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

// 提交记录统计信息
type RecordStatistics struct {
	SubmissionCount   int64    `json:"submission_count,omitempty"`
	JudgementCount    int64    `json:"judgement_count,omitempty"`
	SubmitCountByDate MapCount `json:"submit_count_by_date,omitempty"`
}

// 题目统计信息
type ProblemStatistics struct {
	ProblemCount      int64    `json:"problem_count,omitempty"`
	TagCount          int64    `json:"tag_count,omitempty"`
	TestcaseCount     int64    `json:"testcase_count,omitempty"`
	SolutionCount     int64    `json:"solution_count,omitempty"`
	InsertCountByDate MapCount `json:"insert_count_by_date,omitempty"`
	UpdateCountByDate MapCount `json:"update_count_by_date,omitempty"`
	DeleteCountByDate MapCount `json:"delete_count_by_date,omitempty"`
}
