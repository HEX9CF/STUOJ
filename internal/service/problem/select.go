package problem

import (
	"STUOJ/internal/dao"
	"STUOJ/internal/model"
)

// 根据ID查询题目
func SelectById(id uint64) (model.Problem, error) {
	p, err := dao.SelectProblemById(id)
	if err != nil {
		return model.Problem{}, err
	}

	return p, nil
}

// 根据状态和ID查询题目
func SelectProblemByIdAndStatus(id uint64, s model.ProblemStatus) (model.Problem, error) {
	p, err := dao.SelectProblemByIdAndStatus(id, s)
	if err != nil {
		return model.Problem{}, err
	}

	return p, nil
}

// 查询所有题目
func SelectAll() ([]model.Problem, error) {
	problems, err := dao.SelectAllProblems()
	if err != nil {
		return nil, err
	}

	return problems, nil
}

// 根据状态查询题目
func SelectByStatus(s model.ProblemStatus) ([]model.Problem, error) {
	problems, err := dao.SelectByStatus(s)
	if err != nil {
		return nil, err
	}

	return problems, nil
}

// 根据状态和标签查询题目
func SelectByTagIdAndStatus(tid uint64, s model.ProblemStatus) ([]model.Problem, error) {
	problems, err := dao.SelectProblemsByTagIdAndStatus(tid, s)
	if err != nil {
		return nil, err
	}

	return problems, nil
}

// 根据状态和难度查询题目
func SelectByDifficultyAndStatus(d model.ProblemDifficulty, s model.ProblemStatus) ([]model.Problem, error) {
	problems, err := dao.SelectProblemsByDifficultyAndStatus(d, s)
	if err != nil {
		return nil, err
	}

	return problems, nil
}

// 根据状态查询并根据标题模糊查询题目
func SelectLikeTitleByStatus(title string, s model.ProblemStatus) ([]model.Problem, error) {
	problems, err := dao.SelectProblemsLikeTitleByStatus(title, s)
	if err != nil {
		return nil, err
	}

	return problems, nil
}

// 根据题目ID查询题目历史记录
func SelectHistoriesByProblemId(pid uint64) ([]model.ProblemHistory, error) {
	phs, err := dao.SelectProblemHistoriesByProblemId(pid)
	if err != nil {
		return nil, err
	}

	return phs, nil
}
