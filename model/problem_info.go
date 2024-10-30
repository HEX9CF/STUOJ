package model

// 题目信息（题目+评测点数据）
type ProblemInfo struct {
	Problem   Problem    `json:"problem,omitempty"`
	Testcases []Testcase `json:"testcases,omitempty"`
}
