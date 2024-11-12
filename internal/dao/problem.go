package dao

import (
	"STUOJ/internal/db"
	"STUOJ/internal/entity"
)

// 插入题目
func InsertProblem(p entity.Problem) (uint64, error) {
	tx := db.Db.Create(&p)
	if tx.Error != nil {
		return 0, tx.Error
	}

	return p.Id, nil
}

// 根据ID查询题目
func SelectProblemById(id uint64) (entity.Problem, error) {
	var p entity.Problem

	tx := db.Db.Where("id = ?", id).First(&p)
	if tx.Error != nil {
		return entity.Problem{}, tx.Error
	}

	return p, nil
}

// 根据状态和ID查询题目
func SelectProblemByIdAndStatus(id uint64, s entity.ProblemStatus) (entity.Problem, error) {
	var p entity.Problem

	tx := db.Db.Where("status = ? AND id = ?", s, id).First(&p)
	if tx.Error != nil {
		return entity.Problem{}, tx.Error
	}

	return p, nil
}

// 查询所有题目
func SelectAllProblems() ([]entity.Problem, error) {
	var problems []entity.Problem

	tx := db.Db.Find(&problems)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return problems, nil
}

// 根据状态查询题目
func SelectProblemsByStatus(s entity.ProblemStatus) ([]entity.Problem, error) {
	var problems []entity.Problem

	tx := db.Db.Where("status = ?", s).Find(&problems)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return problems, nil
}

// 根据状态和标签查询题目
func SelectProblemsByTagIdAndStatus(tid uint64, s entity.ProblemStatus) ([]entity.Problem, error) {
	var problems []entity.Problem

	tx := db.Db.Where("status = ? AND id IN (SELECT problem_id FROM tbl_problem_tag WHERE tag_id = ?)", s, tid).Find(&problems)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return problems, nil
}

// 根据状态和难度查询题目
func SelectProblemsByDifficultyAndStatus(d entity.Difficulty, s entity.ProblemStatus) ([]entity.Problem, error) {
	var problems []entity.Problem

	tx := db.Db.Where("status = ? AND difficulty = ?", s, d).Find(&problems)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return problems, nil
}

// 根据状态查询并根据标题模糊查询题目
func SelectProblemsLikeTitleByStatus(title string, s entity.ProblemStatus) ([]entity.Problem, error) {
	var problems []entity.Problem

	tx := db.Db.Where("status = ? AND title like ?", s, "%"+title+"%").Find(&problems)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return problems, nil
}

// 根据ID更新题目
func UpdateProblemById(p entity.Problem) error {
	tx := db.Db.Model(&p).Where("id = ?", p.Id).Updates(p)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// 根据ID删除题目
func DeleteProblemById(id uint64) error {
	tx := db.Db.Where("id = ?", id).Delete(&entity.Problem{})
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
