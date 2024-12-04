package dao

import (
	"STUOJ/internal/db"
	"STUOJ/internal/entity"
	"STUOJ/internal/model"
	"time"
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
func SelectProblemsByStatus(s entity.ProblemStatus, page uint64, size uint64) ([]entity.Problem, error) {
	var problems []entity.Problem

	tx := db.Db.Offset(int((page-1)*size)).Limit(int(size)).Where("status = ?", s).Find(&problems)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return problems, nil
}

// 根据状态和标签查询题目
func SelectProblemsByTagIdAndStatus(tid uint64, s entity.ProblemStatus, page uint64, size uint64) ([]entity.Problem, error) {
	var problems []entity.Problem

	tx := db.Db.Offset(int((page-1)*size)).Limit(int(size)).Where("status = ? AND id IN (SELECT problem_id FROM tbl_problem_tag WHERE tag_id = ?)", s, tid).Find(&problems)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return problems, nil
}

// 根据状态和难度查询题目
func SelectProblemsByDifficultyAndStatus(d entity.Difficulty, s entity.ProblemStatus, page uint64, size uint64) ([]entity.Problem, error) {
	var problems []entity.Problem

	tx := db.Db.Offset(int((page-1)*size)).Limit(int(size)).Where("status = ? AND difficulty = ?", s, d).Find(&problems)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return problems, nil
}

// 根据状态查询并根据标题模糊查询题目
func SelectProblemsLikeTitleByStatus(title string, s entity.ProblemStatus, page uint64, size uint64) ([]entity.Problem, error) {
	var problems []entity.Problem

	tx := db.Db.Offset(int((page-1)*size)).Limit(int(size)).Where("status = ? AND title like ?", s, "%"+title+"%").Find(&problems)
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

// 根据ID更新题目更新时间
func UpdateProblemUpdateTimeById(id uint64) error {
	tx := db.Db.Model(&entity.Problem{}).Where("id = ?", id).Update("update_time", time.Now())
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

// 统计题目数量
func CountProblems() (uint64, error) {
	var count int64

	tx := db.Db.Model(&entity.Problem{}).Count(&count)
	if tx.Error != nil {
		return 0, tx.Error
	}

	return uint64(count), nil
}

// 根据创建时间统计用户数量
func CountProblemsBetweenCreateTime(startTime time.Time, endTime time.Time) ([]model.CountByDate, error) {
	var countByDate []model.CountByDate

	tx := db.Db.Model(&entity.Problem{}).Where("create_time between ? and ?", startTime, endTime).Select("date(create_time) as date, count(*) as count").Group("date").Scan(&countByDate)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return countByDate, nil
}
