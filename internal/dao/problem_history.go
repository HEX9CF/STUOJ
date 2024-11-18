package dao

import (
	"STUOJ/internal/db"
	"STUOJ/internal/entity"
	"STUOJ/internal/model"
	"time"
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

	tx := db.Db.Model(&entity.ProblemHistory{}).Where("problem_id = ?", pid).Find(&phs)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return phs, nil
}

// 根据创建时间统计用户数量
func CountProblemHistoriesBetweenCreateTimeByOperation(startTime time.Time, endTime time.Time, operation entity.Operation) ([]model.CountByDate, error) {
	var countByDate []model.CountByDate

	tx := db.Db.Model(&entity.ProblemHistory{}).Where("operation = ? AND create_time between ? and ?", startTime, endTime, operation).Select("date(create_time) as date, count(*) as count").Group("date").Scan(&countByDate)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return countByDate, nil
}
