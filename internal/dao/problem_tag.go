package dao

import (
	"STUOJ/internal/db"
	"STUOJ/internal/entity"
)

// 给题目添加标签
func InsertProblemTag(pt entity.ProblemTag) error {
	tx := db.Db.Create(&pt)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func SelectTagsByProblemId(pid uint64) ([]entity.Tag, error) {
	var tags []entity.Tag

	tx := db.Db.Table("tbl_tag").Select("tbl_tag.id, tbl_tag.name").Joins("JOIN tbl_problem_tag ON tbl_tag.id = tbl_problem_tag.tag_id").Where("tbl_problem_tag.problem_id = ?", pid).Scan(&tags)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return tags, nil
}

// 查询题目标签关系是否存在
func CountProblemTag(pt entity.ProblemTag) (int64, error) {
	var count int64

	tx := db.Db.Model(&entity.ProblemTag{}).Where("problem_id = ? AND tag_id = ?", pt.ProblemId, pt.TagId).Count(&count)
	if tx.Error != nil {
		return 0, tx.Error
	}

	return count, nil
}

// 删除题目的某个标签
func DeleteProblemTag(pt entity.ProblemTag) error {
	tx := db.Db.Where("problem_id = ? AND tag_id = ?", pt.ProblemId, pt.TagId).Delete(&entity.ProblemTag{})
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// 删除题目的所有标签
func DeleteProblemTagsByProblemId(pid uint64) error {
	tx := db.Db.Where("problem_id = ?", pid).Delete(&entity.ProblemTag{})
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
