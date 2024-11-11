package solution

import (
	"STUOJ/internal/dao"
	"STUOJ/internal/model"
)

// 根据ID更新题解
func UpdateSolutionById(s model.Solution) error {
	err := dao.UpdateSolutionById(s)
	if err != nil {
		return err
	}

	return nil
}
