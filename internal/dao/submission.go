package dao

import (
	"STUOJ/internal/db"
	"STUOJ/internal/entity"
	"time"
)

// 插入提交记录
func InsertSubmission(s entity.Submission) (uint64, error) {
	updateTime := time.Now()
	s.UpdateTime = updateTime
	s.CreateTime = updateTime
	tx := db.Db.Create(&s)
	if tx.Error != nil {
		return 0, tx.Error
	}

	return s.Id, nil
}

// 查询所有提交记录
func SelectAllSubmissions() ([]entity.Submission, error) {
	var submissions []entity.Submission

	tx := db.Db.Find(&submissions)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return submissions, nil
}

// 根据ID查询提交记录
func SelectSubmissionById(id uint64) (entity.Submission, error) {
	var s entity.Submission

	tx := db.Db.Where("id = ?", id).First(&s)
	if tx.Error != nil {
		return entity.Submission{}, tx.Error
	}

	return s, nil
}

// 根据用户ID查询提交记录（不返回源代码）
func SelectSubmissionsByUserId(userId uint64) ([]entity.Submission, error) {
	var submissions []entity.Submission

	tx := db.Db.Where("user_id = ?", userId).Find(&submissions)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return submissions, nil
}

// 根据题目ID查询提交记录（不返回源代码）
func SelectSubmissionsByProblemId(problemId uint64) ([]entity.Submission, error) {
	var submissions []entity.Submission

	tx := db.Db.Where("problem_id = ?", problemId).Find(&submissions)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return submissions, nil
}

// 更新提交记录
func UpdateSubmissionById(s entity.Submission) error {
	tx := db.Db.Model(&s).Where("id = ?", s.Id).Updates(s)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// 根据ID删除提交记录
func DeleteSubmissionById(id uint64) error {
	tx := db.Db.Where("id = ?", id).Delete(&entity.Submission{})
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
