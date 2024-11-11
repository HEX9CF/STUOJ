package problem

import (
	"STUOJ/internal/dao"
	"STUOJ/internal/model"
	"time"
)

// 根据ID更新题目
func UpdateById(p model.Problem) error {
	updateTime := time.Now()
	p.UpdateTime = updateTime

	err := dao.UpdateProblemById(p)
	if err != nil {
		return err
	}

	return nil
}

// 根据ID更新提交记录状态更新时间
func UpdateUpdateTimeById(id uint64) error {
	updateTime := time.Now()

	p := model.Problem{
		Id:         id,
		UpdateTime: updateTime,
	}

	err := dao.UpdateProblemById(p)
	if err != nil {
		return err
	}

	return nil
}
