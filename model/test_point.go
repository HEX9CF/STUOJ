package model

type TestPoint struct {
	Id         uint64 `json:"id,omitempty"`
	Serial     uint64 `json:"serial,omitempty"`
	ProblemId  uint64 `json:"problem_id,omitempty"`
	TestInput  string `json:"test_input,omitempty"`
	TestOutput string `json:"test_output,omitempty"`
}
