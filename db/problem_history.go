package db

import (
	"STUOJ/model"
	"time"
)

// 插入题目历史记录
func InsertProblemHistory(p model.Problem, uid uint64, op model.Operation) (uint64, error) {
	updateTime := time.Now()
	ph := model.ProblemHistory{
		UserId:       uid,
		ProblemId:    p.Id,
		Title:        p.Title,
		Source:       p.Source,
		Difficulty:   p.Difficulty,
		TimeLimit:    p.TimeLimit,
		MemoryLimit:  p.MemoryLimit,
		Description:  p.Description,
		Input:        p.Input,
		Output:       p.Output,
		SampleInput:  p.SampleInput,
		SampleOutput: p.SampleOutput,
		Hint:         p.Hint,
		Operation:    op,
		CreateTime:   updateTime,
	}
	tx := Db.Create(&ph)
	if tx.Error != nil {
		return 0, tx.Error
	}

	return ph.Id, nil
}

// 根据题目ID查询题目历史记录
func SelectProblemHistoriesByProblemId(pid uint64) ([]model.ProblemHistory, error) {
	var phs []model.ProblemHistory

	tx := Db.Table("tbl_problem_history").Where("problem_id = ?", pid).Find(&phs)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return phs, nil
}
