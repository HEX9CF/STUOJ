package db

import (
	"STUOJ/internal/model"
)

// 给题目添加标签
func InsertProblemTag(pid uint64, tid uint64) error {
	pt := model.ProblemTag{
		ProblemId: pid,
		TagId:     tid,
	}
	tx := Db.Create(&pt)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func SelectTagsByProblemId(pid uint64) ([]model.Tag, error) {
	var tags []model.Tag

	tx := Db.Table("tbl_tag").Select("tbl_tag.id, tbl_tag.name").Joins("JOIN tbl_problem_tag ON tbl_tag.id = tbl_problem_tag.tag_id").Where("tbl_problem_tag.problem_id = ?", pid).Scan(&tags)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return tags, nil
}

// 查询题目标签关系是否存在
func CountProblemTagByProblemIdAndTagId(pid uint64, tid uint64) (int64, error) {
	var count int64

	tx := Db.Model(&model.ProblemTag{}).Where("problem_id = ? AND tag_id = ?", pid, tid).Count(&count)
	if tx.Error != nil {
		return 0, tx.Error
	}

	return count, nil
}

// 删除题目的某个标签
func DeleteProblemTagByProblemIdAndTagId(pid uint64, tid uint64) error {
	tx := Db.Where("problem_id = ? AND tag_id = ?", pid, tid).Delete(&model.ProblemTag{})
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
