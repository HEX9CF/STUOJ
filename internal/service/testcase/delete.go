package testcase

import "STUOJ/internal/dao"

// 根据ID删除评测点数据
func DeleteById(id uint64) error {
	err := dao.DeleteTestcaseById(id)
	if err != nil {
		return err
	}

	return nil
}
