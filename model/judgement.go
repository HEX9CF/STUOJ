package model

// 单个评测点结果
type Judgement struct {
	Id            uint64       `json:"id"`
	SubmissionId  uint64       `json:"submission_id"`
	TestPointId   uint64       `json:"test_point_id"`
	Time          float64      `json:"time"`
	Memory        uint64       `json:"memory"`
	Stdout        string       `json:"stdout,omitempty"`
	Stderr        string       `json:"stderr,omitempty"`
	CompileOutput string       `json:"compile_output,omitempty"`
	Message       string       `json:"message,omitempty"`
	Status        SubmitStatus `json:"status"`
}
