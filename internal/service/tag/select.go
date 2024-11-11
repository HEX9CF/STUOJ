package tag

import (
	"STUOJ/internal/dao"
	"STUOJ/internal/model"
)

// 根据ID查询标签
func SelectById(id uint64) (model.Tag, error) {
	t, err := dao.SelectTagById(id)
	if err != nil {
		return model.Tag{}, err
	}

	return t, nil
}

// 查询所有标签
func SelectAll() ([]model.Tag, error) {
	tags, err := dao.SelectAllTags()
	if err != nil {
		return nil, err
	}

	return tags, nil
}
