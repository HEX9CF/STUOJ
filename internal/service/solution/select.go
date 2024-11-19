package solution

import (
	"STUOJ/internal/dao"
	"STUOJ/internal/entity"
	"errors"
)

// 根据ID查询题解
func SelectById(id uint64) (entity.Solution, error) {
	s, err := dao.SelectSolutionById(id)
	if err != nil {
		return entity.Solution{}, errors.New("获取题解失败")
	}

	return s, nil
}

// 查询所有题解（不返回源代码）
func SelectAll() ([]entity.Solution, error) {
	solutions, err := dao.SelectAllSolutions()
	if err != nil {
		return nil, err
	}

	hideSourceCode(solutions)

	return solutions, nil
}

// 根据题目ID查询题解（不返回源代码）
func SelectByProblemId(pid uint64) ([]entity.Solution, error) {
	solutions, err := dao.SelectSolutionsByProblemId(pid)
	if err != nil {
		return nil, err
	}

	hideSourceCode(solutions)

	return solutions, nil
}

// 隐藏源代码
func hideSourceCode(solutions []entity.Solution) {
	for i := range solutions {
		solutions[i].SourceCode = ""
	}
}
