package dao

import (
	"STUOJ/internal/db"
	"STUOJ/internal/entity"
)

// 插入题目历史记录
func InsertProblemHistory(ph entity.ProblemHistory) (uint64, error) {
	tx := db.Db.Create(&ph)
	if tx.Error != nil {
		return 0, tx.Error
	}

	return ph.Id, nil
}

// 根据题目ID查询题目历史记录
func SelectProblemHistoriesByProblemId(pid uint64) ([]entity.ProblemHistory, error) {
	var phs []entity.ProblemHistory

	tx := db.Db.Table("tbl_problem_history").Where("problem_id = ?", pid).Find(&phs)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return phs, nil
}
