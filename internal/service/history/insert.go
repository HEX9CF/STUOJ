package history

import (
	"STUOJ/internal/dao"
	"STUOJ/internal/model"
	"time"
)

// 插入题目历史记录
func InsertProblemHistory(p model.Problem, uid uint64, op model.Operation) (uint64, error) {
	var err error

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

	ph.Id, err = dao.InsertProblemHistory(ph)
	if err != nil {
		return 0, err
	}

	return ph.Id, nil
}
