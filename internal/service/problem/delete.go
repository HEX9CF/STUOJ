package problem

import (
	"STUOJ/internal/dao"
	"STUOJ/internal/entity"
)

// 根据ID删除题目
func DeleteProblemById(id uint64) error {
	err := dao.DeleteProblemById(id)
	if err != nil {
		return err
	}

	return nil
}

// 删除题目的某个标签
func DeleteTag(pid uint64, tid uint64) error {
	pt := entity.ProblemTag{
		ProblemId: pid,
		TagId:     tid,
	}

	err := dao.DeleteProblemTag(pt)
	if err != nil {
		return err
	}

	return nil
}
