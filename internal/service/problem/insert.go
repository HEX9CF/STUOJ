package problem

import (
	"STUOJ/internal/dao"
	"STUOJ/internal/model"
	"time"
)

// 插入题目
func Insert(p model.Problem) (uint64, error) {
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
