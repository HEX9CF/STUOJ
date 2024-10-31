package model

// 提交记录（提交信息+评测结果）
type Record struct {
	Submission Submission  `json:"submission,omitempty"`
	Judgements []Judgement `json:"judgements,omitempty"`
}
