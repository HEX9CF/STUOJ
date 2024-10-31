package model

// 评测点结果
type Judgement struct {
	Id            uint64       `json:"id,omitempty"`
	SubmissionId  uint64       `json:"submission_id,omitempty"`
	TestcaseId    uint64       `json:"testcase_id,omitempty"`
	Time          float64      `json:"time,omitempty"`
	Memory        uint64       `json:"memory,omitempty"`
	Stdout        string       `json:"stdout,omitempty"`
	Stderr        string       `json:"stderr,omitempty"`
	CompileOutput string       `json:"compile_output,omitempty"`
	Message       string       `json:"message,omitempty"`
	Status        SubmitStatus `json:"status,omitempty"`
}
