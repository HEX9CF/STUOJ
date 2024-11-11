package tag

import (
	"STUOJ/internal/dao"
)

// 根据ID删除标签
func DeleteById(id uint64) error {
	err := dao.DeleteTagById(id)
	if err != nil {
		return err
	}

	return nil
}
