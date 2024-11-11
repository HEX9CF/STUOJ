package tag

import (
	"STUOJ/internal/dao"
	"STUOJ/internal/entity"
)

// 根据ID查询标签
func SelectById(id uint64) (entity.Tag, error) {
	t, err := dao.SelectTagById(id)
	if err != nil {
		return entity.Tag{}, err
	}

	return t, nil
}

// 查询所有标签
func SelectAll() ([]entity.Tag, error) {
	tags, err := dao.SelectAllTags()
	if err != nil {
		return nil, err
	}

	return tags, nil
}

// 根据题目ID查询标签
func SelectByProblemId(pid uint64) ([]entity.Tag, error) {
	tags, err := dao.SelectTagsByProblemId(pid)
	if err != nil {
		return nil, err
	}

	return tags, nil
}

// 查询题目标签关系是否存在
func CountProblemTag(pid uint64, tid uint64) (int64, error) {
	pt := entity.ProblemTag{
		ProblemId: pid,
		TagId:     tid,
	}

	count, err := dao.CountProblemTag(pt)
	if err != nil {
		return 0, err
	}

	return count, nil
}
