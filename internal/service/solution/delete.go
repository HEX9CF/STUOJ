package solution

import "STUOJ/internal/dao"

// 根据ID删除题解
func DeleteSolutionById(id uint64) error {
	err := dao.DeleteSolutionById(id)
	if err != nil {
		return err
	}

	return nil
}
