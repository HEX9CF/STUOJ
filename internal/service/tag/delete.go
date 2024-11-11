package tag

import (
	"STUOJ/internal/dao"
	"STUOJ/internal/model"
)

// 根据ID删除标签
func DeleteById(id uint64) error {
	err := dao.DeleteTagById(id)
	if err != nil {
		return err
	}

	return nil
}

// 删除题目的某个标签
func DeleteProblemTag(pid uint64, tid uint64) error {
	pt := model.ProblemTag{
		ProblemId: pid,
		TagId:     tid,
	}

	err := dao.DeleteProblemTag(pt)
	if err != nil {
		return err
	}

	return nil
}
