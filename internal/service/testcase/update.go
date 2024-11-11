package testcase

import (
	"STUOJ/internal/dao"
	"STUOJ/internal/entity"
)

// 根据ID更新评测点数据
func UpdateById(t entity.Testcase) error {
	err := dao.UpdateTestcaseById(t)
	if err != nil {
		return err
	}

	return nil
}
