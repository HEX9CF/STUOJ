package model

import "STUOJ/internal/entity"

// 题目信息（题目+标签+评测点数据+题解）
type ProblemInfo struct {
	Problem   entity.Problem    `json:"problem,omitempty"`
	Tags      []entity.Tag      `json:"tags,omitempty"`
	Testcases []entity.Testcase `json:"testcases,omitempty"`
	Solutions []entity.Solution `json:"solution,omitempty"`
}
