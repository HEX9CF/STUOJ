package tag

import (
	"STUOJ/internal/dao"
	"STUOJ/internal/entity"
)

// 根据ID更新标签
func UpdateById(t entity.Tag) error {
	err := dao.UpdateTagById(t)
	if err != nil {
		return err
	}

	return nil
}
