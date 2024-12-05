package model

// Judge0统计信息
type Judge0Statistics struct {
	LanguageCount   uint64          `json:"language_count,omitempty"`
	JudgeStatistics JudgeStatistics `json:"judge_statistics"`
	JudgeSystemInfo JudgeSystemInfo `json:"judge_system_info"`
}

// 用户统计信息
type UserStatistics struct {
	UserCount uint64 `json:"user_count,omitempty"`
}

// 提交记录统计信息
type RecordStatistics struct {
	SubmissionCount uint64 `json:"submission_count,omitempty"`
	JudgementCount  uint64 `json:"judgement_count,omitempty"`
}

// 博客统计信息
type BlogStatistics struct {
	BlogCount       uint64   `json:"blog_count,omitempty"`
	CommentCount    uint64   `json:"comment_count,omitempty"`
	BlogCountByDate MapCount `json:"blog_count_by_date,omitempty"`
}

// 评论统计信息
type CommentStatistics struct {
	CommentCountByDate MapCount `json:"comment_count_by_date,omitempty"`
}

// 题目统计信息
type ProblemStatistics struct {
	ProblemCount  uint64 `json:"problem_count,omitempty"`
	TestcaseCount uint64 `json:"testcase_count,omitempty"`
	SolutionCount uint64 `json:"solution_count,omitempty"`
}

// 标签统计信息
type TagStatistics struct {
	TagCount          uint64   `json:"tag_count,omitempty"`
	ProblemCountByTag MapCount `json:"problem_count_by_tag,omitempty"`
}
