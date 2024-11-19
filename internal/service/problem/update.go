package problem

import (
	"STUOJ/internal/dao"
	"STUOJ/internal/entity"
	"errors"
	"log"
	"time"
)

// 根据ID更新题目
func UpdateById(p entity.Problem, uid uint64) error {
	// 读取题目
	p0, err := dao.SelectProblemById(p.Id)
	if err != nil {
		log.Println(err)
		return errors.New("题目不存在")
	}

	updateTime := time.Now()

	// 添加题目历史记录
	ph := entity.History{
		UserId:       uid,
		ProblemId:    p0.Id,
		Title:        p0.Title,
		Source:       p0.Source,
		Difficulty:   p0.Difficulty,
		TimeLimit:    p0.TimeLimit,
		MemoryLimit:  p0.MemoryLimit,
		Description:  p0.Description,
		Input:        p0.Input,
		Output:       p0.Output,
		SampleInput:  p0.SampleInput,
		SampleOutput: p0.SampleOutput,
		Hint:         p0.Hint,
		Operation:    entity.OperationUpdate,
		CreateTime:   updateTime,
	}
	ph.Id, err = dao.InsertHistory(ph)
	if err != nil {
		log.Println(err)
		return errors.New("更新题目成功，但插入题目历史记录失败")
	}

	// 更新题目
	p0.UpdateTime = updateTime
	err = dao.UpdateProblemById(p0)
	if err != nil {
		return errors.New("更新题目失败")
	}

	return nil
}
