package testcase

import (
	"STUOJ/internal/dao"
	"STUOJ/internal/model"
)

// 根据ID更新评测点数据
func UpdateById(t model.Testcase) error {
	err := dao.UpdateTestcaseById(t)
	if err != nil {
		return err
	}

	return nil
}
