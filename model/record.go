package model

// 提交记录（包含提交信息和评测点结果）
type Record struct {
	Submission Submission  `json:"submission"`
	Judgements []Judgement `json:"judgements"`
}
