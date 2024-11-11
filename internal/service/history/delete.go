package history

import (
	"STUOJ/internal/dao"
	"STUOJ/internal/model"
)

// 根据题目ID查询题目历史记录
func SelectProblemHistoriesByProblemId(pid uint64) ([]model.ProblemHistory, error) {
	phs, err := dao.SelectProblemHistoriesByProblemId(pid)
	if err != nil {
		return nil, err
	}

	return phs, nil
}
