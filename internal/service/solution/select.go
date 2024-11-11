package solution

import (
	"STUOJ/internal/dao"
	"STUOJ/internal/entity"
)

// 根据ID查询题解
func SelectById(id uint64) (entity.Solution, error) {
	s, err := dao.SelectSolutionById(id)
	if err != nil {
		return entity.Solution{}, err
	}

	return s, nil
}

// 查询所有题解（不返回源代码）
func SelectAllSolutions() ([]entity.Solution, error) {
	solutions, err := dao.SelectAllSolutions()
	if err != nil {
		return nil, err
	}

	// 不返回源代码
	for i := range solutions {
		solutions[i].SourceCode = ""
	}

	return solutions, nil
}

// 根据题目ID查询题解（不返回源代码）
func SelectSolutionsByProblemId(pid uint64) ([]entity.Solution, error) {
	solutions, err := dao.SelectSolutionsByProblemId(pid)
	if err != nil {
		return nil, err
	}

	// 不返回源代码
	for i := range solutions {
		solutions[i].SourceCode = ""
	}

	return solutions, nil
}
