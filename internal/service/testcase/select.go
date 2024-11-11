package testcase

import (
	"STUOJ/internal/dao"
	"STUOJ/internal/entity"
)

// 根据ID查询评测点数据
func SelectById(id uint64) (entity.Testcase, error) {
	t, err := dao.SelectTestcaseById(id)
	if err != nil {
		return entity.Testcase{}, err
	}

	return t, nil
}

// 通过题目ID查询评测点数据
func SelectByProblemId(problemId uint64) ([]entity.Testcase, error) {
	testcases, err := dao.SelectTestcasesByProblemId(problemId)
	if err != nil {
		return nil, err
	}

	return testcases, nil
}
