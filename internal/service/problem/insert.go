package problem

import (
	"STUOJ/internal/dao"
	"STUOJ/internal/entity"
	"time"
)

// 插入题目
func Insert(p entity.Problem) (uint64, error) {
	var err error

	updateTime := time.Now()
	p.UpdateTime = updateTime
	p.CreateTime = updateTime

	p.Id, err = dao.InsertProblem(p)
	if err != nil {
		return 0, err
	}

	return p.Id, nil
}

// 给题目添加标签
func InsertTag(pid uint64, tid uint64) error {
	pt := entity.ProblemTag{
		ProblemId: pid,
		TagId:     tid,
	}

	err := dao.InsertProblemTag(pt)
	if err != nil {
		return err
	}

	return nil
}

// 插入题目历史记录
func InsertHistory(p entity.Problem, uid uint64, op entity.Operation) (uint64, error) {
	var err error

	updateTime := time.Now()
	ph := entity.ProblemHistory{
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
