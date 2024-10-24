package model

// 评测状态
type JudgementStatus uint8

// 单个评测点结果
type Judgement struct {
	Id            uint64  `json:"id"`
	SubmissionId  uint64  `json:"submission_id"`
	TestPointId   uint64  `json:"test_point_id"`
	Time          float64 `json:"time"`
	Memory        uint64  `json:"memory"`
	Stdout        string  `json:"stdout"`
	Stderr        string  `json:"stderr"`
	CompileOutput string  `json:"compile_output"`
	Message       string  `json:"message"`
	Status        uint64  `json:"status"`
}
