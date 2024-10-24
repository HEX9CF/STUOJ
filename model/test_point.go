package model

type TestPoint struct {
	Id         uint64 `json:"id"`
	Serial     uint64 `json:"serial"`
	ProblemId  uint64 `json:"problem_id"`
	TestInput  string `json:"test_input"`
	TestOutput string `json:"test_output"`
}
