package model

// 题目信息（题目+标签+评测点数据）
type ProblemInfo struct {
	Problem   Problem    `json:"problem,omitempty"`
	Tags      []Tag      `json:"tags,omitempty"`
	Testcases []Testcase `json:"testcases,omitempty"`
}
