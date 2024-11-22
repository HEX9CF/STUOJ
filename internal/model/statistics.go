package model

// Judge0统计信息
type Judge0Statistics struct {
	LanguageCount   int64           `json:"language_count,omitempty"`
	JudgeStatistics JudgeStatistics `json:"judge_statistics"`
	JudgeSystemInfo JudgeSystemInfo `json:"judge_system_info"`
}

// 用户统计信息
type UserStatistics struct {
	UserCount int64 `json:"user_count,omitempty"`
}

// 提交记录统计信息
type RecordStatistics struct {
	SubmissionCount int64 `json:"submission_count,omitempty"`
	JudgementCount  int64 `json:"judgement_count,omitempty"`
}

// 博客统计信息
type BlogStatistics struct {
	BlogCount       int64    `json:"blog_count,omitempty"`
	BlogCountByDate MapCount `json:"blog_count_by_date,omitempty"`
}

// 评论统计信息
type CommentStatistics struct {
	CommentCount       int64    `json:"comment_count,omitempty"`
	CommentCountByDate MapCount `json:"comment_count_by_date,omitempty"`
}

// 题目统计信息
type ProblemStatistics struct {
	ProblemCount  int64 `json:"problem_count,omitempty"`
	TestcaseCount int64 `json:"testcase_count,omitempty"`
	SolutionCount int64 `json:"solution_count,omitempty"`
}

// 标签统计信息
type TagStatistics struct {
	TagCount          int64    `json:"tag_count,omitempty"`
	ProblemCountByTag MapCount `json:"problem_count_by_tag,omitempty"`
}
