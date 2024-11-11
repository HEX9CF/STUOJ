package solution

import (
	"STUOJ/internal/dao"
	"STUOJ/internal/entity"
)

// 根据ID更新题解
func UpdateSolutionById(s entity.Solution) error {
	err := dao.UpdateSolutionById(s)
	if err != nil {
		return err
	}

	return nil
}
