package tag

import (
	"STUOJ/internal/dao"
	"STUOJ/internal/model"
)

// 根据ID更新标签
func UpdateById(t model.Tag) error {
	err := dao.UpdateTagById(t)
	if err != nil {
		return err
	}

	return nil
}
