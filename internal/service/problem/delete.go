package problem

import "STUOJ/internal/dao"

// 根据ID删除题目
func DeleteProblemById(id uint64) error {
	err := dao.DeleteProblemById(id)
	if err != nil {
		return err
	}

	return nil
}
